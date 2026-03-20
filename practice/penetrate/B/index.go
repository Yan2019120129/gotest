package main

import (
	"fmt"
	"io"
	"net"
	"os"
	"time"
)

// B 完全透明，不需要发任何标识，连上就直接收发数据
// 就像直连 25565 一样使用
func main() {
	time.Sleep(300 * time.Millisecond)

	conn, err := net.Dial("tcp", "127.0.0.1:50001")
	if err != nil {
		fmt.Println("[B] 连接失败:", err)
		os.Exit(1)
	}
	defer conn.Close()
	fmt.Printf("[B] 已连入 :50001 (pid=%d)\n", os.Getpid())

	// 并发读：把服务端（即 25565）返回的数据打印出来
	go func() {
		io.Copy(os.Stdout, conn)
	}()

	// 发数据（25565 是什么服务就发什么，这里用文本模拟）
	payload := fmt.Sprintf("Hello from B pid=%d\n", os.Getpid())
	conn.Write([]byte(payload))
	fmt.Print("[B] 已发送:", payload)

	time.Sleep(20 * time.Second)
	fmt.Println("[B] 结束")
}
