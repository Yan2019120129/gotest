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
	v := 520709961.0
	fmt.Printf("%f", v/1000/1000/1000)
	return

	data, err := os.ReadFile(enum.PathBwFile)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	strData := string(data)
	strDataSplit := strings.Split(strData, " ")
	if len(strDataSplit) <= 2 {
		fmt.Println("The data in the BW file does not meet expectations, and the data should be separated by at least three spaces")
		return
	}

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

	if enum.Env == "dev" {
		hostname = "m-jiangsu-yangzhou-user3-1677888060-034"
	}

	networkCard := strDataSplit[0]
	bwTmp := strDataSplit[1]
	actionTmp := strDataSplit[2]
	bw, _ := strconv.ParseFloat(bwTmp, 64)
	action, _ := strconv.ParseInt(actionTmp, 10, 64)
	if enum.BusinessTypeDounYIN == appID {
		// 针对抖音特殊处理，除以85%恢复到原值,并保留两位小数
		bw = math.Round(bw/0.85*100) / 100
		fmt.Printf("Ready to get started：hostname：%s ，App ID:：%s，Network Card: %s, Bandwidth: %f, Action: %s\n", hostname, appID, networkCard, bw, actionTmp)
		v, err := core.ModifyDouYinBandwidth(hostname, bw, action, networkCard, appID)
		if err != nil {
			fmt.Println("Error:", string(v), err)
			return
		}
		fmt.Println("Return Succeed：", string(v))
	}
}
