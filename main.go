package main

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

const notifyURL = "http://020200.xyz:1077/pay/notify"

func main() {
	// client 用于向支付宝异步通知路由发送 HTTP 请求。
	client := &http.Client{Timeout: 10 * time.Second}
	response, err := client.Post(notifyURL, "application/x-www-form-urlencoded", nil)
	if err != nil {
		fmt.Println("发送异步通知请求失败：", err)
		return
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("读取异步通知响应失败：", err)
		return
	}
	fmt.Printf("异步通知响应：状态码=%d，内容=%s\n", response.StatusCode, body)
}
