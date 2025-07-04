package core

import (
	"business/enum"
	"business/model"
	"business/utils"
	"encoding/json"
	"fmt"
	"math"
	"time"
)

// ReportMinRunBandwidth 上报混跑业务带宽数据
func ReportMinRunBandwidth(baseUrl, hostname, appid string, bwSum float64, dockerInstanceInfoList [][]any) error {
	bwTmp, err := utils.GetBwTmpAll()
	if err != nil {
		return fmt.Errorf("get bw tmp all error: %v", err)
	}

	// 记录存储文件是否为空，如果为空设置初始比例
	err = detectIsInitBwTmp(appid, bwTmp, dockerInstanceInfoList)
	if err != nil {
		return err
	}

	for _, dockerInstanceInfo := range dockerInstanceInfoList {
		if len(dockerInstanceInfo) < 2 {
			fmt.Println("docker instance info is not out of specifications")
			continue
		}

		count, _ := dockerInstanceInfo[0].(float64)
		businessAppid, _ := dockerInstanceInfo[1].(string)
		businessBwSum, _ := dockerInstanceInfo[2].(float64)
		splicingBusinessAppid := appid + "_" + businessAppid

		params := model.HostInfoReport{
			HostName:  hostname,
			BandWidth: businessBwSum,
			Appid:     splicingBusinessAppid,
		}

		paramsStr := utils.ObjToString(params)
		httpInstance := utils.NewHttp()
		pathUrl := baseUrl + "/agent/report/host_info"
		respByte := httpInstance.Post(pathUrl, paramsStr)
		resp := model.ResMessage{}
		_ = json.Unmarshal(respByte, &resp)
		fmt.Printf("report host info [bwSum:%f， Count:%f，HostName:%s，BusinessBwSum::%f，Appid:%s，Resp：%v]\n",
			bwSum, count, hostname, businessBwSum, splicingBusinessAppid, resp)
		if resp.Code != 0 {
			return fmt.Errorf("report min run bandwidth err ,value:%v,resp:%v", params, resp.Message)
		}
	}

	return nil
}

// ModifyMinRunBandwidth 修改混跑带宽
func ModifyMinRunBandwidth(baseUrl, hostname string, bwSum float64, networkCard, appid string, dockerInstanceInfoList [][]any) error {
	if appid == "" {
		return fmt.Errorf("the appid cannot be empty")
	}

	if hostname == "" {
		return fmt.Errorf("the hostname cannot be empty")
	}

	bwTmp, err := utils.GetBwTmpAll()
	if err != nil {
		return fmt.Errorf("get bw tmp all error: %v", err)
	}

	// 记录存储文件是否为空，如果为空设置初始比例
	err = detectIsInitBwTmp(appid, bwTmp, dockerInstanceInfoList)
	if err != nil {
		return err
	}

	bizConf, err := utils.GetBizConf()
	if err != nil {
		return fmt.Errorf("get bizConf error: %v", err)
	}

	pathUrl := baseUrl + "/agent/get/host_sched_bandwidth"

	isRebootAgent := false
	for _, dockerInstanceInfo := range dockerInstanceInfoList {
		if len(dockerInstanceInfo) < 2 {
			fmt.Println("docker instance info is not out of specifications")
			continue
		}

		count, _ := dockerInstanceInfo[0].(float64)
		businessAppid, _ := dockerInstanceInfo[1].(string)
		businessBwSum, _ := dockerInstanceInfo[2].(float64)
		splicingBusinessAppid := appid + "_" + businessAppid

		param := model.HostInfoParams{
			HostName: hostname,
			Appid:    splicingBusinessAppid,
		}

		httpInstance := utils.NewHttp()
		respByte := httpInstance.Post(pathUrl, utils.ObjToString(param))
		var resp model.ResMessage
		_ = json.Unmarshal(respByte, &resp)
		if resp.Code != 0 {
			return fmt.Errorf("get host sched_bandwidth err ,resp:%v", resp.Message)
		}

		bandwidthInfo := model.BandwidthInfo{}
		_ = json.Unmarshal(resp.Data, &bandwidthInfo)

		bwTmpVal, ok := bwTmp[splicingBusinessAppid]
		if !ok {
			return fmt.Errorf("bw_tmp.json val not exist")
		}

		var limitCount uint8

		// 判断使用那个容器的最大值限制
		// 限制：根据业务比例，限制容器数量
		switch businessAppid {
		case enum.BusinessTypeMixRunBZ:
			limitCount = utils.GetBZInstanceCount()
		case enum.BusinessTypeMixRunG3:
			limitCount = 60
		}

		switch bandwidthInfo.Ret {
		case 0:
			fmt.Printf("the corresponding code is zero:%v\n", bandwidthInfo)
			continue
		case 1: // 根据带宽控制实例数

			// 业务总带宽(服务端计算值)/容器数量=每个容器的带宽
			//averageBandwidth := bandwidthInfo.BandwidthOrig / count

			// 业务总带宽(本地计算值)/容器数量=每个容器的带宽
			averageBandwidth := businessBwSum / count

			// 目标带宽/单个实例带宽量=实例数
			targetCount := math.Round(bandwidthInfo.Bandwidth / averageBandwidth)

			// 不能超过最大值
			if targetCount > float64(limitCount) {
				targetCount = float64(limitCount)
			}

			// 不能低于最小值
			if targetCount < 1 {
				targetCount = 1
			}

			if bwTmpVal.Count != uint16(targetCount) {
				isRebootAgent = true
			}

			// 更新临时文件
			bwTmpVal.Count = uint16(targetCount)
			bwTmpVal.Bandwidth = bandwidthInfo.Bandwidth
			bwTmpVal.NetworkCard = networkCard
			bwTmpVal.UpdateAt = time.Now().Format(time.DateTime)
			fmt.Printf("request information-%s ref:%d, count:%f，maxcount:%f,targetCount:%f,actualBw:%f,bw:%f,targetBw:%f\n",
				splicingBusinessAppid, bandwidthInfo.Ret, count, limitCount, targetCount, businessBwSum, bandwidthInfo.BandwidthOrig, bandwidthInfo.Bandwidth)
		case -1: // 强制设置为2个实例
			if bwTmpVal.Count != 0 {
				isRebootAgent = true
			}
			count = 2
			bwTmpVal.Count = 2
			bwTmpVal.Bandwidth = 0
			bwTmpVal.NetworkCard = networkCard
			bwTmpVal.UpdateAt = time.Now().Format(time.DateTime)
			isRebootAgent = true
			fmt.Printf("request information-%s ref:%d, count:%f，maxcount:%f,targetCount:%f,actualBw:%f,bw:%f,targetBw:%f\n",
				splicingBusinessAppid, bandwidthInfo.Ret, count, limitCount, 0.0, businessBwSum, bandwidthInfo.BandwidthOrig, bandwidthInfo.Bandwidth)
		}

		bwTmp[splicingBusinessAppid] = bwTmpVal
		bizConf[businessAppid] = model.BizConf{
			InstanceCount: bwTmpVal.Count,
		}
	}

	if isRebootAgent {
		err = utils.SaveBwTmpAll(bwTmp)
		if err != nil {
			return fmt.Errorf("save bw_tmp.json error: %v", err)
		}

		err = utils.SaveBizConf(bizConf)
		if err != nil {
			return fmt.Errorf("save biz_conf.josn error: %v", err)
		}

		agent := utils.NewAgent()
		v, err := agent.Reboot()
		if err != nil {
			return fmt.Errorf("reboot error: %v-%v", err, v)
		}
		fmt.Println("bw_tmp.json biz_conf.json file is successfully saved")
	}

	return nil
}

// detectIsInitBwTmp 检测是否初始化bw_tmp 初始化文件
func detectIsInitBwTmp(appid string, bwTmp map[string]model.BwTmp, dockerInstanceInfoList [][]any) error {
	// 记录存储文件是否为空，如果为空设置初始比例
	if len(bwTmp) == 0 {
		for _, dockerInstanceInfo := range dockerInstanceInfoList {
			if len(dockerInstanceInfo) <= 1 {
				fmt.Println("docker instance info is not out of specifications")
				continue
			}

			businessAppid, _ := dockerInstanceInfo[1].(string)
			splicingBusinessAppid := appid + "_" + businessAppid

			bwTmp[splicingBusinessAppid] = model.BwTmp{
				AppID:    splicingBusinessAppid,
				UpdateAt: time.Now().Format(time.DateTime),
			}
		}

		// 保存初始数据
		err := utils.SaveBwTmpAll(bwTmp)
		if err != nil {
			return fmt.Errorf("save bw tmp all error: %v", err)
		}
	}
	return nil
}
