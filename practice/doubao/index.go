package main

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

// 请求结构体
type ArkRequest struct {
	Model                            string         `json:"model"`
	Prompt                           string         `json:"prompt"`
	Image                            []string       `json:"image,omitempty"` // URL 或 base64
	SequentialImageGeneration        string         `json:"sequential_image_generation,omitempty"`
	SequentialImageGenerationOptions map[string]int `json:"sequential_image_generation_options,omitempty"`
	Size                             string         `json:"size,omitempty"`
	OutputFormat                     string         `json:"output_format,omitempty"`
	Watermark                        bool           `json:"watermark,omitempty"`
}

// 响应结构体
type ArkResponse struct {
	Data []struct {
		URL  string `json:"url"`
		Size string `json:"size"`
	} `json:"data"`
}

func main() {
	prompt := flag.String("prompt", "", "生成图片的文本描述")
	dir := flag.String("dir", ".", "保存图片的目录（默认当前目录）")
	images := flag.String("images", "", "图片路径，多个用逗号分隔，本地路径或URL都支持")
	flag.Parse()

	if *prompt == "" {
		fmt.Println("请提供 -prompt 参数")
		return
	}

	var imageList []string
	if *images != "" {
		paths := splitAndTrim(*images)
		for _, p := range paths {
			if fileExists(p) {
				// 本地文件，转base64
				b64, err := fileToBase64(p)
				if err != nil {
					fmt.Println("读取本地图片失败:", err)
					continue
				}
				// 注意 base64 前缀，很多接口要求 "data:image/png;base64,xxxx"
				imageList = append(imageList, "data:image/png;base64,"+b64)
			} else {
				// 当做URL
				imageList = append(imageList, p)
			}
		}
	}

	reqBody := ArkRequest{
		// Model:                     "doubao-seedream-4-5-251128", // 这里用你的model
		Model:                     "doubao-seedream-4-0-250828", // 这里用你的model
		Prompt:                    *prompt,
		Image:                     imageList,
		SequentialImageGeneration: "auto",
		SequentialImageGenerationOptions: map[string]int{
			"max_images": 5,
		},
		// Size: "2K",
		Watermark: false,
	}

	data, err := json.Marshal(reqBody)
	if err != nil {
		fmt.Println("JSON序列化失败:", err)
		return
	}

	apiKey := os.Getenv("ARK_API_KEY")
	if apiKey == "" {
		fmt.Println("请设置环境变量 ARK_API_KEY")
		return
	}

	client := &http.Client{Timeout: 120 * time.Second}
	req, err := http.NewRequest("POST", "https://ark.cn-beijing.volces.com/api/v3/images/generations", bytes.NewReader(data))
	if err != nil {
		fmt.Println("创建请求失败:", err)
		return
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+apiKey)

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("请求失败:", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		fmt.Println("请求错误:", resp.Status, string(body))
		return
	}

	var arkResp ArkResponse
	if err := json.NewDecoder(resp.Body).Decode(&arkResp); err != nil {
		fmt.Println("解析响应失败:", err)
		return
	}

	for i, item := range arkResp.Data {
		if item.URL == "" {
			continue
		}
		err := downloadImage(item.URL, *dir, fmt.Sprintf("image_%d.png", i+1))
		if err != nil {
			fmt.Println("下载图片失败:", err)
		} else {
			fmt.Println("图片保存成功:", filepath.Join(*dir, fmt.Sprintf("image_%d.png", i+1)))
		}
	}
}

func downloadImage(url, dir, filename string) error {
	// 如果目录不存在，创建
	if err := os.MkdirAll(dir, 0755); err != nil {
		return fmt.Errorf("创建目录失败: %v", err)
	}

	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("下载失败: %s", resp.Status)
	}

	path := filepath.Join(dir, filename)
	out, err := os.Create(path)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	return err
}

// 本地文件转base64
func fileToBase64(path string) (string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(data), nil
}

// 检查文件是否存在
func fileExists(path string) bool {
	info, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

// 辅助函数：拆分逗号并去掉空格
func splitAndTrim(s string) []string {
	var res []string
	for _, part := range bytes.Split([]byte(s), []byte(",")) {
		p := string(bytes.TrimSpace(part))
		if p != "" {
			res = append(res, p)
		}
	}
	return res
}
