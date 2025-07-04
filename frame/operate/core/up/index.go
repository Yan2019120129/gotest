package up

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"operate/utils"
	"os"
)

// Up 上传文件
func Up(path string) error {
	if utils.IsExistFile(path) {
		return fmt.Errorf(path, "is not exist")
	}

	// 打开文件
	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// 创建 multipart 表单
	var requestBody bytes.Buffer
	writer := multipart.NewWriter(&requestBody)

	// 创建文件字段
	part, err := writer.CreateFormFile(fileField, file.Name())
	if err != nil {
		panic(err)
	}

	// 写入文件内容
	_, err = io.Copy(part, file)
	if err != nil {
		panic(err)
	}

	// 关闭 writer 完成 multipart 构造
	err = writer.Close()
	if err != nil {
		panic(err)
	}

	return nil
}
