package main

import (
	"flag"
	"fmt"
	"io"
	"net"
	"time"
)

func main() {
	var (
		serverAddr string
		localAddr  string
	)

	flag.StringVar(&serverAddr, "s", "8.138.57.34:1041", "服务端地址端口")
	flag.StringVar(&localAddr, "l", "127.0.0.1:25565", "本地映射地址端口")
	flag.Parse()

	fmt.Printf("[A] 启动 | 服务端=%s 本地=%s\n", serverAddr, localAddr)
	go addToPool(serverAddr, localAddr)
	select {}
}

func addToPool(serverAddr, localAddr string) {
	for {
		serverConn, err := net.Dial("tcp", serverAddr)
		if err != nil {
			fmt.Println("[A] 连服务端失败:", err, "3s 后重试")
			time.Sleep(3 * time.Second)
			continue
		}
		fmt.Println("[A] 预连接就绪，等待配对")

		// 阻塞等待 Server 的配对信号
		buf := make([]byte, 1)
		if _, err = serverConn.Read(buf); err != nil {
			fmt.Println("[A] 等待配对失败:", err)
			serverConn.Close()
			continue
		}

		// 被配对了，立即补充一条新的预连接保证池子不空
		go addToPool(serverAddr, localAddr)
		fmt.Println("[A] 已配对，连接本地服务")

		// 连本地目标服务
		localConn, err := net.Dial("tcp", localAddr)
		if err != nil {
			fmt.Println("[A] 连本地服务失败:", err)
			serverConn.Close()
			return
		}

		// 双向转发
		forward(serverConn, localConn)
		return
	}
}

func forward(serverConn, localConn net.Conn) {
	defer serverConn.Close()
	defer localConn.Close()
	done := make(chan struct{}, 2)
	go func() { io.Copy(localConn, serverConn); done <- struct{}{} }()
	go func() { io.Copy(serverConn, localConn); done <- struct{}{} }()
	<-done
	fmt.Println("[A] 一条转发结束")
}
