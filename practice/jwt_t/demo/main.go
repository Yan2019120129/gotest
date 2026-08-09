// Command demo 演示 JWT 令牌在服务端与客户端之间的完整使用流程：
// 登录签发令牌 -> 携带令牌访问受保护接口 -> 未登录访问被拒绝 ->
// 解析并篡改令牌用户内容后重新封装，仍被服务端签名校验拒绝。
package main

import (
	"context"
	"errors"
	"fmt"
	"net"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"

	"gotest/practice/jwt_t"
)

func main() {
	server, err := jwt_t.NewServer()
	if err != nil {
		fmt.Println("初始化服务端失败：", err)
		os.Exit(1)
	}

	addr := os.Getenv("JWT_DEMO_ADDR")
	if addr == "" {
		addr = "127.0.0.1:18080"
	}
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		fmt.Println("监听端口失败：", err)
		os.Exit(1)
	}

	httpServer := &http.Server{Handler: server.ServeMux()}
	go func() {
		// 服务端在后台 goroutine 中运行，演示结束后关闭
		_ = httpServer.Serve(listener)
	}()

	fmt.Printf("== JWT 使用流程演示 ==\n服务端已启动：http://%s\n\n", listener.Addr())
	ctx := context.Background()
	client := &jwt_t.Client{BaseURL: "http://" + listener.Addr().String()}

	// 流程1：客户端登录，服务端校验用户名密码并签发 RS256 令牌
	loginResp, err := client.Login(ctx, "admin", "123456")
	if err != nil {
		fmt.Println("登录失败：", err)
		os.Exit(1)
	}
	token := loginResp.Token
	if len(token) > 60 {
		token = token[:40] + "..." + token[len(token)-20:]
	}
	fmt.Printf("[1] 登录成功，服务端签发令牌（有效期至 %s）\n", loginResp.ExpiresAt.Format("2006-01-02 15:04:05"))
	fmt.Printf("    令牌：%s\n\n", token)

	// 流程2：客户端携带 Bearer 令牌访问受保护接口
	me, err := client.GetMe(ctx)
	if err != nil {
		fmt.Println("访问受保护接口失败：", err)
		os.Exit(1)
	}
	fmt.Printf("[2] 携带令牌访问 %s 成功，当前用户：%s（角色：%s）\n\n", jwt_t.MePath, me.Username, me.Role)

	// 流程3：客户端登出清除令牌，再访问受保护接口应被拒绝
	client.SetToken("")
	_, err = client.GetMe(ctx)
	if errors.Is(err, jwt_t.ErrUnauthorized) {
		fmt.Println("[3] 未登录访问受保护接口：401 未授权，客户端令牌已清除")
	} else {
		fmt.Println("[3] 未登录访问结果异常：", err)
	}

	// 流程4：解析真实令牌并篡改其中的用户内容，沿用原签名重新封装后访问应被拒绝
	forged, err := forgeToken(loginResp.Token, func(claims *jwt_t.Claims) {
		claims.Username = "root"
		claims.Role = "superadmin"
	})
	if err != nil {
		fmt.Println("篡改令牌失败：", err)
		os.Exit(1)
	}
	fmt.Println("[4] 已解析原令牌并篡改用户内容：username admin -> root，role admin -> superadmin")
	client.SetToken(forged)
	_, err = client.GetMe(ctx)
	if errors.Is(err, jwt_t.ErrUnauthorized) {
		fmt.Println("    重新封装的令牌访问受保护接口：401 未授权，服务端通过签名校验识破了篡改")
	} else {
		fmt.Println("    重新封装的令牌访问结果异常：", err)
	}

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	_ = httpServer.Shutdown(shutdownCtx)
	fmt.Println("\n演示结束，服务端已关闭")
}

// forgeToken 解析原始令牌的载荷，经 mutate 篡改用户内容后重新封装令牌。
// 攻击者不持有服务端私钥，只能沿用原签名，因此篡改后的令牌无法通过签名校验。
func forgeToken(raw string, mutate func(*jwt_t.Claims)) (string, error) {
	claims := new(jwt_t.Claims)
	// 仅解析不校验签名，读取原令牌中的用户内容
	_, parts, err := jwt.NewParser().ParseUnverified(raw, claims)
	if err != nil {
		return "", err
	}
	mutate(claims)
	// 用篡改后的 claims 重建头部与载荷，拼接原签名完成重新封装
	rebuilt, err := jwt.NewWithClaims(jwt.SigningMethodRS256, claims).SigningString()
	if err != nil {
		return "", err
	}
	return rebuilt + "." + parts[2], nil
}
