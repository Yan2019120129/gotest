package main

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"gotest/common/utils"
	"io"
	"strings"
	"time"
)

var encryptionData = "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"
var key = "tradingeconomics-charts-core-api-key"

const (
	klinePath = "https://d3ii0wo49og5mi.cloudfront.net/markets"
)

func Kline() {
	instance := utils.NewHttp()
	instance.AddParam("interval", "5m") // 1分:1m 5分:5m 15分:15m 1小时:1h 1天:1d 1周:1w 1月:1month
	//instance.AddParam("span", "1d")     // 1天:1d 1周:1w 1月:1m 6月:6m 1年:1y 5年:5y 10年:10y 25年:25y 50年:50y
	instance.AddParam("n", "20") // 1m:300
	instance.AddParam("ohlc", "1")
	instance.AddParam("key", "20240229:nazare")
	v := instance.Get(klinePath + "/jpyusd:cur")
	val, err := conversion(v)
	ticker := val[len(val)-1]
	date := time.Unix(ticker.CreatedAt, 0)
	fmt.Printf("%+v\n%v\n%v", val[len(val)-1], err, date)
	//fmt.Printf("%+v", len(val.Series[0].Data))
}

// conversion 转换数据
func conversion(tmpVal []byte) ([]*KlineAttrs, error) {
	vStr := strings.Trim(string(tmpVal), "\"")
	decryptData, err := dataMagic(vStr, key)
	if err != nil {
		return nil, err
	}
	fmt.Println(string(decryptData))
	tmp := Response{}
	err = json.Unmarshal(decryptData, &tmp)
	if err != nil {
		return nil, err
	}
	kline := make([]*KlineAttrs, 0)
	tmpData := make([][]any, 0)
	if len(tmp.Series) != 0 && len(tmp.Series[0].Data) != 0 {
		tmpData = tmp.Series[0].Data
		for _, a := range tmpData[0] {
			if a == nil {
				tmpData = tmpData[1:]
			}
		}
	}
	for _, datum := range tmpData {
		kline = append(kline, &KlineAttrs{
			OpenPrice:  datum[4].(float64),
			HighPrice:  datum[5].(float64),
			LowsPrice:  datum[6].(float64),
			ClosePrice: datum[7].(float64),
			CreatedAt:  int64(datum[0].(float64)),
		})
	}

	return kline, err
}

func dataMagic(b64Data, dk string) ([]byte, error) {
	// 1. Base64 解码
	decodedData, err := base64.StdEncoding.DecodeString(b64Data)
	if err != nil {
		return nil, fmt.Errorf("failed to decode base64: %v", err)
	}

	// 2. 按字节进行异或操作
	keyBytes := []byte(dk)
	for i := 0; i < len(decodedData); i++ {
		decodedData[i] ^= keyBytes[i%len(keyBytes)]
	}

	// 3. 解压数据
	decompressedData, err := decompressGzipData(decodedData)
	if err != nil {
		return nil, fmt.Errorf("failed to decompress encryptionData: %v", err)
	}

	return decompressedData, nil
}

func decompressGzipData(compressedData []byte) ([]byte, error) {
	// 使用 gzip.NewReader 创建解压读取器
	reader, err := gzip.NewReader(bytes.NewReader(compressedData))
	if err != nil {
		return nil, fmt.Errorf("failed to create gzip reader: %v", err)
	}
	defer reader.Close()

	// 使用 bytes.Buffer 来存储解压后的数据
	var decompressedData bytes.Buffer
	_, err = io.Copy(&decompressedData, reader)
	if err != nil {
		return nil, fmt.Errorf("failed to decompress encryptionData: %v", err)
	}

	// 返回解压后的字节数据
	return decompressedData.Bytes(), nil
}
