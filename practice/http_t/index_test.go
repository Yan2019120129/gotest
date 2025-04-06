package http_t

import (
	"fmt"
	"github.com/cheggaaa/pb/v3"
	"io"
	"net/http"
	"os"
	"testing"
	"time"
)

// TestHttp 测试http post 请求获取Eh数据
func TestHttp(t *testing.T) {
	_, err := Post("USDCNY", "1m")
	if err != nil {
		return
	}
}

// 测试http 下载安装
func TestHttpDownload(t *testing.T) {
	url := "https://releases.ubuntu.com/24.04/ubuntu-24.04.2-desktop-amd64.iso"
	filePath := "ubuntu-24.04.2-desktop-amd64.iso"

	// 1. 检查本地文件状态
	file, fileSize, err := checkLocalFile(filePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// 2. 创建HTTP客户端
	client := &http.Client{
		Timeout: 30 * time.Minute,
	}

	// 3. 创建支持断点的请求
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(fmt.Sprintf("创建请求失败: %v", err))
	}

	// 设置Range头
	if fileSize > 0 {
		req.Header.Set("Range", fmt.Sprintf("bytes=%d-", fileSize))
	}

	// 4. 发送请求
	resp, err := client.Do(req)
	if err != nil {
		panic(fmt.Sprintf("请求失败: %v", err))
	}
	defer resp.Body.Close()

	// 5. 检查服务器是否支持断点续传
	switch resp.StatusCode {
	case http.StatusPartialContent:
		fmt.Println("服务器支持断点续传，继续下载...")
	case http.StatusOK:
		fmt.Println("服务器不支持断点续传，重新下载...")
		fileSize = 0 // 重置已下载大小
	default:
		panic(fmt.Sprintf("服务器返回错误状态码: %d", resp.StatusCode))
	}

	// 6. 准备进度条
	contentLength := fileSize + resp.ContentLength
	tmpl := `{{ red "With funcs:" }} {{ bar . "<" "-" (cycle . "↖" "↗" "↘" "↙" ) "." ">"}} {{speed . | rndcolor }} {{percent .}} {{string . "my_green_string" | green}} {{string . "my_blue_string" | blue}}`
	// 基于 pb 的模板开启一个进度条
	bar := pb.ProgressBarTemplate(tmpl).Start64(contentLength)
	// 为 string 元素设置值
	bar.Set("my_green_string", "green").
		Set("my_blue_string", "blue")

	//bar := pb.New64(contentLength)
	bar.Set(pb.Bytes, true)
	bar.SetCurrent(fileSize)
	bar.Start()

	// 7. 创建带进度条的Reader
	proxyReader := bar.NewProxyReader(resp.Body)

	// 8. 执行下载
	_, err = io.Copy(file, proxyReader)
	bar.Finish()

	if err != nil {
		panic(fmt.Sprintf("下载失败: %v", err))
	}

	fmt.Printf("\n下载完成! 文件保存至: %s\n", filePath)
}

// 检查本地文件并返回可写文件句柄
func checkLocalFile(path string) (*os.File, int64, error) {
	// 获取文件信息
	fileInfo, err := os.Stat(path)
	if os.IsNotExist(err) {
		// 文件不存在，创建新文件
		file, err := os.Create(path)
		return file, 0, err
	}

	// 文件存在，以追加模式打开
	file, err := os.OpenFile(path, os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		return nil, 0, err
	}

	// 获取当前文件大小
	currentSize := fileInfo.Size()
	fmt.Printf("发现已下载文件: %s (%.2f MB)\n", path, float64(currentSize)/(1024*1024))

	return file, currentSize, nil
}
