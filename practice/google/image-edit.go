package main

import (
	"context"
	"fmt"
	"net"
	"os"

	"google.golang.org/genai"
)

func main() {
	net.DefaultResolver = &net.Resolver{
		PreferGo: true,
		Dial: func(ctx context.Context, network, address string) (net.Conn, error) {
			d := net.Dialer{}
			return d.DialContext(ctx, "tcp4", address) // 强制 IPv4
		},
	}
	// apiKey := os.Getenv("GEMINI_API_KEY")
	apiKey := "REMOVED_API_KEY"

	ctx := context.Background()

	client, err := genai.NewClient(ctx, &genai.ClientConfig{
		APIKey:  apiKey,
		Backend: genai.BackendGeminiAPI,
	})

	if err != nil {
		panic(err)
	}

	// 读取图片
	imgBytes, err := os.ReadFile("/home/yan/Pictures/test/1.png")
	if err != nil {
		panic(err)
	}

	prompt := `
Add a realistic black hanging brand tag to the shoe.

Requirements:
- tag color: matte black
- logo text: YIFEIGE
- gold foil text
- font similar to Microsoft YaHei
- attached to shoelace
- realistic lighting
`

	// 构造请求
	contents := []*genai.Content{
		{
			Parts: []*genai.Part{
				{
					Text: prompt,
				},
				{
					InlineData: &genai.Blob{
						MIMEType: "image/jpeg",
						Data:     imgBytes,
					},
				},
			},
		},
	}

	resp, err := client.Models.GenerateContent(
		ctx,
		"gemini-2.0-flash",
		contents,
		nil,
	)

	if err != nil {
		panic(err)
	}

	// 解析返回图片
	for _, cand := range resp.Candidates {

		for _, part := range cand.Content.Parts {

			if part.InlineData != nil {

				err := os.WriteFile("output.jpg", part.InlineData.Data, 0644)
				if err != nil {
					panic(err)
				}

				fmt.Println("image saved: output.jpg")
			}
		}
	}
}
