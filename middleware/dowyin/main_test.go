package main_test

import (
	"fmt"
	"gotest/middleware/dowyin/model"
	"gotest/middleware/dowyin/utils"
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
	v := utils.GetBwTmp("ljljsadlfjlajfl")
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

	netInfo.GetInterfaceStats()

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
