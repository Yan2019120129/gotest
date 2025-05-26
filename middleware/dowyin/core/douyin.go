package core

import (
	"bytes"
	"crypto/md5"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"gotest/middleware/dowyin/model"
	"io"
	"net/http"
	"net/url"
	"time"
)

// ModifyDouYinBandwidth 修改抖音带宽
func ModifyDouYinBandwidth(hostname string, bandwidth float64, action int64) ([]byte, error) {
	if hostname == "" {
		return nil, fmt.Errorf("The hostname cannot be empty")
	}

	if action == 0 {
		return nil, fmt.Errorf("The action cannot be zero")
	}

	return DYControlBG(hostname, bandwidth, 2, 2)
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
