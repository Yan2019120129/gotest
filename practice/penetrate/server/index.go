package main

import (
	"flag"
	"fmt"
	"io"
	"net"
)

func main() {
	aPort := flag.String("a", ":1041", "A 预连接监听端口")
	bPort := flag.String("b", ":1040", "B 公开监听端口")
	poolSize := flag.Int("pool", 5, "连接池数量")
	flag.Parse()

	aLn, err := net.Listen("tcp", *aPort)
	if err != nil {
		fmt.Println("监听失败:", err)
		return
	}
	defer aLn.Close()

	bLn, err := net.Listen("tcp", *bPort)
	if err != nil {
		fmt.Println("监听失败:", err)
		return
	}
	defer bLn.Close()

	fmt.Printf("[Server] A 预连接端口 %s | B 公开端口 %s | 连接池 %d\n", *aPort, *bPort, *poolSize)

	pool := make(chan net.Conn, (*poolSize)*2)

	// 持续接收 A 的预连接
	go func() {
		for {
			conn, err := aLn.Accept()
			if err != nil {
				return
			}
			fmt.Printf("[Server] 池 +1，当前 %d\n", len(pool)+1)
			pool <- conn
		}
	}()

	// B 来了，从池子取一条 A 的连接，发信号后打通
	for {
		bConn, err := bLn.Accept()
		if err != nil {
			return
		}
		fmt.Println("[Server] B 连入:", bConn.RemoteAddr())

		aConn := <-pool
		fmt.Printf("[Server] 配对，池剩 %d\n", len(pool))

		// 通知 A：你被配对了，去连本地服务
		aConn.Write([]byte{0x01})

		go bridge(bConn, aConn)
	}
}

func bridge(b, a net.Conn) {
	defer b.Close()
	defer a.Close()
	done := make(chan struct{}, 2)
	go func() { io.Copy(a, b); done <- struct{}{} }()
	go func() { io.Copy(b, a); done <- struct{}{} }()
	<-done
	fmt.Printf("[Server] 转发结束 B=%s\n", b.RemoteAddr())
}
