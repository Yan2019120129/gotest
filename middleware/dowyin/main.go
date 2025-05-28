package main

import (
	"fmt"
	"gotest/middleware/dowyin/core"
	"gotest/middleware/dowyin/enum"
	"gotest/middleware/dowyin/utils"
	"os"
	"strconv"
	"strings"
)

func main() {
	data, err := os.ReadFile(enum.PathBwFile)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	strData := string(data)
	strDataSplit := strings.Split(strData, " ")

	appID, err := utils.GetAppID()
	if err != nil {
		fmt.Println("Error getting App ID:", err)
		return
	}

	hostname, err := os.Hostname()
	if err != nil {
		fmt.Println("Error getting hostname:", err)
		return
	}

	networkCard := strDataSplit[0]
	bw_tmp := strDataSplit[1]
	action_tmp := strDataSplit[2]
	bw, _ := strconv.ParseFloat(bw_tmp, 64)
	action, _ := strconv.ParseInt(action_tmp, 10, 64)

	switch appID {
	case enum.BusinessTypeDounYIN:
		fmt.Printf("Ready to get started：hostname：%s ，App ID:：%s，Network Card: %s, Bandwidth: %s, Action: %s", hostname, appID, networkCard, bw_tmp, action_tmp)
		v, err := core.ModifyDouYinBandwidth(hostname, bw, action, networkCard, appID)
		if err != nil {
			fmt.Println("Error:", string(v), err)
			return
		}
		fmt.Println("Return：", string(v))
	case enum.BusinessTypeMixRun:
		if len(strDataSplit) < 3 {
			fmt.Println("not find bw value")
			return
		}
		bwSum_tmp := strDataSplit[3]
		bwSum, _ := strconv.ParseFloat(bwSum_tmp, 64)
		dockerInstanceInfo, err := utils.GetDockerInstanceInfo()
		if err != nil {
			fmt.Println("get docker instance info error:", err)
			return
		}

		err = core.ReportMinRunBandwidth(hostname, appID, bwSum, dockerInstanceInfo)
		if err != nil {
			fmt.Println("report min run bandwidth error:", err)
			return
		}
	}
}
