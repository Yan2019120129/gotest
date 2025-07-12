package utils

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"strings"
)

type Http struct {
	header     http.Header
	respHeader http.Header
	params     url.Values
}

func NewHttp() *Http {
	return &Http{header: make(http.Header)}
}

// Set 设置请求头信息
func (h *Http) Set(key, val string) *Http {
	h.header.Set(key, val)
	return h
}

// AddParam 设置get参数
func (h *Http) AddParam(key, val string) *Http {
	if h.params == nil {
		h.params = make(url.Values)
	}
	h.params.Set(key, val)
	return h
}

// GetParam 获取get参数
func (h *Http) GetParam() string {
	return h.params.Encode()
}

// GetRHeader 获取http相应头
func (h *Http) GetRHeader() http.Header {
	return h.respHeader
}

// GetHeader 获取http请求头
func (h *Http) GetHeader() http.Header {
	return h.header
}

// ask 发起请求
func (h *Http) ask(method, url string, body io.Reader) ([]byte, error) {
	request, err := http.NewRequest(method, url, body)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	request.Header = h.header
	do, err := http.DefaultClient.Do(request)
	if err != nil {
		return nil, err
	}
	h.respHeader = do.Header
	return h.result(do.Body)
}

// Get Get请求
func (h *Http) Get(path string) ([]byte, error) {
	if h.params != nil {
		path += "?" + h.params.Encode()
	}
	return h.ask("GET", path, nil)
}

// Post Post请求
func (h *Http) Post(url string, s string) ([]byte, error) {
	return h.PostFormat(url, "application/json", s)
}

// PostFormat Post请求,带格式
func (h *Http) PostFormat(url string, ctxType string, s string) ([]byte, error) {
	h.Set("Content-Type", ctxType)
	var params io.Reader
	if s != "" {
		params = strings.NewReader(s)
	}
	return h.ask("POST", url, params)
}

// PostFile 发起 multipart/form-data 文件上传请求
func (h *Http) PostFile(targetURL, filePath string) ([]byte, error) {
	// 打开文件
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("打开文件失败: %w", err)
	}
	defer file.Close()

	// 创建 multipart 表单体
	var requestBody bytes.Buffer
	writer := multipart.NewWriter(&requestBody)

	// 添加文件字段
	part, err := writer.CreateFormFile("upload", file.Name())
	if err != nil {
		return nil, fmt.Errorf("创建文件字段失败: %w", err)
	}

	// 拷贝文件内容
	if _, err := io.Copy(part, file); err != nil {
		return nil, fmt.Errorf("写入文件内容失败: %w", err)
	}

	// 关闭 multipart writer（自动添加 boundary）
	if err := writer.Close(); err != nil {
		return nil, fmt.Errorf("关闭 multipart writer 失败: %w", err)
	}

	// 设置请求头
	h.Set("Content-Type", writer.FormDataContentType())

	// 发起 POST 请求
	return h.ask("POST", targetURL, &requestBody)
}

func (h *Http) result(body io.ReadCloser) ([]byte, error) {
	defer body.Close()
	data, err := io.ReadAll(body)
	if err != nil {
		return nil, err
	}
	return data, nil
}
