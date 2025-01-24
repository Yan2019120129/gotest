package main

import (
	"encoding/json"
	"fmt"
	"gotest/common/utils"
	"math/rand"
	"os"
	"time"
)

func GetKline() {
	nowTime := time.Now()
	rowsAffected := 1
	startTime := nowTime
	startPrice := 100.0
	if rowsAffected > 0 {
		startTime = time.UnixMilli(1735660800000)
		startPrice = 400
	} else {
		startTime = nowTime.AddDate(0, 0, -1)
	}
	afterTime := nowTime.AddDate(0, 0, 1)
	index := int(afterTime.Sub(startTime).Minutes())

	klineList := make([]*utils.KlineAttrs, 0)
	interval := 10
	for i := 0; i < index; i += interval {
		endPrice := floatWithAmplitude(100, 95)
		if interval+i > index {
			interval = index - i
		}
		kLines := utils.GenerateKline(startPrice, endPrice, startTime.Truncate(time.Minute), startTime.Add(time.Duration(interval)*time.Minute).Truncate(time.Minute))
		if len(kLines) > 0 {
			klineList = append(klineList, kLines...)
			startPrice = kLines[len(kLines)-1].ClosePrice
			startTime = startTime.Add(time.Duration(interval) * time.Minute)
			continue
		}
		break
	}

	dataJson, _ := json.Marshal(klineList)
	err := os.WriteFile("/Users/taozijun_1/Documents/vuefile/vuetest/src/assets/kline.json", dataJson, 0644)
	if err != nil {
		return
	}
	fmt.Println(index, len(klineList))
}

// FloatWithAmplitude 根据基本值和振幅范围返回一个浮动后的值
func floatWithAmplitude(base float64, amplitude float64) float64 {
	rand.Seed(time.Now().UnixNano()) // 初始化随机种子
	// 生成范围内的随机浮动值：[-amplitude, +amplitude]
	offset := (rand.Float64()*2 - 1) * amplitude
	return base + offset
}

func NewKline() {
	nowTime := time.Now()
	startTime := nowTime
	targetTime := startTime.Add(50 * time.Minute)
	index := int(targetTime.Sub(startTime).Minutes())
	startPrice := 320.0
	endPrice := startPrice
	price := 250.0
	difference := price - startPrice
	interval := 10
	klineList := make([]*utils.KlineAttrs, 0)
	for i := 0; i < index; i += interval {
		if difference > 100 { // 如果价格差距大于100则根据生成幅度调整
			tmpPrice := difference / float64(index/interval)
			endPrice += tmpPrice
		} else if difference < -100 { // 如果价格差距小于-100则根据生成幅度调整
			tmpPrice := difference / float64(index/interval)
			startPrice -= tmpPrice
		} else {
			interval = index
		}
		if interval+i >= index {
			interval = index - i
		}
		if endPrice > price || interval+i >= index {
			endPrice = price
		}

		kLines := utils.GenerateKline(startPrice, endPrice, startTime.Truncate(time.Minute), startTime.Add(time.Duration(interval)*time.Minute).Truncate(time.Minute))
		if len(kLines) > 0 {
			klineList = append(klineList, kLines...)
			startPrice = kLines[len(kLines)-1].ClosePrice
			startTime = startTime.Add(time.Duration(interval) * time.Minute)
			continue
		}
		break
	}
	dataJson, err := json.Marshal(&klineList)
	if err != nil {
		panic(err)
	}
	err = os.WriteFile("/Users/taozijun_1/Documents/vuefile/vuetest/src/assets/kline.json", dataJson, 0644)
	if err != nil {
		return
	}
	fmt.Println(index, len(klineList))
}

func main() {
	//GetKline()
	NewKline()
}
