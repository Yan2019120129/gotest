package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"google.golang.org/genai"
)

func main() {
	apiKey := os.Getenv("GEMINI_API_KEY")
	if apiKey == "" {
		fmt.Println("环境变量 GEMINI_API_KEY 未设置")
	} else {
		fmt.Println("环境变量 GEMINI_API_KEY =", apiKey)
	}
	ctx := context.Background()
	client, err := genai.NewClient(ctx, &genai.ClientConfig{
		APIKey:  apiKey,
		Backend: genai.BackendGeminiAPI,
	})
	if err != nil {
		log.Fatal(err)
	}

	imagePath := "/home/yan/Pictures/test/1.png"
	imgData, _ := os.ReadFile(imagePath)

	parts := []*genai.Part{
		genai.NewPartFromText("在鞋子上增加一个品牌吊牌，吊牌颜色为黑色，logo为YIFEIGE，黑色烫金微软雅黑体，尽可能贴近现实，响应图片大小与原图近可能保持一致"),
		&genai.Part{
			InlineData: &genai.Blob{
				MIMEType: "image/png",
				Data:     imgData,
			},
		},
	}

	contents := []*genai.Content{
		genai.NewContentFromParts(parts, genai.RoleUser),
	}

	result, err := client.Models.GenerateContent(
		ctx,
		"gemini-3.1-flash-image-preview",
		contents,
		nil,
	)

	fmt.Println("result:", result, err)
	if err != nil {
		return
	}
	for _, part := range result.Candidates[0].Content.Parts {
		if part.Text != "" {
			fmt.Println(part.Text)
		} else if part.InlineData != nil {
			imageBytes := part.InlineData.Data
			outputFilename := "gemini_generated_image.png"
			_ = os.WriteFile(outputFilename, imageBytes, 0644)
		}
	}
}
