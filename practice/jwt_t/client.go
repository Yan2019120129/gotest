package jwt_t

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"
)

// ErrUnauthorized 表示令牌缺失、无效或已过期
var ErrUnauthorized = errors.New("未授权：令牌无效或已过期")

// LoginResponse 登录接口的返回内容
type LoginResponse struct {
	Token     string    `json:"token"`
	ExpiresAt time.Time `json:"expires_at"`
}

// MeResponse 受保护的用户信息接口的返回内容
type MeResponse struct {
	UserID    int64     `json:"uid"`
	Username  string    `json:"username"`
	Role      string    `json:"role"`
	IssuedAt  time.Time `json:"issuedAt"`
	ExpiresAt time.Time `json:"expiresAt"`
}

// apiError 服务端返回的错误信息结构
type apiError struct {
	Error string `json:"error"`
}

// Client JWT 演示客户端，负责登录获取令牌并携带令牌请求受保护接口
type Client struct {
	BaseURL    string
	HTTPClient *http.Client
	token      string
}

// Login 使用用户名密码登录，成功后保存服务端签发的令牌
func (c *Client) Login(ctx context.Context, username, password string) (LoginResponse, error) {
	var resp LoginResponse
	body, err := json.Marshal(map[string]string{
		"username": username,
		"password": password,
	})
	if err != nil {
		return resp, err
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, c.BaseURL+LoginPath, bytes.NewReader(body))
	if err != nil {
		return resp, err
	}
	req.Header.Set("Content-Type", "application/json")
	if err := c.do(req, &resp); err != nil {
		return resp, err
	}
	c.token = resp.Token
	return resp, nil
}

// GetMe 携带当前令牌访问受保护的用户信息接口，
// 服务端返回 401 时清除本地令牌并返回 ErrUnauthorized
func (c *Client) GetMe(ctx context.Context) (MeResponse, error) {
	var resp MeResponse
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, c.BaseURL+MePath, nil)
	if err != nil {
		return resp, err
	}
	if c.token != "" {
		req.Header.Set("Authorization", "Bearer "+c.token)
	}
	if err := c.do(req, &resp); err != nil {
		return resp, err
	}
	return resp, nil
}

// SetToken 手动设置令牌，用于登出后清除或演示伪造令牌场景
func (c *Client) SetToken(token string) {
	c.token = token
}

// Token 返回当前保存的令牌
func (c *Client) Token() string {
	return c.token
}

// do 发送请求并处理响应，401 时清除令牌并返回 ErrUnauthorized，
// 其他非 2xx 状态码返回服务端错误信息
func (c *Client) do(req *http.Request, out interface{}) error {
	httpClient := c.HTTPClient
	if httpClient == nil {
		httpClient = http.DefaultClient
	}
	resp, err := httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusUnauthorized {
		c.token = ""
		return ErrUnauthorized
	}
	if resp.StatusCode < http.StatusOK || resp.StatusCode >= http.StatusMultipleChoices {
		var apiErr apiError
		if err := json.NewDecoder(resp.Body).Decode(&apiErr); err == nil && apiErr.Error != "" {
			return errors.New(apiErr.Error)
		}
		return fmt.Errorf("请求失败：%s", resp.Status)
	}
	if out != nil {
		if err := json.NewDecoder(resp.Body).Decode(out); err != nil && err != io.EOF {
			return err
		}
	}
	return nil
}
