package main

import (
	"business/core"
	"business/enum"
	"business/utils"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	appID, err := utils.GetAppID()
	if err != nil {
		fmt.Println("Error getting app id:", err)
		return
	}

	if appID != enum.BusinessTypeMixRun || appID != enum.BusinessTypeDounYIN {
		fmt.Println("Unsupported business type:", appID)
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

	switch appID {
	case enum.BusinessTypeMixRun:
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
	case enum.BusinessTypeDounYIN:
		networkCard := strDataSplit[0]
		bwTmp := strDataSplit[1]
		actionTmp := strDataSplit[2]
		bw, _ := strconv.ParseFloat(bwTmp, 64)
		action, _ := strconv.ParseInt(actionTmp, 10, 64)
		// 针对抖音特殊处理，除以85%恢复到原值,并保留两位小数
		bw = math.Round(bw/0.85*100) / 100
		fmt.Printf("Ready to get started：hostname：%s ，App ID:：%s，Network Card: %s, Bandwidth: %f, Action: %s\n", hostname, appID, networkCard, bw, actionTmp)
		v, err := core.ModifyDouYinBandwidth(hostname, bw, action, networkCard, appID)
		if err != nil {
			fmt.Println("Error:", string(v), err)
			return
		}
		fmt.Println("Return Succeed：", string(v))
	default:
		fmt.Println("Unsupported business type:", appID)
		return
	}

}
