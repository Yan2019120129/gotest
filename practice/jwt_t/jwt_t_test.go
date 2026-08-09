package jwt_t

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// newTestServer 创建测试用服务端与 httptest 测试服务
func newTestServer(t *testing.T) (*Server, *httptest.Server) {
	t.Helper()
	server, err := NewServer()
	if err != nil {
		t.Fatalf("创建服务端失败：%v", err)
	}
	ts := httptest.NewServer(server.ServeMux())
	t.Cleanup(ts.Close)
	return server, ts
}

// signTestToken 使用服务端私钥签发指定 claims 的令牌，用于构造过期等特殊场景
func signTestToken(t *testing.T, server *Server, claims Claims) string {
	t.Helper()
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	signed, err := token.SignedString(server.privateKey)
	if err != nil {
		t.Fatalf("签发测试令牌失败：%v", err)
	}
	return signed
}

// doLogin 直接请求登录接口并返回响应状态码与内容
func doLogin(t *testing.T, ts *httptest.Server, username, password string) (int, map[string]interface{}) {
	t.Helper()
	body := fmt.Sprintf(`{"username":%q,"password":%q}`, username, password)
	resp, err := ts.Client().Post(ts.URL+LoginPath, "application/json", strings.NewReader(body))
	if err != nil {
		t.Fatalf("请求登录接口失败：%v", err)
	}
	defer resp.Body.Close()
	var data map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		t.Fatalf("解析登录响应失败：%v", err)
	}
	return resp.StatusCode, data
}

// doMe 携带指定令牌请求受保护接口，返回响应状态码
func doMe(t *testing.T, ts *httptest.Server, token string) int {
	t.Helper()
	req, err := http.NewRequest(http.MethodGet, ts.URL+MePath, nil)
	if err != nil {
		t.Fatalf("构造请求失败：%v", err)
	}
	if token != "" {
		req.Header.Set("Authorization", "Bearer "+token)
	}
	resp, err := ts.Client().Do(req)
	if err != nil {
		t.Fatalf("请求受保护接口失败：%v", err)
	}
	defer resp.Body.Close()
	return resp.StatusCode
}

// TestLoginSuccess 验证正确的用户名密码能换取令牌
func TestLoginSuccess(t *testing.T) {
	_, ts := newTestServer(t)
	status, data := doLogin(t, ts, "admin", "123456")
	if status != http.StatusOK {
		t.Fatalf("登录成功应返回 200，实际得到 %d", status)
	}
	if token, _ := data["token"].(string); token == "" {
		t.Fatal("登录成功但未返回令牌")
	}
	if _, ok := data["expires_at"]; !ok {
		t.Fatal("登录成功但未返回过期时间")
	}
}

// TestLoginWrongPassword 验证密码错误返回 401 与错误信息
func TestLoginWrongPassword(t *testing.T) {
	_, ts := newTestServer(t)
	status, data := doLogin(t, ts, "admin", "wrong")
	if status != http.StatusUnauthorized {
		t.Fatalf("密码错误应返回 401，实际得到 %d", status)
	}
	if msg, _ := data["error"].(string); msg != "用户名或密码错误" {
		t.Fatalf("错误信息不正确：%v", data["error"])
	}
}

// TestMeWithValidToken 验证携带有效令牌可以访问受保护接口
func TestMeWithValidToken(t *testing.T) {
	server, ts := newTestServer(t)
	token, _, err := server.createToken(1, "admin", "admin")
	if err != nil {
		t.Fatalf("签发令牌失败：%v", err)
	}
	status := doMe(t, ts, token)
	if status != http.StatusOK {
		t.Fatalf("有效令牌应返回 200，实际得到 %d", status)
	}
}

// TestMeWithoutToken 验证缺少令牌访问受保护接口返回 401
func TestMeWithoutToken(t *testing.T) {
	_, ts := newTestServer(t)
	if status := doMe(t, ts, ""); status != http.StatusUnauthorized {
		t.Fatalf("缺少令牌应返回 401，实际得到 %d", status)
	}
}

// TestMeWithTamperedToken 验证篡改签名后的令牌返回 401
func TestMeWithTamperedToken(t *testing.T) {
	server, ts := newTestServer(t)
	valid, _, err := server.createToken(1, "admin", "admin")
	if err != nil {
		t.Fatalf("签发令牌失败：%v", err)
	}
	parts := strings.Split(valid, ".")
	if len(parts) != 3 {
		t.Fatalf("令牌格式异常：%s", valid)
	}
	// 篡改签名部分最后一个字符
	sig := []byte(parts[2])
	sig[len(sig)-1] ^= 0x01
	tampered := parts[0] + "." + parts[1] + "." + string(sig)
	if status := doMe(t, ts, tampered); status != http.StatusUnauthorized {
		t.Fatalf("篡改令牌应返回 401，实际得到 %d", status)
	}
}

// TestMeWithRepackagedToken 验证篡改载荷用户内容并沿用原签名重新封装的令牌返回 401
func TestMeWithRepackagedToken(t *testing.T) {
	server, ts := newTestServer(t)
	valid, _, err := server.createToken(1, "admin", "admin")
	if err != nil {
		t.Fatalf("签发令牌失败：%v", err)
	}
	// 解析原令牌并篡改用户内容，沿用原签名重新封装
	claims := new(Claims)
	_, parts, err := jwt.NewParser().ParseUnverified(valid, claims)
	if err != nil {
		t.Fatalf("解析令牌失败：%v", err)
	}
	claims.Username = "root"
	claims.Role = "superadmin"
	rebuilt, err := jwt.NewWithClaims(jwt.SigningMethodRS256, claims).SigningString()
	if err != nil {
		t.Fatalf("重新封装令牌失败：%v", err)
	}
	forged := rebuilt + "." + parts[2]
	if status := doMe(t, ts, forged); status != http.StatusUnauthorized {
		t.Fatalf("篡改并重新封装的令牌应返回 401，实际得到 %d", status)
	}
}

// TestMeWithExpiredToken 验证过期令牌返回 401
func TestMeWithExpiredToken(t *testing.T) {
	server, ts := newTestServer(t)
	token := signTestToken(t, server, Claims{
		UserID:   1,
		Username: "admin",
		Role:     "admin",
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    TokenIssuer,
			IssuedAt:  jwt.NewNumericDate(time.Now().Add(-2 * time.Hour)),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(-time.Hour)),
		},
	})
	if status := doMe(t, ts, token); status != http.StatusUnauthorized {
		t.Fatalf("过期令牌应返回 401，实际得到 %d", status)
	}
}

// TestClientEndToEnd 验证客户端登录、访问受保护接口以及 401 清除令牌的完整流程
func TestClientEndToEnd(t *testing.T) {
	_, ts := newTestServer(t)
	client := &Client{BaseURL: ts.URL, HTTPClient: ts.Client()}

	loginResp, err := client.Login(context.Background(), "admin", "123456")
	if err != nil {
		t.Fatalf("客户端登录失败：%v", err)
	}
	if loginResp.Token == "" {
		t.Fatal("客户端登录成功但未保存令牌")
	}

	me, err := client.GetMe(context.Background())
	if err != nil {
		t.Fatalf("客户端访问受保护接口失败：%v", err)
	}
	if me.Username != "admin" || me.Role != "admin" {
		t.Fatalf("用户信息不正确：%+v", me)
	}

	// 服务端拒绝无效令牌后，客户端应清除本地令牌并返回 ErrUnauthorized
	client.SetToken("已失效的令牌")
	_, err = client.GetMe(context.Background())
	if !errors.Is(err, ErrUnauthorized) {
		t.Fatalf("期望 ErrUnauthorized，实际得到：%v", err)
	}
	if client.Token() != "" {
		t.Fatal("401 后客户端应清除令牌")
	}
}
