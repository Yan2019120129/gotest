package net_t

import (
	"fmt"
	"io"
	"net"
	"testing"
)

type Msg struct {
	Val string
}

func TestClient(t *testing.T) {
	conn, err := net.Dial("tcp", ":1071")
	if err != nil {
		fmt.Println("dial err", err)
		return
	}
	//encoder := json.NewEncoder(conn)
	//err = encoder.Encode(&Msg{Val: "ok"})
	//if err != nil {
	//	fmt.Println("Write err", err)
	//	return
	//}

	targetConn, err := net.Dial("tcp", ":3306")
	if err != nil {
		fmt.Println("dial 1 err", err)
		return
	}
	Copy(conn, targetConn)
}

func Copy(conn, targetConn net.Conn) {
	defer conn.Close()
	defer targetConn.Close()
	go func() {
		_, err := io.Copy(conn, targetConn)
		if err != nil {
			fmt.Println("copy err", err)
			return
		}
	}()
	_, err := io.Copy(targetConn, conn)
	if err != nil {
		fmt.Println("copy 1 err", err)
		return
	}
}
