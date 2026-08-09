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
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"sort"
	"strings"
	"time"
)

// Client 支付宝开放平台客户端：负责密钥解析、RSA2 签名、结果验签与网关通信
type Client struct {
	config          Config          // 应用配置
	privateKey      *rsa.PrivateKey // 应用私钥，用于请求签名
	alipayPublicKey *rsa.PublicKey  // 支付宝公钥，用于响应与回调验签
	httpClient      *http.Client    // HTTP 客户端
}

// gatewayResponse 网关通用响应结构，业务字段用 RawMessage 保留原文用于验签
type gatewayResponse struct {
	AlipayTradePrecreateResponse json.RawMessage `json:"alipay_trade_precreate_response"`
	AlipayTradeQueryResponse     json.RawMessage `json:"alipay_trade_query_response"`
	AlipayTradeRefundResponse    json.RawMessage `json:"alipay_trade_refund_response"`
	Sign                         string          `json:"sign"`
}

// raw 按接口方法名返回对应的业务响应字段原文
func (g *gatewayResponse) raw(method string) json.RawMessage {
	switch method {
	case MethodTradePrecreate:
		return g.AlipayTradePrecreateResponse
	case MethodTradeQuery:
		return g.AlipayTradeQueryResponse
	case MethodTradeRefund:
		return g.AlipayTradeRefundResponse
	default:
		return nil
	}
}

// NewClient 根据配置创建支付宝客户端，并解析应用私钥与支付宝公钥
func NewClient(config Config) (*Client, error) {
	config.applyDefaults()

	privateKey, err := parsePrivateKey(config.PrivateKey)
	if err != nil {
		return nil, fmt.Errorf("解析应用私钥失败: %w", err)
	}
	alipayPublicKey, err := parseAlipayPublicKey(config.AlipayPublicKey)
	if err != nil {
		return nil, fmt.Errorf("解析支付宝公钥失败: %w", err)
	}
	return &Client{
		config:          config,
		privateKey:      privateKey,
		alipayPublicKey: alipayPublicKey,
		httpClient:      &http.Client{Timeout: 15 * time.Second},
	}, nil
}

// parsePrivateKey 解析应用私钥，兼容 PKCS8/PKCS1 PEM 与裸 Base64 三种格式
func parsePrivateKey(key string) (*rsa.PrivateKey, error) {
	der, err := decodePEMOrBase64(key)
	if err != nil {
		return nil, err
	}
	if parsed, err := x509.ParsePKCS8PrivateKey(der); err == nil {
		rsaKey, ok := parsed.(*rsa.PrivateKey)
		if !ok {
			return nil, errors.New("私钥不是 RSA 密钥")
		}
		return rsaKey, nil
	}
	if parsed, err := x509.ParsePKCS1PrivateKey(der); err == nil {
		return parsed, nil
	}
	return nil, errors.New("无法解析私钥，请确认是 PKCS8/PKCS1 格式")
}

// parseAlipayPublicKey 解析支付宝公钥，兼容 PKIX/RSA PEM 与裸 Base64 三种格式
func parseAlipayPublicKey(key string) (*rsa.PublicKey, error) {
	der, err := decodePEMOrBase64(key)
	if err != nil {
		return nil, err
	}
	if parsed, err := x509.ParsePKIXPublicKey(der); err == nil {
		rsaKey, ok := parsed.(*rsa.PublicKey)
		if !ok {
			return nil, errors.New("公钥不是 RSA 公钥")
		}
		return rsaKey, nil
	}
	if parsed, err := x509.ParsePKCS1PublicKey(der); err == nil {
		return parsed, nil
	}
	return nil, errors.New("无法解析公钥，请确认是支付宝公钥")
}

// decodePEMOrBase64 将 PEM 或裸 Base64 文本解码为 DER 字节
func decodePEMOrBase64(data string) ([]byte, error) {
	trimmed := strings.TrimSpace(data)
	if block, _ := pem.Decode([]byte(trimmed)); block != nil {
		return block.Bytes, nil
	}
	return base64.StdEncoding.DecodeString(trimmed)
}

// buildSignString 按支付宝规则生成待签名字符串：
// 过滤空值后按 key 升序排列，拼接为 k1=v1&k2=v2...
func buildSignString(params url.Values) string {
	keys := make([]string, 0, len(params))
	for key := range params {
		if params.Get(key) != "" {
			keys = append(keys, key)
		}
	}
	sort.Strings(keys)

	parts := make([]string, 0, len(keys))
	for _, key := range keys {
		parts = append(parts, key+"="+params.Get(key))
	}
	return strings.Join(parts, "&")
}

// signParams 对参数集合进行 RSA2 签名，返回 Base64 编码的签名
func (c *Client) signParams(params url.Values) (string, error) {
	digest := sha256.Sum256([]byte(buildSignString(params)))
	signature, err := rsa.SignPKCS1v15(rand.Reader, c.privateKey, crypto.SHA256, digest[:])
	if err != nil {
		return "", fmt.Errorf("生成签名失败: %w", err)
	}
	return base64.StdEncoding.EncodeToString(signature), nil
}

// verifySign 使用支付宝公钥校验参数集合的 RSA2 签名
func (c *Client) verifySign(params url.Values, sign string) error {
	return c.verifyContent(buildSignString(params), sign)
}

// verifyContent 校验任意待签名内容（参数拼接串或响应 JSON 原文）与签名是否匹配
func (c *Client) verifyContent(content, sign string) error {
	if sign == "" {
		return errors.New("缺少签名参数 sign")
	}
	digest := sha256.Sum256([]byte(content))
	signature, err := base64.StdEncoding.DecodeString(sign)
	if err != nil {
		return fmt.Errorf("签名 Base64 解码失败: %w", err)
	}
	if err := rsa.VerifyPKCS1v15(c.alipayPublicKey, crypto.SHA256, digest[:], signature); err != nil {
		return fmt.Errorf("签名校验失败: %w", err)
	}
	return nil
}

// buildParams 组装某接口的公共参数并完成 RSA2 签名，返回带 sign 的参数集合
func (c *Client) buildParams(method string, bizContent map[string]any) (url.Values, error) {
	bizJSON, err := json.Marshal(bizContent)
	if err != nil {
		return nil, fmt.Errorf("序列化 biz_content 失败: %w", err)
	}

	params := url.Values{}
	params.Set("app_id", c.config.AppID)
	params.Set("method", method)
	params.Set("format", c.config.Format)
	params.Set("charset", c.config.Charset)
	params.Set("sign_type", c.config.SignType)
	params.Set("timestamp", time.Now().Format("2006-01-02 15:04:05"))
	params.Set("version", VersionAPI)
	params.Set("biz_content", string(bizJSON))
	if c.config.NotifyURL != "" {
		params.Set("notify_url", c.config.NotifyURL)
	}

	sign, err := c.signParams(params)
	if err != nil {
		return nil, err
	}
	params.Set("sign", sign)
	return params, nil
}

// doRequest 向支付宝网关发起表单请求，校验响应签名后返回业务响应字段原文
func (c *Client) doRequest(method string, bizContent map[string]any) (json.RawMessage, error) {
	params, err := c.buildParams(method, bizContent)
	if err != nil {
		return nil, err
	}

	resp, err := c.httpClient.PostForm(c.config.GatewayURL, params)
	if err != nil {
		return nil, fmt.Errorf("请求支付宝网关失败: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("读取网关响应失败: %w", err)
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("网关返回异常状态码 %d: %s", resp.StatusCode, string(body))
	}

	var gw gatewayResponse
	if err := json.Unmarshal(body, &gw); err != nil {
		return nil, fmt.Errorf("解析网关响应失败: %w", err)
	}
	raw := gw.raw(method)
	if len(raw) == 0 {
		return nil, fmt.Errorf("网关响应缺少业务字段 %s", method)
	}
	// 支付宝对业务响应字段的 JSON 原文做签名，验签通过才算可信响应
	if err := c.verifyContent(string(raw), gw.Sign); err != nil {
		return nil, err
	}
	return raw, nil
}
