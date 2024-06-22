package main

import (
	"io"
	"log"
	"net"
)

const (
	publicAddr = "8.138.57.34:4000" // 公网服务器的地址和端口
)

func client() {
	conn, err := net.Dial("tcp", publicAddr)
	if err != nil {
		log.Fatalf("Failed to connect to public server: %v", err)
	}
	log.Printf("Connected to public server at %s", publicAddr)

	localListener, err := net.Listen("tcp", ":9090")
	if err != nil {
		log.Fatalf("Failed to start local listener: %v", err)
	}
	log.Printf("Local server listening on %s", localListener.Addr())

	go handlePublicConnection(conn)

	for {
		localConn, err := localListener.Accept()
		if err != nil {
			log.Printf("Failed to accept local connection: %v", err)
			continue
		}

		go handleLocalConnection(localConn, conn)
	}
}

func handlePublicConnection(conn net.Conn) {
	defer conn.Close()
	io.Copy(io.Discard, conn)
}

func handleLocalConnection(localConn, publicConn net.Conn) {
	defer localConn.Close()
	go io.Copy(publicConn, localConn)
	io.Copy(localConn, publicConn)
}
