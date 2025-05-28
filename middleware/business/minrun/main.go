package main

import (
	"fmt"
	"gotest/middleware/business/core"
	"gotest/middleware/business/enum"
	"gotest/middleware/business/utils"
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
	bw_tmp := strDataSplit[1]
	action_tmp := strDataSplit[2]
	bw, _ := strconv.ParseFloat(bw_tmp, 64)
	action, _ := strconv.ParseInt(action_tmp, 10, 64)
	if enum.BusinessTypeMixRun == appID {
		if len(strDataSplit) <= 3 {
			fmt.Println("not find total bandwidth")
			return
		}
		bwSum_tmp := strDataSplit[3]
		bwSum, _ := strconv.ParseFloat(bwSum_tmp, 64)
		fmt.Printf("Ready to get started：hostname：%s ，App ID:：%s，Network Card: %s, Bandwidth: %f, Action: %d, Bandwidth Sum:%f\n", hostname, appID, networkCard, bw, action, bwSum)
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
