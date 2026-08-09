// Package jwt_t 演示 JWT 令牌在服务端与客户端之间的完整使用流程。
// 服务端负责登录签发 RS256 令牌，并通过中间件校验令牌保护接口；
// 客户端负责登录获取令牌、携带令牌请求受保护接口并处理 401。
package jwt_t

import (
	"context"
	"crypto/rand"
	"crypto/rsa"
	"encoding/json"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const (
	// TokenIssuer 令牌签发者标识
	TokenIssuer = "gotest-jwt-practice"

	// TokenTTL 令牌默认有效期
	TokenTTL = 2 * time.Hour

	// LoginPath 登录接口路径
	LoginPath = "/api/login"

	// MePath 受保护的用户信息接口路径
	MePath = "/api/me"
)

// Claims JWT 自定义载荷
type Claims struct {
	UserID   int64  `json:"uid"`
	Username string `json:"username"`
	Role     string `json:"role"`
	jwt.RegisteredClaims
}

// demoUser 演示用用户信息
type demoUser struct {
	ID       int64
	Password string
	Role     string
}

// Server JWT 演示服务端，持有 RSA 密钥对与演示用户表
type Server struct {
	privateKey *rsa.PrivateKey
	publicKey  *rsa.PublicKey
	users      map[string]demoUser
	ttl        time.Duration
}

// NewServer 创建新的 JWT 演示服务端，并生成 2048 位 RSA 密钥对
func NewServer() (*Server, error) {
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return nil, err
	}
	return &Server{
		privateKey: privateKey,
		publicKey:  &privateKey.PublicKey,
		users: map[string]demoUser{
			"admin": {ID: 1, Password: "123456", Role: "admin"},
			"guest": {ID: 2, Password: "123456", Role: "guest"},
		},
		ttl: TokenTTL,
	}, nil
}

// ServeMux 注册全部路由，登录接口公开，用户信息接口受令牌保护
func (s *Server) ServeMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc(LoginPath, s.handleLogin)
	mux.Handle(MePath, s.authMiddleware(http.HandlerFunc(s.handleMe)))
	return mux
}

// createToken 使用私钥按 RS256 算法签发新令牌，返回令牌与过期时间
func (s *Server) createToken(userID int64, username, role string) (string, time.Time, error) {
	now := time.Now()
	expiresAt := now.Add(s.ttl)
	claims := Claims{
		UserID:   userID,
		Username: username,
		Role:     role,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    TokenIssuer,
			IssuedAt:  jwt.NewNumericDate(now),
			ExpiresAt: jwt.NewNumericDate(expiresAt),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	signed, err := token.SignedString(s.privateKey)
	if err != nil {
		return "", time.Time{}, err
	}
	return signed, expiresAt, nil
}

// verifyToken 使用公钥校验令牌签名与有效期，返回解析出的 claims
func (s *Server) verifyToken(raw string) (*Claims, error) {
	claims := new(Claims)
	_, err := jwt.ParseWithClaims(raw, claims, func(token *jwt.Token) (interface{}, error) {
		// 只接受 RS256 算法，防止算法混淆攻击
		if token.Method != jwt.SigningMethodRS256 {
			return nil, jwt.ErrSignatureInvalid
		}
		return s.publicKey, nil
	})
	if err != nil {
		return nil, err
	}
	return claims, nil
}

// handleLogin 处理登录请求，校验用户名密码后签发令牌
func (s *Server) handleLogin(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		writeJSON(w, http.StatusMethodNotAllowed, map[string]string{"error": "请求方法不支持"})
		return
	}
	var req struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "请求体格式错误"})
		return
	}
	user, ok := s.users[req.Username]
	if !ok || user.Password != req.Password {
		writeJSON(w, http.StatusUnauthorized, map[string]string{"error": "用户名或密码错误"})
		return
	}
	token, expiresAt, err := s.createToken(user.ID, req.Username, user.Role)
	if err != nil {
		writeJSON(w, http.StatusInternalServerError, map[string]string{"error": "签发令牌失败"})
		return
	}
	writeJSON(w, http.StatusOK, map[string]interface{}{
		"token":      token,
		"expires_at": expiresAt,
	})
}

// handleMe 返回当前登录用户的信息，仅允许携带有效令牌访问
func (s *Server) handleMe(w http.ResponseWriter, r *http.Request) {
	claims, ok := claimsFromContext(r.Context())
	if !ok {
		writeJSON(w, http.StatusUnauthorized, map[string]string{"error": "未认证"})
		return
	}
	writeJSON(w, http.StatusOK, map[string]interface{}{
		"uid":       claims.UserID,
		"username":  claims.Username,
		"role":      claims.Role,
		"issuedAt":  claims.IssuedAt.Time,
		"expiresAt": claims.ExpiresAt.Time,
	})
}

// claimsContextKey 请求上下文中存放 claims 的键类型
type claimsContextKey struct{}

// authMiddleware 解析 Authorization 头中的 Bearer 令牌并校验，
// 校验通过后将 claims 注入请求上下文，否则返回 401
func (s *Server) authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		raw := bearerToken(r.Header.Get("Authorization"))
		if raw == "" {
			writeJSON(w, http.StatusUnauthorized, map[string]string{"error": "缺少令牌"})
			return
		}
		claims, err := s.verifyToken(raw)
		if err != nil {
			writeJSON(w, http.StatusUnauthorized, map[string]string{"error": "令牌无效或已过期"})
			return
		}
		ctx := context.WithValue(r.Context(), claimsContextKey{}, claims)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// claimsFromContext 从请求上下文中取出 claims
func claimsFromContext(ctx context.Context) (*Claims, bool) {
	claims, ok := ctx.Value(claimsContextKey{}).(*Claims)
	return claims, ok
}

// bearerToken 从 Authorization 头中提取 Bearer 令牌内容
func bearerToken(header string) string {
	const prefix = "Bearer "
	if strings.HasPrefix(header, prefix) {
		return strings.TrimSpace(strings.TrimPrefix(header, prefix))
	}
	return ""
}

// writeJSON 以 JSON 格式写出响应
func writeJSON(w http.ResponseWriter, status int, v interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(v)
}
