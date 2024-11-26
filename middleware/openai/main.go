package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type Resul struct {
	Choices []struct {
		FinishReason string `json:"finish_reason"`
		Index        int    `json:"index"`
		Message      struct {
			Content string `json:"content"`
			Role    string `json:"role"`
		}
		Logprobs any `json:"logprobs"`
	}
	Created int    `json:"created"`
	Id      string `json:"id"`
	Model   string `json:"model"`
	Object  string `json:"object"`
	Usage   struct {
		CompletionTokens int `json:"completion_tokens"`
		PromptTokens     int `json:"prompt_tokens"`
		TotalTokens      int `json:"total_tokens"`
	}
}

func main() {
	url := "https://api.openai.com/v1/chat/completions"

	messages := []map[string]string{
		{"role": "user", "content": "给我介绍一下中秋节"},
	}

	bodyMap := map[string]interface{}{
		"model":    "gpt-4o",
		"messages": messages,
	}

	body, _ := json.Marshal(bodyMap)

	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(body))
	req.Header.Add("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	defer resp.Body.Close()
	var result Resul
	json.NewDecoder(resp.Body).Decode(&result)
	for _, choice := range result.Choices {
		fmt.Println(choice.Index, "-----", choice.Message)
	}
}
