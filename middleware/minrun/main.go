package main

import (
	"fmt"
	"gotest/middleware/minrun/core"
	"gotest/middleware/minrun/enum"
	"gotest/middleware/minrun/utils"
	"os"
	"strconv"
	"strings"
)

func main() {
	data, err := os.ReadFile(enum.Bw_FILE_Path)
	if err != nil {
		fmt.Println("error reading file:", err)
		return
	}
	strData := string(data)
	strDataSplit := strings.Split(strData, " ")

	appID, err := utils.GetAppID()
	if err != nil {
		fmt.Println("error getting App ID:", err)
		return
	}

	hostname, err := os.Hostname()
	if err != nil {
		fmt.Println("error getting hostname:", err)
		return
	}

	if enum.Business_Type_MIX_RUN == appID {
		networkCard := strDataSplit[0]
		bw_tmp := strDataSplit[1]
		action_tmp := strDataSplit[2]
		//appid := strDataSplit[3]
		fmt.Printf("Ready to get started：hostname：%s ，App ID:：%s，Network Card: %s, Bandwidth: %s, Action: %s", hostname, appID, networkCard, bw_tmp, action_tmp)
		bw, _ := strconv.ParseFloat(bw_tmp, 64)
		action, _ := strconv.ParseInt(action_tmp, 10, 64)

		//v, err := core.ModifyMinRunBandwidth(hostname, bw, action)
		//if err != nil {
		//	fmt.Println("The request was successful:", string(v), err)
		//	return
		//}
		//fmt.Println("The request was fail:", string(v), err)
	}

}
