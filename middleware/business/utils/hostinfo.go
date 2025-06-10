package utils

import (
	"business/enum"
	"business/model"
	"encoding/json"
	"fmt"
	"os"
	"regexp"
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
		return fmt.Errorf("the value cannot be empty")
	}

	for _, v := range val {
		if v.UpdateAt == "" {
			v.UpdateAt = time.Now().Format(time.Now().Format(time.DateTime))
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

// GetBizConf 获取容器控制文件
func GetBizConf() (map[string]model.BizConf, error) {
	err := examineBwTmpFile(enum.PathBizConfFile)
	if err != nil {
		return nil, err
	}

	file, err := os.ReadFile(enum.PathBizConfFile)
	zxbiz := make(map[string]model.BizConf)
	_ = json.Unmarshal(file, &zxbiz)
	return zxbiz, err
}

// SaveBizConf 保存容器控制文件
func SaveBizConf(val map[string]model.BizConf) error {
	byteList, err := json.Marshal(val)
	if err != nil {
		return err
	}

	err = os.WriteFile(enum.PathBizConfFile, byteList, os.ModePerm)
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
		command := "docker ps | awk '{print $NF}'| awk -F '[._]' '{print $3}' | sort | uniq -c"
		dockerInstance, err = ExecCommand(command)
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

	dockerInstanceKeys, err := GetDockerInstanceKeys()
	if err != nil {
		return nil, err
	}

	minionInfo, err := GetMinionInfo()
	if err != nil {
		return nil, err
	}

	dockerInstanceMap := map[string]float64{}

	// 添加每个容器的上行数据
	for i, v := range dockerInstanceKeys {
		if len(v) == 1 {
			continue
		}
		keyStr, _ := v[0].(string)
		appidStr, _ := v[1].(string)
		tx, ok := minionInfo[keyStr]
		if ok {
			tx = tx * 8 / 1000 / 1000 / 60
			dockerInstanceKeys[i] = append(dockerInstanceKeys[i], tx)
		}
		dockerInstanceMap[appidStr] += tx
	}

	// 统计每个实例的总量
	for i, v := range dockerInstanceInfo {
		keyStr, _ := v[1].(string)
		sumBw := dockerInstanceMap[keyStr]
		dockerInstanceInfo[i] = append(dockerInstanceInfo[i], sumBw)
	}

	return dockerInstanceInfo, nil
}

// GetDockerInstanceKeys 获取容器信息
func GetDockerInstanceKeys() ([][]any, error) {
	var command string
	if enum.Env == "dev" {
		command = "cat /home/yan/Documents/file/gofile/gotest/middleware/business/tmp.txt | tail -n +2 | awk '{print $1, $NF}'"
	} else {
		command = "docker ps | tail -n +2 | awk '{print $1, $NF}'"
	}

	dockerInstance, err := ExecCommand(command)
	if err != nil {
		return nil, err
	}

	var dockerInstanceKeys [][]any
	dockerLines := strings.Split(dockerInstance, "\n")
	for _, line := range dockerLines {
		if len(line) <= 20 {
			continue
		}
		spaceIndex := strings.Index(line, " ")
		underIndex := strings.LastIndex(line, "_")
		dotIndex := strings.LastIndex(line, ".")
		key := strings.TrimSpace(line[:spaceIndex+1])
		appid := strings.TrimSpace(line[underIndex+1 : dotIndex])
		dockerInstanceKeys = append(dockerInstanceKeys, []any{key, appid})
	}
	return dockerInstanceKeys, nil
}

// GetMinionInfo 获取Minion信息
func GetMinionInfo() (map[string]float64, error) {
	nowTime := time.Now()
	nowTimeStr := nowTime.Format("2006-01-02 15:04")
	if enum.Env == "dev" {
		nowTimeStr = "2025-05-29 18:26"
	}
	var val string
	var err error

	command := fmt.Sprintf("tail -1000 %s | grep '%s.*containerId'", enum.PathMinionLogFile, nowTimeStr)
	val, err = ExecCommand(command)
	if err != nil {
		nowTimeStr = nowTime.Add(-time.Minute).Format("2006-01-02 15:04")
		command = fmt.Sprintf("tail -2000 %s | grep '%s.*containerId'", enum.PathMinionLogFile, nowTimeStr)
		val, err = ExecCommand(command)
		if err != nil {
			return nil, fmt.Errorf("get minion error:%v", err)
		}
	}

	logLines := strings.Split(val, "\n")

	// 编译正则表达式
	re := regexp.MustCompile(`containerId:([a-f0-9]{12}).*?Tx:(\d+)`)

	txMap := make(map[string]float64)

	// 处理每行日志
	for _, line := range logLines {
		matches := re.FindStringSubmatch(line)
		if len(matches) >= 3 { // 完整匹配+两个捕获组
			flowID := matches[1]
			txValue := matches[2]
			txValueF, _ := strconv.ParseFloat(txValue, 64)
			txMap[flowID] = txValueF
		}
	}

	return txMap, nil
}

// GetBZInstanceCount 获取bi站最大实例数
func GetBZInstanceCount() uint8 {
	instanceCountStr, err := ExecCommand(fmt.Sprintf("lsblk |grep %s |wc -l", enum.BusinessTypeMixRunBZ))
	if err != nil {
		fmt.Println("GetBZInstanceCount exec command error: ", err)
		return enum.DefaultContainerMaxBZ
	}

	instanceCountStr = strings.ReplaceAll(instanceCountStr, "\n", "")

	instanceCountInt64, err := strconv.ParseInt(instanceCountStr, 10, 64)
	if err != nil {
		fmt.Println("GetBZInstanceCount parse int error: ", err)
		return enum.DefaultContainerMaxBZ
	}

	return uint8(instanceCountInt64)
}
