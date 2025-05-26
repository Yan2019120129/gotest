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
	data, err := os.ReadFile("../bw")
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

	if len(strDataSplit) <= 3 {
		return
	}

	hostname, err := os.Hostname()
	if err != nil {
		fmt.Println("Error getting hostname:", err)
		return
	}

	switch appID {
	case enum.Business_Type_Doun_YIN:
		networkCard := strDataSplit[0]
		bw_tmp := strDataSplit[1]
		action := strDataSplit[2]
		appid := strDataSplit[3]
		fmt.Printf("Network Card: %s, Bandwidth: %s, Action: %s, App ID: %s\n", networkCard, bw_tmp, action, appid)
		bw, _ := strconv.ParseFloat(bw_tmp, 64)
		if bw <= 0 {
			fmt.Println("Invalid bandwidth value:", bw_tmp)
			return
		}
		core.ModifyDouYinBandwidth(hostname, bw)
	case enum.Business_Type_MIX_RUN:
		// modifyDouYinBandwidth(*result)
	}
}
