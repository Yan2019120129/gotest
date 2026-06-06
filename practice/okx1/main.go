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
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/natefinch/lumberjack"
)

func main() {
	fmt.Println("./okx config.yml BTC-USDT,DOGE-USDT")
	defaultConfigPath := "/home/yan/Documents/file/gofile/gotest/common/config/config.yml"
	defaultSymbols := []string{"ETH-USDT"}

	args := os.Args[1:]

	var (
		configPath string
		symbols    []string
	)

	// 解析参数
	switch len(args) {
	case 0:
		configPath = defaultConfigPath
		symbols = defaultSymbols

	case 1:
		// 只传一个参数：可能是 config，也可能是 symbol
		if strings.Contains(args[0], ".yml") || strings.Contains(args[0], ".yaml") {
			configPath = args[0]
			symbols = defaultSymbols
		} else {
			configPath = defaultConfigPath
			symbols = strings.Split(args[0], ",")
		}

	default:
		configPath = args[0]
		symbols = strings.Split(args[1], ",")
	}

	log.SetOutput(&lumberjack.Logger{
		Filename:   "./logs/app.log",
		MaxSize:    50, // MB
		MaxBackups: 1,  // 保留1个备份
		MaxAge:     1,  // 保留1天
		Compress:   false,
	})

	config.Init(configPath)
	cfg := config.GetConfig()

	emailCfg := cfg.Mail.Providers[cfg.Mail.Default]
	mailClient := email.NewClient(
		emailCfg.Host,
		emailCfg.Port,
		emailCfg.Username,
		emailCfg.Password,
		"行情监控",
	)

	// 最新价格缓存（多 symbol）
	latestPrice := make(map[string]float64)

	// 每分钟采样
	go func() {
		ticker := time.NewTicker(time.Minute)
		defer ticker.Stop()

		priceMonitor := make(map[string]*monitor.PriceMonitor)
		for _, symbol := range symbols {
			pm := monitor.NewPriceMonitor(
				symbol,
				monitor.DefaultConfigs[symbol],
				mailClient,
				[]string{"1556403682@qq.com"},
			)
			priceMonitor[symbol] = pm
		}

		for range ticker.C {
			for _, symbol := range symbols {
				price := latestPrice[symbol]
				if price <= 0 {
					continue
				}

				if pm, ok := priceMonitor[symbol]; ok {
					pm.Add(price)
					log.Printf("采样 %s: %f\n", symbol, price)
				}
			}
		}
	}()

	wsConfig := make(map[string]any)
	if cfg.Gin.Proxy != "" {
		wsConfig["proxy"] = cfg.Gin.Proxy
	}

	okxWs := utils.NewWs(enum.TradesWsUrlOKX, wsConfig)

	okxWs.Run()

	// 为每个 symbol 订阅
	for _, symbol := range symbols {
		subParams := dto.NewSubscribe().
			SetSubParams(enum.OKXChannelTicker, symbol).
			Subscribe()

		okxWs.Send(subParams.ToString())
	}

	// websocket 回调
	okxWs.Read(func(msg []byte) {
		var resp dto.RespJson
		if err := json.Unmarshal(msg, &resp); err != nil {
			return
		}

		if len(resp.Data) == 0 {
			return
		}

		var okxTickers []dto.OkxTickers
		if err := json.Unmarshal(resp.Data, &okxTickers); err != nil {
			return
		}

		price, err := strconv.ParseFloat(okxTickers[0].Last, 64)
		if err != nil {
			return
		}

		symbol := okxTickers[0].InstId
		latestPrice[symbol] = price

		log.Printf("%s: %f\n", symbol, price)
	})
}
