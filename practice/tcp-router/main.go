package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"os"
	"strings"
	"sync"
)

// ================= 数据结构 =================

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Client struct {
	IP string `json:"ip"`
}

type ServerConfig struct {
	WebPort int    `json:"web_port"` // Web端口
	TCPPort int    `json:"tcp_port"` // TCP监听端口
	Target  string `json:"target"`   // 转发地址 127.0.0.1:25565
}

type Config struct {
	Server  ServerConfig `json:"server"`
	Users   []User       `json:"users"`
	Clients []Client     `json:"clients"`
}

var config Config
var mu sync.Mutex
var logger *log.Logger

// ================= 初始化 =================

func initLogger() {
	file, _ := os.OpenFile("app.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	logger = log.New(io.MultiWriter(os.Stdout, file), "", log.LstdFlags)
}

// ================= 配置读写 =================

func loadConfig() {
	file, err := os.ReadFile("config.json")
	if err != nil {
		// 默认配置
		config = Config{
			Server: ServerConfig{
				WebPort: 8080,
				TCPPort: 1040,
				Target:  "127.0.0.1:25565",
			},
			Users: []User{
				{Username: "admin", Password: "123456"},
			},
		}
		saveConfig()
		return
	}

	json.Unmarshal(file, &config)
}

func saveConfig() {
	mu.Lock()
	defer mu.Unlock()

	data, _ := json.MarshalIndent(config, "", "  ")
	os.WriteFile("config.json", data, 0644)
}

// ================= 工具函数 =================

func getClientIP(r *http.Request) string {
	ip := r.Header.Get("X-Real-IP")
	if ip == "" {
		ip = r.Header.Get("X-Forwarded-For")
	}
	if ip == "" {
		ip = strings.Split(r.RemoteAddr, ":")[0]
	}
	return ip
}

func getTCPClientIP(addr string) string {
	return strings.Split(addr, ":")[0]
}

func isIPAllowed(ip string) bool {
	for _, c := range config.Clients {
		if c.IP == ip {
			return true
		}
	}
	return false
}

// ================= Web中间件 =================

func logMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ip := getClientIP(r)
		logger.Printf("[WEB] %s %s IP=%s", r.Method, r.URL.Path, ip)
		next.ServeHTTP(w, r)
	})
}

func writeHTML(w http.ResponseWriter, html string) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Write([]byte(html))
}

// ================= Web部分 =================

// 登录页面
func loginPage(w http.ResponseWriter, r *http.Request) {
	html := `
	<h2>登录</h2>
	<form method="POST" action="/login">
	账号: <input name="username"><br>
	密码: <input name="password" type="password"><br>
	<button type="submit">登录</button>
	</form>
	`
	writeHTML(w, html)
}

// 登录处理
func loginHandler(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	password := r.FormValue("password")
	ip := getClientIP(r)

	for _, u := range config.Users {
		if u.Username == username && u.Password == password {
			logger.Printf("[LOGIN SUCCESS] user=%s ip=%s", username, ip)
			http.Redirect(w, r, "/auth", 302)
			return
		}
	}

	logger.Printf("[LOGIN FAIL] user=%s ip=%s", username, ip)
	w.Write([]byte("登录失败"))
}

// 授权页面
func authPage(w http.ResponseWriter, r *http.Request) {
	ip := getClientIP(r)

	html := fmt.Sprintf(`
	<h2>授权页面</h2>
	<p>你的IP: %s</p>
	<form method="POST" action="/authorize">
	<button type="submit">授权访问</button>
	</form>
	`, ip)

	writeHTML(w, html)
}

// 授权接口
func authorizeHandler(w http.ResponseWriter, r *http.Request) {
	ip := getClientIP(r)

	for _, c := range config.Clients {
		if c.IP == ip {
			logger.Printf("[AUTH] 已存在 ip=%s", ip)
			w.Write([]byte("已授权过"))
			return
		}
	}

	config.Clients = append(config.Clients, Client{IP: ip})
	saveConfig()

	logger.Printf("[AUTH SUCCESS] ip=%s", ip)
	w.Write([]byte("授权成功: " + ip))
}

// ================= TCP代理 =================

func startTCPProxy() {
	addr := fmt.Sprintf(":%d", config.Server.TCPPort)

	ln, err := net.Listen("tcp", addr)
	if err != nil {
		panic(err)
	}

	logger.Printf("[SYSTEM] TCP监听端口 %d", config.Server.TCPPort)

	for {
		conn, err := ln.Accept()
		if err != nil {
			logger.Println("[ERROR] accept失败:", err)
			continue
		}
		go handleConn(conn)
	}
}

func handleConn(client net.Conn) {
	defer client.Close()

	ip := getTCPClientIP(client.RemoteAddr().String())
	logger.Printf("[TCP CONNECT] ip=%s", ip)

	if !isIPAllowed(ip) {
		logger.Printf("[TCP REJECT] ip=%s 未授权", ip)
		return
	}

	targetAddr := config.Server.Target
	logger.Printf("[TCP ALLOW] ip=%s -> %s", ip, targetAddr)

	target, err := net.Dial("tcp", targetAddr)
	if err != nil {
		logger.Printf("[TCP ERROR] 连接目标失败 err=%v", err)
		return
	}
	defer target.Close()

	go io.Copy(target, client)
	io.Copy(client, target)

	logger.Printf("[TCP CLOSE] ip=%s", ip)
}

// ================= 主函数 =================

func main() {
	initLogger()
	loadConfig()

	// Web路由
	http.HandleFunc("/", loginPage)
	http.HandleFunc("/login", loginHandler)
	http.HandleFunc("/auth", authPage)
	http.HandleFunc("/authorize", authorizeHandler)

	// 启动Web
	go func() {
		addr := fmt.Sprintf(":%d", config.Server.WebPort)
		logger.Printf("[SYSTEM] Web启动端口 %d", config.Server.WebPort)
		http.ListenAndServe(addr, logMiddleware(http.DefaultServeMux))
	}()

	// 启动TCP
	startTCPProxy()
}
