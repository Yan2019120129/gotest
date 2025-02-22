package main

import (
	"fmt"
	"gotest/common/utils"
	"gotest/practice/okx1/dto"
	"gotest/practice/okx1/enum"
)

func main() {
	okxWs := utils.NewWs(enum.TradesWsUrlOKX)
	subParams := dto.NewSubscribe().
		SetSubParams(enum.OKXChannelTrades, "DOGE-USDT").
		Subscribe()
	okxWs.Run()
	okxWs.Send(subParams.ToString())
	okxWs.Read(func(bytes []byte) {
		fmt.Println(string(bytes))
	})
}
