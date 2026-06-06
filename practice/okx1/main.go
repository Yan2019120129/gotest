package main

import (
	"encoding/json"
	"fmt"
	"gotest/common/config"
	"gotest/common/utils"
	"gotest/practice/email"
	"gotest/practice/okx1/dto"
	"gotest/practice/okx1/enum"
	"gotest/practice/okx1/monitor"
	"log"
	"strconv"
	"time"
)

func main() {
	cfg := config.GetMailConfig()
	emailCfg := cfg.Providers[cfg.Default]
	mailClient := email.NewClient(
		emailCfg.Host,
		emailCfg.Port,
		emailCfg.Username,
		emailCfg.Password,
		"行情监控",
	)

	symbol := "ETH-USDT"
	// 监控器
	pm := monitor.NewPriceMonitor(
		symbol,
		monitor.DefaultConfigs[symbol],
		mailClient,
		[]string{
			"1556403682@qq.com",
		},
	)

	var latestPrice float64

	// 每分钟采样一次
	go func() {

		ticker := time.NewTicker(time.Minute)
		defer ticker.Stop()

		for range ticker.C {

			if latestPrice <= 0 {
				continue
			}

			pm.Add(latestPrice)

			log.Printf(
				"采样价格: %.2f\n",
				latestPrice,
			)
		}
	}()

	okxWs := utils.NewWs(enum.TradesWsUrlOKX, map[string]any{
		"proxy": "http://127.0.0.1:7890",
	})
	subParams := dto.NewSubscribe().
		SetSubParams(enum.OKXChannelTicker, symbol).
		Subscribe()
	okxWs.Run()
	okxWs.Send(subParams.ToString())
	okxWs.Read(func(msg []byte) {
		var resp dto.RespJson

		if err := json.Unmarshal(msg, &resp); err != nil {
			return
		}

		var okxTickers []dto.OkxTickers
		if err := json.Unmarshal(resp.Data, &okxTickers); err != nil {
			return
		}

		if len(resp.Data) == 0 {
			return
		}

		price, err := strconv.ParseFloat(
			okxTickers[0].Last,
			64,
		)

		if err != nil {
			return
		}

		latestPrice = price

		fmt.Printf(
			symbol+": %.2f\n",
			price,
		)
	})
}
