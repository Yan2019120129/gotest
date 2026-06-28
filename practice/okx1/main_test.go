package main

import (
	"encoding/json"
	"fmt"
	"gotest/common/utils"
	"gotest/practice/okx1/monitor"
	"log"
	"os"
	"strconv"
	"testing"
	"time"
)

type KlineResp struct {
	Code string     `json:"code"`
	Msg  string     `json:"msg"`
	Data [][]string `json:"data"`
}

func TestPriceMonitor(m *testing.T) {
	mailClient := &mockEmailSender{}
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

	data, err := os.ReadFile("./tmp.json")
	if err != nil {
		log.Fatal(err)
	}

	var resp KlineResp

	if err := json.Unmarshal(data, &resp); err != nil {
		log.Fatal(err)
	}

	for _, item := range resp.Data {

		if len(item) < 5 {
			continue
		}

		price, err := strconv.ParseFloat(
			item[4], // close价格
			64,
		)

		if err != nil {
			continue
		}

		latestPrice = price

		fmt.Printf(
			"模拟行情 %.2f\n",
			price,
		)

		pm.Add(price)

		// 模拟1分钟K线推送
		time.Sleep(time.Second)
	}
}

type mockEmailSender struct{}

func (m *mockEmailSender) Send(to []string, subject, body string) error {
	fmt.Println("=== 模拟发送邮件 ===")
	fmt.Println("收件人:", to)
	fmt.Println("标题:", subject)
	fmt.Println("内容:\n", body)
	fmt.Println("===================")
	return nil
}

func TestPrice(t *testing.T) {
	h := utils.NewHttp()
	v, e := h.Get("https://www.okx.com/api/v5/market/index-candles?instId=BTC-USDT")
	if e != nil {
		panic(e)
	}
	fmt.Println(string(v))
}

func TestAccountBalance(t *testing.T) {
	client := monitor.NewClient(
		"589978b6-8b1a-40c9-89ee-c33403451cf7",
		"71EFC47547FA4D820BDA97D2D676F5ED",
		"Yjj1323106558.",
	)

	// 获取全部资产
	res, err := client.GetAccountBalance("ALLO")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(res)

	fmt.Println("账户总权益:", res.Data[0].TotalEq)

	for _, d := range res.Data[0].Details {
		fmt.Printf("币种: %s 余额: %s 可用: %s\n",
			d.Ccy, d.CashBal, d.AvailBal)
	}
}
