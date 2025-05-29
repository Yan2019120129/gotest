package core

import (
	"business/model"
	"business/utils"
	"bytes"
	"crypto/md5"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"math"
	"net/http"
	"net/url"
	"time"
)

// ModifyDouYinBandwidth 修改抖音带宽
func ModifyDouYinBandwidth(hostname string, bandwidth float64, action int64, networkCard, appID string) ([]byte, error) {
	if appID == "" {
		return nil, fmt.Errorf("the appid cannot be empty")
	}

	if hostname == "" {
		return nil, fmt.Errorf("the hostname cannot be empty")
	}

	if action == 0 {
		return nil, fmt.Errorf("the action cannot be zero")
	}

	bwTmp, ok := utils.GetBwTmp(appID)
	abs := math.Abs(bwTmp.Bandwidth - bandwidth)
	fmt.Printf("bandwidth abs: 【|%f-%f|=%v】 \n", bwTmp.Bandwidth, bandwidth, abs)
	if ok && bwTmp.Bandwidth != 0 && abs < 0.2 {
		return nil, fmt.Errorf("the bandwidth cannot be more than 200M")
	}

	res, err := DYControlBG(hostname, bandwidth, 2, 2)
	if err != nil {
		return nil, err
	}

	resMessage := model.ResMessage{}
	_ = json.Unmarshal(res, &resMessage)

	if resMessage.Code != 0 {
		return nil, fmt.Errorf("the request is incorrect%s", resMessage.Message)
	}

	newBwTmp := model.BwTmp{
		NetworkCard: networkCard,
		Bandwidth:   bandwidth,
		AppID:       appID,
		Count:       1,
		UpdateAt:    time.Now().Format(time.DateTime),
	}

	err = utils.SetBwTmp(newBwTmp)
	if err != nil {
		return res, fmt.Errorf("set bandwidth fail:%v", err)
	}

	fmt.Println("save the bw_tmp.json file")

	return res, nil
}

// DYControlBG 控制抖音带宽
func DYControlBG(hostname string, BandWidth, BandWay float64, from int) ([]byte, error) {
	hBand := &model.BandHost{Host: hostname, BandWidth: BandWidth, BandWay: BandWay, From: from}

	u := SignUrl("https://mc-service.vod.snv1.com/index.php", "Node.SetNodeProvideBand", "kmanage", "d683575969b52144f29da0efcf391454")

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	reqData, _ := json.Marshal(hBand)
	req, err := client.Post(u, "application/json", bytes.NewReader(reqData))
	if err != nil {
		return nil, err
	}
	defer req.Body.Close()
	data, err := io.ReadAll(req.Body)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// SignUrl 路由签名
func SignUrl(path, action, user, key string) string {
	ts := time.Now().Unix()
	md5Base := fmt.Sprintf("%s%s%s%d", action, user, key, ts)
	token := md5.Sum([]byte(md5Base))
	param := url.Values{}
	param.Set("Action", action)
	param.Set("Sign", fmt.Sprintf("%s-%d-%x", user, ts, token))
	return path + "?" + param.Encode()
}
