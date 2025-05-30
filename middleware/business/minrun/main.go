package main

import (
	"business/core"
	"business/enum"
	"business/utils"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	dockerInstanceInfo, err := utils.GetDockerInstanceInfo()
	if err != nil {
		fmt.Println("get docker instance info error:", err)
		return
	}

	dockerInstanceKeys, err := utils.GetDockerInstanceKeys()
	if err != nil {
		fmt.Println("get docker instance keys error:", err)
		return
	}

	minionInfo, err := utils.GetMinionInfo()
	if err != nil {
		fmt.Println("get docker instance info error:", err)
		return
	}

	dockerInstance := map[string]float64{}

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
		dockerInstance[appidStr] += tx
	}

	// 统计每个实例的总量
	for i, v := range dockerInstanceInfo {
		keyStr, _ := v[1].(string)
		sumBw := dockerInstance[keyStr]
		dockerInstanceInfo[i] = append(dockerInstanceInfo[i], sumBw)
	}
	for _, v := range dockerInstanceInfo {
		fmt.Println("dockerInstanceInfo", v)
	}
}

func main1() {
	appID, err := utils.GetAppID()
	if err != nil {
		fmt.Println("Error getting app id:", err)
		return
	}

	if enum.BusinessTypeMixRun != appID {
		return
	}

	data, err := os.ReadFile(enum.PathBwFile)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	strData := string(data)
	strDataSplit := strings.Split(strData, " ")
	if len(strDataSplit) < 4 {
		fmt.Println("The data in the BW file does not meet expectations, and the data should be separated by at least three spaces")
		return
	}

	hostname, err := os.Hostname()
	if err != nil {
		fmt.Println("Error getting hostname:", err)
		return
	}

	if enum.Env == "dev" {
		hostname = "m-jiangsu-yangzhou-user3-1677888060-051"
	}

	networkCard := strDataSplit[0]
	bwSumTmp := strings.ReplaceAll(strDataSplit[3], "\n", "")
	bwSum, err := strconv.ParseFloat(bwSumTmp, 64)
	if bwSum <= 0 || err != nil {
		fmt.Printf("bandwidth sum is %v", bwSumTmp)
		return
	}

	dockerInstanceInfo, err := utils.GetDockerInstanceInfo()
	if err != nil {
		fmt.Println("get docker instance info error:", err)
		return
	}

	baseConfig, err := utils.GetBaseConfig()
	if err != nil {
		fmt.Printf("get base config error: %v\n", err)
		return
	}

	// 上报带宽
	err = core.ReportMinRunBandwidth(baseConfig.AgentSche.URL, hostname, appID, bwSum, dockerInstanceInfo)
	if err != nil {
		fmt.Println("report min run bandwidth error:", err)
		return
	}

	fmt.Println("Report Succeed")

	// 调整带宽
	err = core.ModifyMinRunBandwidth(baseConfig.AgentSche.URL, hostname, bwSum, networkCard, appID, dockerInstanceInfo)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Modify Succeed")
}
