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
		hostname = "m-jiangsu-yangzhou-user3-1677888060-051"
	}

	networkCard := strDataSplit[0]
	bwTmp := strDataSplit[1]
	actionTmp := strDataSplit[2]
	bw, _ := strconv.ParseFloat(bwTmp, 64)
	action, _ := strconv.ParseInt(actionTmp, 10, 64)
	if enum.BusinessTypeMixRun == appID {
		if len(strDataSplit) <= 3 {
			fmt.Println("not find total bandwidth")
			return
		}
		bwSumTmp := strings.ReplaceAll(strDataSplit[3], "\n", "")
		bwSum, err := strconv.ParseFloat(bwSumTmp, 64)
		if bwSum <= 0 || err != nil {
			fmt.Printf("bandwidth sum is %v", bwSumTmp)
			return
		}

		fmt.Printf("Ready to get started：hostname：%s ，App ID:：%s，Network Card: %s, Bandwidth: %f, Action: %d, Bandwidth Sum:%f\n", hostname, appID, networkCard, bw, action, bwSum)
		dockerInstanceInfo, err := utils.GetDockerInstanceInfo()
		if err != nil {
			fmt.Println("get docker instance info error:", err)
			return
		}

		fmt.Println("docker instance info", dockerInstanceInfo)
		err = core.ReportMinRunBandwidth(hostname, appID, bwSum, dockerInstanceInfo)
		if err != nil {
			fmt.Println("report min run bandwidth error:", err)
			return
		}
	}
}
