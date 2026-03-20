package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

const baseURL = "http://localhost:8000"

// 通用 HTTP 请求
func doRequest(t *testing.T, method, url string, body any) []byte {
	var reader io.Reader
	if body != nil {
		bs, err := json.Marshal(body)
		require.NoError(t, err)
		reader = bytes.NewReader(bs)
	}

	req, err := http.NewRequest(method, url, reader)
	require.NoError(t, err)

	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	resp, err := http.DefaultClient.Do(req)
	require.NoError(t, err)
	defer resp.Body.Close()

	bs, err := io.ReadAll(resp.Body)
	require.NoError(t, err)

	require.Equal(t, http.StatusOK, resp.StatusCode)
	return bs
}

func Test_WeChat_Login_Flow(t *testing.T) {

	/************ 1️⃣ 创建登录 ************/
	createResp := doRequest(t, "GET", baseURL+"/api/v1/wechat/login-qrcode", nil)

	var createData struct {
		Code int `json:"code"`
		Data struct {
			State string `json:"state"`
		} `json:"data"`
	}
	require.NoError(t, json.Unmarshal(createResp, &createData))
	require.Equal(t, 200, createData.Code)
	require.NotEmpty(t, createData.Data.State)

	state := createData.Data.State
	t.Log("state:", state)

	/************ 2️⃣ 登录回调（模拟微信） ************/
	callbackURL := baseURL + "/api/v1/wechat/login-callback" +
		"?code=TEST_CODE&state=" + state

	_ = doRequest(t, "GET", callbackURL, nil)

	/************ 3️⃣ 状态检测（轮询） ************/
	var appCode string

	for i := 0; i < 5; i++ {
		stateResp := doRequest(
			t,
			"GET",
			baseURL+"/api/v1/wechat/login-state?code=ANY&state="+state,
			nil,
		)

		var stateData struct {
			Code int `json:"code"`
			Data struct {
				Success bool   `json:"success"`
				Code    string `json:"code"`
			} `json:"data"`
		}

		require.NoError(t, json.Unmarshal(stateResp, &stateData))
		require.Equal(t, 200, stateData.Code)

		if stateData.Data.Success {
			appCode = stateData.Data.Code
			break
		}

		time.Sleep(500 * time.Millisecond)
	}

	require.NotEmpty(t, appCode)
	t.Log("appCode:", appCode)

	/************ 4️⃣ 登录 ************/
	loginResp := doRequest(
		t,
		"POST",
		baseURL+"/api/v1/login",
		map[string]string{
			"appCode": appCode,
		},
	)

	var loginData map[string]any
	require.NoError(t, json.Unmarshal(loginResp, &loginData))
	//require.Equal(t, 200, loginData["code"])
	fmt.Println(loginData)
}

// TestServer 服务端：
//
//	:5000  控制通道 —— A 连入注册
//	:50001 数据通道 —— B 连入后通知 A 做转发
func TestServer(t *testing.T) {

}

// TestA 客户端 A：
//  1. 连接服务端 :5000 注册
//  2. 阻塞等待通知
//  3. 收到通知后连本地 :25565，再连服务端 :50001，双向转发
func TestA(t *testing.T) {

}

// TestB 客户端 B：连接服务端 :50001，发数据，等 echo
func TestB(t *testing.T) {
	time.Sleep(500 * time.Millisecond) // 等 Server/A 就绪

	conn, err := net.Dial("tcp", "127.0.0.1:50001")
	if err != nil {
		t.Fatal("连接服务端 :50001 失败:", err)
	}
	defer conn.Close()
	fmt.Println("[B] 已连入服务端 :50001")

	// 发送测试数据（本地 25565 需是 echo 服务才能收到回显）
	payload := "Hello from B!\n"
	_, err = conn.Write([]byte(payload))
	if err != nil {
		t.Fatal("发送失败:", err)
	}
	fmt.Println("[B] 已发送:", payload)

	buf := make([]byte, 64)
	n, err := conn.Read(buf)
	if err != nil {
		t.Fatal("读取回显失败:", err)
	}
	fmt.Println("[B] 收到回显:", string(buf[:n]))
}
