package scheduled

import (
	"bandwidth_summary/enum"
	"bandwidth_summary/model"
	"bandwidth_summary/utils"
	"encoding/json"
	"fmt"
	"log"
	"sync"
)

var m sync.Map = sync.Map{}

// GetClentBw 获取客户端带宽数据
func GetClentBw(serverPath string) (model.NetInfo, error) {
	var netInfo model.NetInfo
	httpClient := utils.NewHttp()
	httpClient.AddParam("token", "e508a5785365cf0a4703469c5d2d0bcb")
	val, err := httpClient.Get(serverPath + enum.RouterNetInfo)
	if err != nil {
		log.Printf("get net info error:%s", err.Error())
		return netInfo, fmt.Errorf("get net info error:%s", err.Error()) // 如果没有数据，返回空的 NetInfo
	}
	_ = json.Unmarshal(val, &netInfo)

	return netInfo, nil
}

// GetClentBwSummary 汇总各服务器带宽
func GetClentBwSummary(kgGroup string, addrs []string) model.ReportTcInfo {
	reportTcInfo := model.ReportTcInfo{
		KgGroup: kgGroup,
	}

	// 调用获取客户端带宽数据的函数
	for _, addr := range addrs {
		netInfo, err := GetClentBw(addr)
		if err != nil {
			log.Printf("error getting bandwidth info from %s: %v\n", addr, err)
			continue
		}
		key := netInfo.Hostname + addr
		var bandwidth float64
		netInfoOld := model.NetInfo{}
		netInfoOld_tmp, ok := m.Load(key)
		if ok {
			netInfoOld = netInfoOld_tmp.(model.NetInfo)
		}

		bandwidth = float64(netInfo.RXBytes-netInfoOld.RXBytes) / float64(netInfo.Timestamp-netInfoOld.Timestamp) / 1000 / 1000 * 8
		if bandwidth == 0 {
			log.Printf("hostname:%s this bandwidth is zeor\n", netInfo.Hostname)
			continue
		}
		reportTcInfo.Bandwidth += bandwidth
		m.Store(key, netInfo)
	}

	return reportTcInfo
}

// ReportTcInfo 上报机房带宽
func ReportTcInfo(url string, reportInfo model.ReportTcInfo) error {
	log.Println("start report--- ", "url:", url, "report info:", reportInfo)
	httpInstance := utils.NewHttp()
	resp, err := httpInstance.Post(url+enum.RouterReportDCInfo, utils.ObjToString(reportInfo))
	if err != nil {
		return err
	}
	respMesage := model.ResMessage{}
	_ = json.Unmarshal(resp, &respMesage)

	if respMesage.Message != "success" || respMesage.Code != 0 {
		log.Printf("report tc info error: message=%s,code=%d", respMesage.Message, respMesage.Code)
		return fmt.Errorf("report tc info error: message=%s,code=%d", respMesage.Message, respMesage.Code)
	}
	log.Println("report success--- ", respMesage)

	return nil
}
