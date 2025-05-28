package utils

import (
	"bufio"
	"business/enum"
	"encoding/json"
	"errors"
	"fmt"
	"io/fs"
	"net"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

// NetWorkInfo 结构体封装网卡信息获取逻辑
type NetWorkInfo struct {
	sysNetPath      string
	procNetDev      string
	virtualPrefixes []string
}

// NewNetWorkInfo 构造函数初始化路径和虚拟网卡前缀
func NewNetWorkInfo() *NetWorkInfo {
	return &NetWorkInfo{
		sysNetPath:      "/sys/class/net",
		procNetDev:      "/proc/net/dev",
		virtualPrefixes: []string{"docker", "br-", "veth", "lo"},
	}
}

// GetInterfaces 获取所有物理网卡名称
func (n *NetWorkInfo) GetInterfaces() ([]string, error) {
	interfaces, err := net.Interfaces()
	if err != nil {
		return nil, fmt.Errorf("failed to get interfaces: %v", err)
	}

	var validIfaces []string
	for _, iface := range interfaces {
		// 过滤虚拟网卡
		isVirtual := false
		for _, prefix := range n.virtualPrefixes {
			if strings.HasPrefix(iface.Name, prefix) {
				isVirtual = true
				break
			}
		}
		if !isVirtual {
			validIfaces = append(validIfaces, iface.Name)
		}
	}
	return validIfaces, nil
}

// GetInterfaceInfo 获取单个网卡详细信息
func (n *NetWorkInfo) GetInterfaceInfo(ifaceName string) (map[string]string, error) {
	info := make(map[string]string)
	interfacePath := filepath.Join(n.sysNetPath, ifaceName)

	// 获取MAC地址
	addressPath := filepath.Join(interfacePath, "address")
	if data, err := os.ReadFile(addressPath); err == nil {
		info["mac"] = strings.TrimSpace(string(data))
	}

	// 获取操作状态
	operstatePath := filepath.Join(interfacePath, "operstate")
	if data, err := os.ReadFile(operstatePath); err == nil {
		info["operstate"] = strings.TrimSpace(string(data))
	}

	// 获取网卡类型
	info["type"] = n.getInterfaceType(ifaceName)

	// 获取物理速率
	speedPath := filepath.Join(interfacePath, "speed")
	if data, err := os.ReadFile(speedPath); err == nil {
		speed := strings.TrimSpace(string(data))
		if s, err := strconv.Atoi(speed); err == nil {
			info["speed"] = fmt.Sprintf("%dMbps", s)
		} else {
			info["speed"] = "unknown"
		}
	} else {
		info["speed"] = "unknown"
	}

	return info, nil
}

// getInterfaceType 判断网卡类型
func (n *NetWorkInfo) getInterfaceType(ifaceName string) string {
	if strings.HasPrefix(ifaceName, "br-") {
		return "bridge"
	} else if strings.HasPrefix(ifaceName, "docker") {
		return "docker"
	} else if strings.HasPrefix(ifaceName, "veth") {
		return "virtual"
	} else if ifaceName == "lo" {
		return "loopback"
	}
	return "physical"
}

// GetInterfaceStats 获取网卡统计信息
func (n *NetWorkInfo) GetInterfaceStats(ifaceName string) (uint64, uint64, error) {
	data, err := os.ReadFile(n.procNetDev)
	if err != nil {
		return 0, 0, fmt.Errorf("failed to read proc net dev: %v", err)
	}

	scanner := bufio.NewScanner(strings.NewReader(string(data)))
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, ifaceName+":") {
			fields := strings.Fields(line)
			rxBytes := parseHexToUint64(fields[1])
			txBytes := parseHexToUint64(fields[9])
			return rxBytes, txBytes, nil
		}
	}
	return 0, 0, fmt.Errorf("interface %s not found", ifaceName)
}

// parseHexToUint64 将十六进制字符串转换为uint64
func parseHexToUint64(hexStr string) uint64 {
	value, _ := strconv.ParseUint(strings.Replace(hexStr, ":", "", -1), 16, 64)
	return value
}

// GetInterfaceSpeed 获取网卡实时速率
func (n *NetWorkInfo) GetInterfaceSpeed(ifaceName string, interval time.Duration) (float64, float64, error) {
	rx1, tx1, err := n.GetInterfaceStats(ifaceName)
	if err != nil {
		return 0, 0, err
	}

	time.Sleep(interval)

	rx2, tx2, err := n.GetInterfaceStats(ifaceName)
	if err != nil {
		return 0, 0, err
	}

	rxSpeed := float64(rx2-rx1) / interval.Seconds() / 1024 / 1024 // MB/s
	txSpeed := float64(tx2-tx1) / interval.Seconds() / 1024 / 1024 // MB/s

	return rxSpeed, txSpeed, nil
}

// MonitorNetwork 流量监控协程
func MonitorNetwork(ifaceName string, interval time.Duration) {
	netInfo := NewNetWorkInfo()
	for {
		rxSpeed, txSpeed, err := netInfo.GetInterfaceSpeed(ifaceName, interval)
		if err != nil {
			fmt.Printf("Error getting speed: %v\n", err)
			return
		}
		fmt.Printf("\r%s - RX: %.2f MB/s | TX: %.2f MB/s", ifaceName, rxSpeed, txSpeed)
		os.Stdout.Sync()
		time.Sleep(interval)
	}
}

// businessMacConfig 定义业务网卡配置结构
type businessMacConfig struct {
	BusinessNIC string        `json:"business_nic"`
	PPPoENics   []pppoENicDef `json:"pppoe_nics,omitempty"`
}

// pppoENicDef 定义PPPoE网卡定义结构
type pppoENicDef struct {
	PPPoEDevice string `json:"pppoe_device"`
	Eth         string `json:"eth"`
}

// networkV2Config 定义网络配置结构
type networkV2Config struct {
	PPPoEConfig struct {
		PPPoENics []pppoENicDef `json:"pppoe_nics"`
	} `json:"pppoe_config"`
}

// GetRunDev 获取运行网卡
func GetRunDev() (string, error) {
	var config businessMacConfig

	// 读取业务配置文件
	data, err := os.ReadFile(enum.PathBusinessMacFile)
	if err != nil {
		if !errors.Is(err, fs.ErrNotExist) {
			return "", fmt.Errorf("读取业务配置失败: %w", err)
		}
		// 文件不存在时使用默认值
		return enum.DefaultNic, nil
	}

	// 解析JSON配置
	if err := json.Unmarshal(data, &config); err != nil {
		return "", fmt.Errorf("解析业务配置失败: %w", err)
	}

	// 获取基础网卡名称
	retDev := config.BusinessNIC
	if retDev == "" {
		retDev = enum.DefaultNic
	}

	// 处理特殊场景
	if retDev == "vr_veth_host" && len(config.PPPoENics) > 0 {
		pppoeDevice := config.PPPoENics[0].PPPoEDevice
		return GetRealDef(pppoeDevice)
	}

	return retDev, nil
}

// GetRealDef 获取实际网卡名称
func GetRealDef(pppoeDev string) (string, error) {
	var config networkV2Config

	// 读取网络配置文件
	data, err := os.ReadFile(enum.PathNetworkV2File)
	if err != nil {
		if !errors.Is(err, fs.ErrNotExist) {
			return "", fmt.Errorf("读取网络配置失败: %w", err)
		}
		return pppoeDev, nil
	}

	// 解析JSON配置
	if err := json.Unmarshal(data, &config); err != nil {
		return "", fmt.Errorf("解析网络配置失败: %w", err)
	}

	// 查找匹配的PPPoE设备
	for _, nic := range config.PPPoEConfig.PPPoENics {
		if nic.PPPoEDevice == pppoeDev {
			return nic.Eth, nil
		}
	}

	return pppoeDev, nil
}
