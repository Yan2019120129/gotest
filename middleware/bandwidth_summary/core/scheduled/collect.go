package scheduled

import "bandwidth_summary/utils"

// GetClentBw 获取客户端带宽数据
func GetClentBw(serverPath string) any {
	httpClient := utils.NewHttp()
	httpClient.Get(serverPath + "/api/bandwidth_summary/client")
	return nil
}

func GetClentBwSummary(addrs []string) {

	// 调用获取客户端带宽数据的函数
	GetClentBw("")
}
