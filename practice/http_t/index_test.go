package http_t

import (
	"testing"
)

// TestHttp 测试http post 请求获取Eh数据
func TestHttp(t *testing.T) {
	_, err := Post("USDCNY", "1m")
	if err != nil {
		return
	}
}
