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
	data, err := os.ReadFile(enum.Bw_FILE_Path)
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

	fmt.Println("开始请求：", hostname, appID)
	switch appID {
	case enum.Business_Type_Doun_YIN:
		networkCard := strDataSplit[0]
		bw_tmp := strDataSplit[1]
		action_tmp := strDataSplit[2]
		//appid := strDataSplit[3]
		fmt.Printf("Network Card: %s, Bandwidth: %s, Action: %s, App ID: %s\n", networkCard, bw_tmp, action_tmp, appID)
		bw, _ := strconv.ParseFloat(bw_tmp, 64)
		if bw <= 0 {
			fmt.Println("Invalid bandwidth value:", bw_tmp)
			return
		}

		action, err := strconv.ParseInt(action_tmp, 10, 64)
		if err != nil {
			fmt.Println("Invalid bandwidth value:", bw_tmp)
		}
		v, err := core.ModifyDouYinBandwidth(hostname, bw, action)
		if err != nil {
			fmt.Println("请求失败:", string(v), err)
		}
		fmt.Println("请求成功:", string(v), err)
	case enum.Business_Type_MIX_RUN:
		// modifyDouYinBandwidth(*result)
	}
}
