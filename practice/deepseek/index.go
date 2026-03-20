package main

import (
	"fmt"
	"gotest/common/utils"
	"os"
)

func main() {
	apiKey := os.Getenv("deepseek_api_key")
	h := utils.NewHttp()
	h.Set("Content-Type", "application/json")
	h.Set("Authorization", "Bearer "+apiKey)
	// 转换为结构体
	type DeepSeekRequest struct {
		Model    string `json:"model"`
		Messages []struct {
			Role    string `json:"role"`
			Content string `json:"content"`
		} `json:"messages"`
		Stream bool `json:"stream"`
	}
	request := DeepSeekRequest{
		Model: "deepseek-chat",
		Messages: []struct {
			Role    string `json:"role"`
			Content string `json:"content"`
		}{
			{Role: "system", Content: "You are a helpful assistant."},
			{Role: "user", Content: "Hello!"},
		},
		Stream: false,
	}

	v, err := h.Post("https://api.deepseek.com/chat/completions", utils.ObjToString(request))
	fmt.Println(string(v), err)
}
