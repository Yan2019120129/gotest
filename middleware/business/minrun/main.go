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
	agent := utils.NewAgent()
	v, err := agent.Reboot()
	if err != nil {
		fmt.Printf("reboot error: %v-%v\n", err, v)
	}
	fmt.Printf("reboot succeed: %v-%v\n", err, v)
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
