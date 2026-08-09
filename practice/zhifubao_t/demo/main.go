// Command demo 演示支付宝扫码支付完整流程：扫码下单 -> 用户扫码付款 -> 异步通知 -> 查询 -> 退款。
// 运行前请先在 practice/zhifubao_t/config.yaml 中填入沙箱应用的真实密钥。
package main

import (
	"fmt"
	"net"
	"net/http"
	"os"

	"gotest/practice/zhifubao_t"
)

func main() {
	configPath := os.Getenv("ALIPAY_CONFIG")
	if configPath == "" {
		configPath = "practice/zhifubao_t/config.yaml"
	}

	config, err := zhifubao_t.LoadConfig(configPath)
	if err != nil {
		fmt.Println("加载配置失败：", err)
		os.Exit(1)
	}
	client, err := zhifubao_t.NewClient(config)
	if err != nil {
		fmt.Println("初始化支付宝客户端失败（请确认 config.yaml 已填入沙箱应用私钥与支付宝公钥）：", err)
		os.Exit(1)
	}

	server := zhifubao_t.NewServer(client, config)

	addr := os.Getenv("ALIPAY_DEMO_ADDR")
	if addr == "" {
		addr = "127.0.0.1:18081"
	}
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		fmt.Println("监听端口失败：", err)
		os.Exit(1)
	}

	fmt.Println("== 支付宝支付流程演示 ==")
	fmt.Printf("服务地址：http://%s\n", listener.Addr())
	fmt.Println("扫码下单：/pay?amount=0.01&subject=测试商品")
	fmt.Println("异步通知：/pay/notify")
	fmt.Println("查询：/pay/query?out_trade_no=xxx    退款：POST /pay/refund")
	fmt.Println("注意：notify_url 需为支付宝可访问的公网地址，本地演示请使用内网穿透工具")
	_ = http.Serve(listener, server.Handler())
}
