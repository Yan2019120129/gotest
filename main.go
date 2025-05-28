package main

import (
	"fmt"
	"math"
)

const (
	RootPath               = "/xyapp/system/miner.plugin-zxagent.ipk"
	Log_FiLE_Path          = RootPath + "/logs/app.log"
	Bw_FILE_Path           = RootPath + "/bw"
	Bw_tmp_FILE_Path       = RootPath + "/bw_tmp"
	Business_Type_Doun_YIN = "52d531d3ea193a292485d06517b4b5fd" // 抖音 appid 标识
	Business_Type_MIX_RUN  = "be37b71de68ba3339cc196b6ef802706" // 混跑 appid 标识
)

func main() {
	// 执行计算并保留两位小数
	bw := 4.2
	bw = math.Round(bw/0.85*100) / 100
	fmt.Println("bw:", bw)
}
