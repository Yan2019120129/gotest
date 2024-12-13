package main

import (
	"fmt"
	"time"
)

var (
	address = []string{
		"4aMBgyyHoe8nFw81H9yZmDXP95CkmH1WGh5DQGNGpump",
		"7imiPydi4TwAHYNWeRVdmNcmdeVD5k7JqFXTiZgL18ce",
	}

	urls = []string{
		fmt.Sprintf("https://api-v2.solscan.io/v2/account?address=%s", address[0]),
		fmt.Sprintf("https://api-v2.solscan.io/v2/token/transfer?address=%s&page=%d&page_size=%d&exclude_amount_zero=%t", address[0], 1, 10, false),
		fmt.Sprintf("https://gmgn.ai/defi/quotation/v1/wallet_activity/sol?type=buy&wallet=%s&limit=%d&cost=%d", address[1], 1, 10),
		"https://api.devnet.solana.com",
		"wss://ws.gmgn.ai/stream",
		"wss://frontend-api-v2.pump.fun/socket.io/?EIO=4&transport=websocket",
	}

	megs = []string{
		`{"action":"subscribe","channel":"token_activity","id":"de301ffa0c0eee33","data":{"chain":"sol","address":"6m2SeubvmKc1MGqGxQToy9i7tdbkxhnMxqgEhtgFpump"}}`,
		`40`,
		"3",
	}
)

func main() {
	instance := NewWs(urls[5], "")
	instance.Send(megs[1])
	go func() {
		for {
			time.Sleep(1)
			instance.Send(megs[2])
		}
	}()
	instance.Read(func(val []byte) {
		fmt.Println(string(val))
	})

	//data := utils.ObjToString(map[string]interface{}{
	//	"jsonrpc": "2.0",
	//	"id":      1,
	//	"method":  "getBlocks",
	//	"params": []any{
	//		5, 10,
	//	},
	//})
	//val := NewHttp().Post(urls[2], "application/json", data)
	//fmt.Println(string(val))

	//val := NewHttp().
	//	Get(urls[2])
	//fmt.Println(string(val))

	//// 获取最新的 slot
	//solana := NewSolan(urls[2])
	//latestSlot, err := solana.getLatestSlot()
	//if err != nil {
	//	fmt.Printf("Error getting latest slot: %v\n", err)
	//	return
	//}
	//
	//// 获取区块数据
	//blockData, err := solana.getBlock(latestSlot)
	//if err != nil {
	//	fmt.Printf("Error getting block data: %v\n", err)
	//	return
	//}
	//
	//// 解析交易数据
	//solana.parseTransactions(blockData)
}
