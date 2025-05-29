package core

import (
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
	isBwTmpIsNull := len(bwTmp) == 0
	if isBwTmpIsNull {
		for _, dockerInstanceInfo := range dockerInstanceInfoList {
			if len(dockerInstanceInfo) <= 1 {
				fmt.Println("dockerInstanceInfo is not out of specifications")
				continue
			}

			businessAppid, _ := dockerInstanceInfo[1].(string)
			splicingBusinessAppid := appid + "_" + businessAppid

			bwTmp[splicingBusinessAppid] = model.BwTmp{
				Percentage: 1,
				AppID:      splicingBusinessAppid,
				UpdateAt:   time.Now().Format(time.DateTime),
			}
		}

		// 保存初始数据
		err = utils.SaveBwTmpAll(bwTmp)
		if err != nil {
			return fmt.Errorf("save bw tmp all error: %v", err)
		}
	}

	//  percentageSum 占比总和
	var percentageSum uint8 = 0
	for _, v := range bwTmp {
		if v.Percentage == 0 {
			v.Percentage = 1
		}
		percentageSum += v.Percentage
	}

	for _, dockerInstanceInfo := range dockerInstanceInfoList {
		if len(dockerInstanceInfo) <= 1 {
			fmt.Println("docker instance info is not out of specifications")
			continue
		}

		count, _ := dockerInstanceInfo[0].(float64)
		businessAppid, _ := dockerInstanceInfo[1].(string)
		splicingBusinessAppid := appid + "_" + businessAppid

		bwTmpVal, ok := bwTmp[splicingBusinessAppid]
		if !ok {
			return fmt.Errorf("bw tmp val not exist")
		}

		// 占比/总量*总带宽=业务总带宽
		businessBwSum := float64(bwTmpVal.Percentage) / float64(percentageSum) * bwSum

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
		fmt.Printf("report host info [bwSum:%f， Count:%f，HostName:%s，BusinessBwSum::%f，Appid:%s\n，Resp：%v]", bwSum, count, hostname, businessBwSum, splicingBusinessAppid, resp)
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
	isBwTmpIsNull := len(bwTmp) == 0
	if isBwTmpIsNull {
		for _, dockerInstanceInfo := range dockerInstanceInfoList {
			if len(dockerInstanceInfo) <= 1 {
				fmt.Println("docker instance info is not out of specifications")
				continue
			}

			businessAppid, _ := dockerInstanceInfo[1].(string)
			splicingBusinessAppid := appid + "_" + businessAppid

			bwTmp[splicingBusinessAppid] = model.BwTmp{
				Percentage: 1,
				AppID:      splicingBusinessAppid,
				UpdateAt:   time.Now().Format(time.DateTime),
			}
		}

		// 保存初始数据
		err = utils.SaveBwTmpAll(bwTmp)
		if err != nil {
			return fmt.Errorf("save bw tmp all error: %v", err)
		}
	}

	//  percentageSum 占比总和
	var percentageSum uint8 = 0
	for _, v := range bwTmp {
		if v.Percentage == 0 {
			v.Percentage = 1
		}
		percentageSum += v.Percentage
	}

	bizConf, err := utils.GetBizConf()
	if err != nil {
		return fmt.Errorf("get bizConf error: %v", err)
	}

	pathUrl := baseUrl + "/agent/get/host_sched_bandwidth"

	isRebootAgent := false

	for _, dockerInstanceInfo := range dockerInstanceInfoList {
		if len(dockerInstanceInfo) <= 1 {
			fmt.Println("docker instance info is not out of specifications")
			continue
		}

		count, _ := dockerInstanceInfo[0].(float64)
		businessAppid, _ := dockerInstanceInfo[1].(string)
		splicingBusinessAppid := appid + "_" + businessAppid
		//dockerInstanceInfo[1] = splicingBusinessAppid

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
			return fmt.Errorf("bw tmp val not exist")
		}

		// 占比/总量*总带宽=业务总带宽
		businessBwSum := float64(bwTmpVal.Percentage) / float64(percentageSum) * bwSum

		switch bandwidthInfo.Ret {
		case 0:
			continue
		case 1: // 根据带宽控制实例数

			// 业务总带宽/容器数量=每个容器的带宽
			averageBandwidth := businessBwSum / count

			// 目标带宽/单个实例带宽量=实例数
			count = math.Round(bandwidthInfo.Bandwidth / averageBandwidth)

			// 更新临时文件
			bwTmpVal.Bandwidth = bandwidthInfo.Bandwidth
			bwTmpVal.NetworkCard = networkCard
			bwTmpVal.UpdateAt = time.Now().Format(time.DateTime)
			fmt.Printf("Modify host info [BwSum:%f， Count:%f，HostName:%s，BusinessBwSum:%f，Appid:%s\n，Resp：%v]", bwSum, count, hostname, businessBwSum, splicingBusinessAppid, bandwidthInfo)
			isRebootAgent = true
		case -1: // 强制设置为0个实例
			count = 0
			bwTmpVal.Bandwidth = 0
			bwTmpVal.NetworkCard = networkCard
			bwTmpVal.UpdateAt = time.Now().Format(time.DateTime)

			isRebootAgent = true
			fmt.Printf("Modify host info [BwSum:%f， Count:%f，HostName:%s，BusinessBwSum:%f，Appid:%s\n，Resp：%v]", bwSum, count, hostname, businessBwSum, splicingBusinessAppid, bandwidthInfo)
		}

		bwTmp[splicingBusinessAppid] = bwTmpVal
		bizConf[businessAppid] = model.BizConf{
			InstanceCount: uint16(count),
		}
	}

	if isRebootAgent {
		err = utils.SaveBwTmpAll(bwTmp)
		if err != nil {
			return fmt.Errorf("save bw tmp all error: %v", err)
		}

		err = utils.SaveBizConf(bizConf)
		if err != nil {
			return fmt.Errorf("save zx_biz error: %v", err)
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
