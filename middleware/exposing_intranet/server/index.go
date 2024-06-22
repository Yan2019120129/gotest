package server

import (
	"io"
	"log"
	"net"
)

const (
	publicAddr   = "0.0.0.0:4000"   // 公网服务器监听的地址和端口
	internalAddr = "localhost:9090" // 内网服务器的地址和端口
)

func server() {
	ln, err := net.Listen("tcp", publicAddr)
	if err != nil {
		log.Fatalf("Failed to bind address: %v", err)
	}
	log.Printf("Public server listening on %s", publicAddr)

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Printf("Failed to accept connection: %v", err)
			continue
		}

		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	internalConn, err := net.Dial("tcp", internalAddr)
	if err != nil {
		log.Printf("Failed to connect to internal server: %v", err)
		return
	}
	defer internalConn.Close()

	go io.Copy(internalConn, conn)
	io.Copy(conn, internalConn)
}
