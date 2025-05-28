package utils

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"gotest/middleware/business/enum"
	"gotest/middleware/business/model"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
	"time"
)

// GetAppID 获取AppId
func GetAppID() (string, error) {
	configFile := enum.PathRecruitResultFile
	type ResultItem struct {
		AppID string `json:"appid"`
	}
	type Data struct {
		Result []ResultItem `json:"result"`
	}

	if _, err := os.Stat(configFile); err == nil {
		data, err := os.ReadFile(configFile)
		if err != nil {
			return "", err
		}
		var result Data
		if err := json.Unmarshal(data, &result); err == nil && len(result.Result) > 0 {
			return result.Result[0].AppID, nil
		}
	}
	return "", fmt.Errorf("config file %s not found", configFile)
}

// examineBwTmpFile 检查BwTmp文件是否存在，不存在则创建
func examineBwTmpFile(path string) error {
	_, err := os.Stat(path)
	if err != nil {
		bytes, _ := json.Marshal(map[string]model.BwTmp{})
		err = os.WriteFile(path, bytes, os.ModePerm)
		return err
	}
	return nil
}

// GetBwTmp 获取带宽临时文件数据
func GetBwTmp(appid string) (model.BwTmp, bool) {
	v := model.BwTmp{}
	err := examineBwTmpFile(enum.PathBwTmpFile)
	if err != nil {
		fmt.Println("Err:", err)
		return v, false
	}

	file, err := os.ReadFile(enum.PathBwTmpFile)
	if err != nil {
		fmt.Println("Err:", err)
		return v, false
	}

	bwTmp := make(map[string]model.BwTmp)
	_ = json.Unmarshal(file, &bwTmp)
	v, ok := bwTmp[appid]

	return v, ok
}

// GetBwTmpAll 获取全部带宽临时文件数据
func GetBwTmpAll() (map[string]model.BwTmp, error) {
	err := examineBwTmpFile(enum.PathBwTmpFile)
	if err != nil {
		return nil, err
	}
	file, err := os.ReadFile(enum.PathBwTmpFile)
	if err != nil {
		return nil, err
	}

	bwTmp := make(map[string]model.BwTmp)
	_ = json.Unmarshal(file, &bwTmp)

	return bwTmp, nil
}

// SaveBwTmpAll 保存全部带宽临时文件数据
func SaveBwTmpAll(val map[string]model.BwTmp) error {
	if len(val) == 0 {
		return fmt.Errorf("The value cannot be empty")
	}

	for _, v := range val {
		if v.Percentage == 0 {
			v.Percentage = 1
		}
	}

	byteList, err := json.Marshal(val)
	if err != nil {
		return err
	}

	err = os.WriteFile(enum.PathBwTmpFile, byteList, os.ModePerm)
	if err != nil {
		return err
	}

	return nil
}

// SetBwTmp 设置带宽临时文件数据
func SetBwTmp(bwTmp model.BwTmp) error {
	if bwTmp.AppID == "" {
		return fmt.Errorf("the AppID can't be empty")
	}

	if bwTmp.Percentage == 0 {
		bwTmp.Percentage = 1
	}

	bwTmpList, err := GetBwTmpAll()
	if err != nil {
		return err
	}

	if bwTmp.UpdateAt == "" {
		bwTmp.UpdateAt = time.Now().Format(time.DateTime)
	}

	bwTmpList[bwTmp.AppID] = bwTmp

	valBytes, err := json.Marshal(bwTmpList)
	if err != nil {
		return err
	}

	err = os.WriteFile(enum.PathBwTmpFile, valBytes, os.ModePerm)
	if err != nil {
		return err
	}

	return nil
}

// GetDockerInstanceInfo 获取docker 实例信息
func GetDockerInstanceInfo() ([][]any, error) {
	var dockerInstance string
	var err error
	if enum.Env == "dev" {
		dockerInstance = `      1 
     48 2698d6c20affd188754ca34f17f43918
     30 be37b71de68ba3339cc196b6ef802706
`
	} else {
		dockerInstance, err = execDockerCommand()
		if err != nil {
			return nil, err
		}
	}

	var dockerInstanceInfo [][]any
	dockerInstanceLines := strings.Split(dockerInstance, "\n")
	for _, dockerInstanceLine := range dockerInstanceLines {
		dockerInstanceLine = strings.TrimSpace(dockerInstanceLine)
		dockerInstanceColumn := strings.Split(dockerInstanceLine, " ")
		if len(dockerInstanceColumn) <= 1 || dockerInstanceColumn[1] == "" {
			continue
		}

		count, _ := strconv.ParseFloat(dockerInstanceColumn[0], 64)
		tmp := []any{count, dockerInstanceColumn[1]}
		dockerInstanceInfo = append(dockerInstanceInfo, tmp)
	}

	return dockerInstanceInfo, nil
}

// execDockerCommand 执行docker命令
func execDockerCommand() (string, error) {
	command := "docker ps | awk '{print $NF}'| awk -F '[._]' '{print $3}' | sort | uniq -c"
	shell := "/bin/sh"
	if runtime.GOOS == "windows" {
		shell = "cmd"
	}
	cmd := exec.Command(shell, "-c", command)

	// 捕获标准输出和标准错误
	var stdout, stderr bytes.Buffer

	// 标准输出
	cmd.Stdout = &stdout

	// 标准错误
	cmd.Stderr = &stderr

	var err error
	// 执行命令
	err = cmd.Run()
	if err != nil {
		return "", err
	}

	if stderr.String() != "" {
		return "", errors.New(stderr.String())
	}
	return stdout.String(), nil
}
