package dto

import "net/http"

// Massage 返回类型。
type Massage struct {
	Code int         `json:"code"` // 返回状态码
	Msg  string      `json:"msg"`  // 返回信息
	Data interface{} `json:"data"` // 返回数据
}

// Success 返回成功信息
func Success(data interface{}) (int, interface{}) {
	return http.StatusOK, data
}

// Error 返回错处信息
func Error(err error) (int, interface{}) {
	return http.StatusInternalServerError, err.Error()
}
