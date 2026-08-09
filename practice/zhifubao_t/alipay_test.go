package zhifubao_t

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"encoding/pem"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"path/filepath"
	"testing"
)

// newTestKey 生成临时 RSA 密钥对，返回私钥与 PEM 编码的应用私钥、支付宝公钥
func newTestKey(t *testing.T) (*rsa.PrivateKey, string, string) {
	t.Helper()
	key, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		t.Fatalf("生成 RSA 密钥失败: %v", err)
	}
	privDER, err := x509.MarshalPKCS8PrivateKey(key)
	if err != nil {
		t.Fatalf("编码私钥失败: %v", err)
	}
	pubDER, err := x509.MarshalPKIXPublicKey(&key.PublicKey)
	if err != nil {
		t.Fatalf("编码公钥失败: %v", err)
	}
	return key,
		string(pem.EncodeToMemory(&pem.Block{Type: "PRIVATE KEY", Bytes: privDER})),
		string(pem.EncodeToMemory(&pem.Block{Type: "PUBLIC KEY", Bytes: pubDER}))
}

// newClientFromPEM 使用指定密钥 PEM 构造指向目标网关的测试客户端
func newClientFromPEM(t *testing.T, privPEM, pubPEM, gatewayURL string) *Client {
	t.Helper()
	client, err := NewClient(Config{
		AppID:           "2021000000000000",
		PrivateKey:      privPEM,
		AlipayPublicKey: pubPEM,
		GatewayURL:      gatewayURL,
		NotifyURL:       "https://example.com/api/payment/alipay/notify",
	})
	if err != nil {
		t.Fatalf("创建客户端失败: %v", err)
	}
	return client
}

// newTestClient 生成临时 RSA 密钥并构造一个仅用于本地验签测试的客户端
func newTestClient(t *testing.T) *Client {
	t.Helper()
	_, privPEM, pubPEM := newTestKey(t)
	return newClientFromPEM(t, privPEM, pubPEM, DefaultGatewayURL)
}

// TestBuildSignString 验证签名字符串按 key 升序拼接并过滤空值
func TestBuildSignString(t *testing.T) {
	params := url.Values{}
	params.Set("b", "2")
	params.Set("a", "1")
	params.Set("empty", "")
	if got, want := buildSignString(params), "a=1&b=2"; got != want {
		t.Fatalf("buildSignString = %q, 期望 %q", got, want)
	}
}

// TestSignAndVerify 验证 RSA2 签名与验签往返，以及篡改参数后验签失败
func TestSignAndVerify(t *testing.T) {
	client := newTestClient(t)

	params := url.Values{}
	params.Set("app_id", client.config.AppID)
	params.Set("out_trade_no", "202608071200000001")
	params.Set("total_amount", "0.01")
	params.Set("trade_status", "TRADE_SUCCESS")

	sign, err := client.signParams(params)
	if err != nil {
		t.Fatalf("签名失败: %v", err)
	}
	if err := client.verifySign(params, sign); err != nil {
		t.Fatalf("验签失败: %v", err)
	}

	// 篡改任意参数后验签必须失败
	params.Set("total_amount", "0.02")
	if err := client.verifySign(params, sign); err == nil {
		t.Fatal("篡改参数后验签仍通过，签名校验存在漏洞")
	}
}

// TestVerifyNotify 模拟支付宝异步通知的签名与验签流程
func TestVerifyNotify(t *testing.T) {
	client := newTestClient(t)

	// 模拟支付宝服务端：对除 sign/sign_type 外的参数签名
	signContent := url.Values{}
	signContent.Set("app_id", client.config.AppID)
	signContent.Set("out_trade_no", "202608071200000002")
	signContent.Set("total_amount", "0.01")
	signContent.Set("trade_status", "TRADE_SUCCESS")
	signContent.Set("timestamp", "2026-08-07 12:00:00")
	sign, err := client.signParams(signContent)
	if err != nil {
		t.Fatalf("生成通知签名失败: %v", err)
	}

	// 通知表单包含 sign 与 sign_type 字段
	notify := url.Values{}
	for key, values := range signContent {
		notify[key] = values
	}
	notify.Set("sign_type", SignTypeRSA2)
	notify.Set("sign", sign)

	verified, err := client.VerifyNotify(notify)
	if err != nil {
		t.Fatalf("异步通知验签失败: %v", err)
	}
	if verified.Get("out_trade_no") != "202608071200000002" {
		t.Fatalf("验签后订单号不一致: %s", verified.Get("out_trade_no"))
	}
	if verified.Get("sign") != "" || verified.Get("sign_type") != "" {
		t.Fatal("验签结果不应包含 sign 或 sign_type")
	}

	// 篡改金额后验签必须失败
	notify.Set("total_amount", "100.00")
	if _, err := client.VerifyNotify(notify); err == nil {
		t.Fatal("篡改金额后异步通知验签仍通过")
	}
}

// TestLoadConfig 验证 YAML 配置加载与默认值填充
func TestLoadConfig(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "config.yaml")
	content := `alipay:
  app_id: "2021000000000000"
  private_key: "PRIVATE_KEY"
  alipay_public_key: "PUBLIC_KEY"
  notify_url: "https://example.com/notify"
`
	if err := os.WriteFile(path, []byte(content), 0o600); err != nil {
		t.Fatalf("写入配置文件失败: %v", err)
	}
	cfg, err := LoadConfig(path)
	if err != nil {
		t.Fatalf("加载配置失败: %v", err)
	}
	if cfg.AppID != "2021000000000000" {
		t.Fatalf("AppID = %s", cfg.AppID)
	}
	if cfg.GatewayURL != DefaultGatewayURL {
		t.Fatalf("网关默认值未填充: %s", cfg.GatewayURL)
	}
	if cfg.SignType != SignTypeRSA2 {
		t.Fatalf("签名类型默认值未填充: %s", cfg.SignType)
	}
}

// signRaw 使用私钥对原始内容做 RSA2 签名，模拟支付宝服务端对响应原文签名
func signRaw(t *testing.T, key *rsa.PrivateKey, content string) string {
	t.Helper()
	digest := sha256.Sum256([]byte(content))
	signature, err := rsa.SignPKCS1v15(rand.Reader, key, crypto.SHA256, digest[:])
	if err != nil {
		t.Fatalf("生成签名失败: %v", err)
	}
	return base64.StdEncoding.EncodeToString(signature)
}

// TestPrecreate 验证扫码支付预下单：mock 网关校验请求参数并返回签名响应，
// 客户端应能完成响应验签并解析出二维码内容。
func TestPrecreate(t *testing.T) {
	key, privPEM, pubPEM := newTestKey(t)
	gateway := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if err := r.ParseForm(); err != nil {
			http.Error(w, "请求解析失败", http.StatusBadRequest)
			return
		}
		if got := r.Form.Get("method"); got != MethodTradePrecreate {
			t.Errorf("method = %s, 期望 %s", got, MethodTradePrecreate)
		}
		var biz map[string]any
		if err := json.Unmarshal([]byte(r.Form.Get("biz_content")), &biz); err != nil {
			t.Errorf("biz_content 解析失败: %v", err)
		}
		outTradeNo, _ := biz["out_trade_no"].(string)

		// 模拟支付宝：对业务响应字段原文签名后返回
		bizRaw, err := json.Marshal(map[string]string{
			"code":         "10000",
			"msg":          "Success",
			"out_trade_no": outTradeNo,
			"qr_code":      "https://qr.alipay.com/example123",
		})
		if err != nil {
			t.Errorf("序列化业务响应失败: %v", err)
			return
		}
		sign := signRaw(t, key, string(bizRaw))
		resp := struct {
			AlipayTradePrecreateResponse json.RawMessage `json:"alipay_trade_precreate_response"`
			Sign                         string          `json:"sign"`
		}{json.RawMessage(bizRaw), sign}
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(resp); err != nil {
			t.Errorf("写响应失败: %v", err)
		}
	}))
	defer gateway.Close()

	client := newClientFromPEM(t, privPEM, pubPEM, gateway.URL)
	result, err := client.Precreate(PrecreateParams{
		OutTradeNo:  "202608071200000003",
		TotalAmount: "0.01",
		Subject:     "测试商品",
	})
	if err != nil {
		t.Fatalf("扫码预下单失败: %v", err)
	}
	if result.QrCode != "https://qr.alipay.com/example123" {
		t.Fatalf("QrCode = %s", result.QrCode)
	}
	if result.OutTradeNo != "202608071200000003" {
		t.Fatalf("OutTradeNo = %s", result.OutTradeNo)
	}
}

// TestPrecreateBadSign 验证网关响应签名被篡改时预下单必须失败
func TestPrecreateBadSign(t *testing.T) {
	_, privPEM, pubPEM := newTestKey(t)
	gateway := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		// 返回签名与业务响应不匹配的响应
		fmt.Fprint(w, `{"alipay_trade_precreate_response":{"code":"10000","msg":"Success","out_trade_no":"x","qr_code":"y"},"sign":"bm90LWEtdmFsaWQtc2lnbmF0dXJl"}`)
	}))
	defer gateway.Close()

	client := newClientFromPEM(t, privPEM, pubPEM, gateway.URL)
	if _, err := client.Precreate(PrecreateParams{
		OutTradeNo:  "202608071200000004",
		TotalAmount: "0.01",
		Subject:     "测试商品",
	}); err == nil {
		t.Fatal("篡改响应签名后预下单仍通过")
	}
}
