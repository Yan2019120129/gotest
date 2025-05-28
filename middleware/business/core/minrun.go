package core

import (
	"business/model"
	"business/utils"
	"encoding/json"
	"fmt"
	"time"
)

// ReportMinRunBandwidth 上报混跑业务带宽数据
func ReportMinRunBandwidth(hostname, appid string, bwSum float64, dockerInstanceInfoList [][]any) error {
	baseConfig, err := utils.GetBaseConfig()
	if err != nil {
		return fmt.Errorf("GetBaseConfig error: %v", err)
	}

	pathUrl := baseConfig.AgentSche.URL + "/agent/report/host_info"

	bwTmp, err := utils.GetBwTmpAll()
	if err != nil {
		return fmt.Errorf("GetBwTmpAll error: %v", err)
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
			businessAppid = appid + "_" + businessAppid

			bwTmp[businessAppid] = model.BwTmp{
				Percentage: 1,
				AppID:      businessAppid,
				UpdateAt:   time.Now().Format(time.DateTime),
			}
		}

		// 保存初始数据
		err = utils.SaveBwTmpAll(bwTmp)
		if err != nil {
			return fmt.Errorf("SaveBwTmpAll error: %v", err)
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
			fmt.Println("dockerInstanceInfo is not out of specifications")
			continue
		}

		count, _ := dockerInstanceInfo[0].(float64)
		businessAppid, _ := dockerInstanceInfo[1].(string)
		businessAppid = appid + "_" + businessAppid
		dockerInstanceInfo[1] = businessAppid

		bwTmpVal, ok := bwTmp[businessAppid]
		if !ok {
			return fmt.Errorf("bwTmpVal not exist")
		}

		businessBwSum := float64(bwTmpVal.Percentage) / float64(percentageSum) * bwSum
		//averageBandwidth := businessBwSum / count
		bandWidth := businessBwSum / count
		dockerInstanceInfo = append(dockerInstanceInfo, bandWidth)

		params := model.ReportHostInfo{
			HostName:  hostname,
			BandWidth: bandWidth,
			Appid:     businessAppid,
		}

		paramsStr := utils.ObjToString(params)
		httpInstance := utils.NewHttp()
		respByte := httpInstance.Post(pathUrl, paramsStr)
		resp := model.ResMessage{}
		_ = json.Unmarshal(respByte, &resp)
		fmt.Printf("report host info【bwSum:%f， Count:%f，HostName:%s，BandWidth::%f，Appid:%s\n，Resp：%v】", bwSum, count, hostname, bandWidth, businessAppid, resp)
		if resp.Code != 0 {
			return fmt.Errorf("report min run bandwidth err ,value:%v,resp:%v", params, resp.Message)
		}
	}

	return nil
}

// ModifyMinRunBandwidth 修改混跑带宽
func ModifyMinRunBandwidth(hostname string, bandwidth float64, action int64, networkCard, appID string) ([]byte, error) {
	if appID == "" {
		return nil, fmt.Errorf("the appid cannot be empty")
	}

	if hostname == "" {
		return nil, fmt.Errorf("the hostname cannot be empty")
	}

	if action == 0 {
		return nil, fmt.Errorf("the action cannot be zero")
	}

	return nil, nil
}
