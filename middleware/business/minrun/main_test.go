package main_test

import (
	"business/model"
	"business/utils"
	"encoding/json"
	"fmt"
	"math"
	"testing"
	"time"
)

func TestMain01(t *testing.T) {
	//path := "/home/yan/Documents/file/gofile/gotest/middleware/dowyin/bw_tmp"
	//err := utils.examineBwTmpFile(path)
	//fmt.Println(err)
}

func TestMain02(t *testing.T) {
	v, _ := utils.GetBwTmp("ljljsadlfjlajfl")
	fmt.Println(v)
	v.Bandwidth = 1000000
	err := utils.SetBwTmp(v)
	fmt.Println(err)
}

func TestMain03(t *testing.T) {
	bandwidth := 1.1
	bwTmp := model.BwTmp{
		NetworkCard: "",
		Bandwidth:   0,
		AppID:       "",
		UpdateAt:    "",
	}
	if bwTmp.Bandwidth != 0 && math.Abs(bwTmp.Bandwidth-bandwidth) > 0.2 {
		fmt.Println("0")
		return
	}
	fmt.Println("1")
}

func TestMain04(t *testing.T) {
	netInfo := utils.NewNetWorkInfo()

	// 获取所有物理网卡
	ifaces, err := netInfo.GetInterfaces()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	fmt.Printf("Physical Interfaces: %v\n\n", ifaces)

	// 显示网卡详细信息
	for _, iface := range ifaces {
		info, _ := netInfo.GetInterfaceInfo(iface)
		fmt.Printf("Interface: %s\n", iface)
		fmt.Printf("  MAC:      %s\n", info["mac"])
		fmt.Printf("  State:    %s\n", info["operstate"])
		fmt.Printf("  Type:     %s\n", info["type"])
		fmt.Printf("  Speed:    %s\n\n", info["speed"])
	}

	// 启动监控示例（监控eth0）
	fmt.Println("Starting network monitoring (Ctrl+C to stop)...")
	go utils.MonitorNetwork("eth0", 1*time.Second)

	// 保持主程序运行
	time.Sleep(30 * time.Second)
}

func TestMain05(t *testing.T) {
	pathUrl := "http://111.63.205.238:80" + "/agent/report/host_info"

	params := model.HostInfoReport{
		HostName:  "m-jiangsu-yangzhou-user3-1677888060-051",
		BandWidth: 0,
		Appid:     "be37b71de68ba3339cc196b6ef802706_2698d6c20affd188754ca34f17f43918",
	}

	paramsStr := utils.ObjToString(params)

	httpInstance := utils.NewHttp()
	respByte := httpInstance.Post(pathUrl, paramsStr)
	resp := model.ResMessage{}
	_ = json.Unmarshal(respByte, &resp)
	if resp.Code != 0 {
		panic(fmt.Errorf("ReportMinRunBandwidth err %v", resp.Message))
	}
	fmt.Println(resp)
}

func TestMain06(t *testing.T) {
	for i := 0; i < 100; i++ {
		v, ok := utils.GetBwTmp("be37b71de68ba3339cc196b6ef802706_2698d6c20affd188754ca34f17f43918")
		fmt.Println(v, ok)
		time.Sleep(1 * time.Second)
	}
}

func TestMain07(t *testing.T) {
	m := map[int]model.BizConf{}
	for i := 0; i < 100; i++ {
		m[i] = model.BizConf{InstanceCount: 2}
	}
	fmt.Println(m)
}

func TestMain08(t *testing.T) {
	data := []float64{2.434, 2.525, 2.999, 2.010}
	for _, v := range data {
		vTmp := utils.Round(v, 2)
		fmt.Println(v, vTmp)
	}
}

func TestMain09(t *testing.T) {
	v, _ := utils.GetMinionInfo()
	val := 0.0
	for _, f := range v {
		tmp := f * 8 / 1000 / 1000 / 60
		fmt.Println(tmp, f)
		val += tmp
	}
	fmt.Println(val)
}

// GetDockerData 获取docker数据
func TestMain10(t *testing.T) {
	v, err := utils.GetDockerInstanceKeys()

	fmt.Println(v, err, len(v))
}

// GetDockerData 获取docker数据
func TestMain11(t *testing.T) {
	fmt.Println(75879651 / 1000 / 1000)
}
