package net_t

import (
	"fmt"
	"net"
	"testing"
)

var connTarget net.Conn

func TestServer(t *testing.T) {
	listen, err := net.Listen("tcp", ":1070")
	if err != nil {
		fmt.Println("listen err", err)
		return
	}
	defer listen.Close()
	for {
		accept, err := listen.Accept()
		if err != nil {
			fmt.Println("accept err", err)
			return
		}
		connTarget = accept
		fmt.Println("add target", connTarget)
	}
}

func TestServer1(t *testing.T) {
	listen, err := net.Listen("tcp", ":1071")
	if err != nil {
		fmt.Println("listen err", err)
		return
	}
	defer listen.Close()
	for {
		accept, err := listen.Accept()
		if err != nil {
			fmt.Println("accept err", err)
			return
		}
		fmt.Println("link target", connTarget, accept)
		go Copy(connTarget, accept)
	}
}
