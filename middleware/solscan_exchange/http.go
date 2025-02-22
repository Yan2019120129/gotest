package main

import (
	"io"
	"log"
	"net/http"
	"strings"
)

type Http struct {
	header http.Header
}

func NewHttp() *Http {
	return &Http{header: make(http.Header)}
}

// Set 设置请求头信息
func (h *Http) Set(key, val string) *Http {
	h.header.Set(key, val)
	return h
}

// ask 发起请求
func (h *Http) ask(method, url string, body io.Reader) []byte {
	request, err := http.NewRequest(method, url, body)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	request.Header = h.header
	do, err := http.DefaultClient.Do(request)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	return h.result(do.Body)
}

// Get Get请求
func (h *Http) Get(url string) []byte {
	return h.ask("GET", url, nil)
}

// Post Post请求
func (h *Http) Post(url string, ctxType string, s string) []byte {
	h.Set("Content-Type", ctxType)
	var params io.Reader
	if s != "" {
		params = strings.NewReader(s)
	}
	return h.ask("POST", url, params)
}

func (h *Http) result(body io.ReadCloser) []byte {
	defer body.Close()
	data, err := io.ReadAll(body)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	return data
}
