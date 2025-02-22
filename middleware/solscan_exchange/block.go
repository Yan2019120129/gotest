package main

import (
	"encoding/json"
	"fmt"
	"gotest/common/utils"
	"log"
)

type Solana struct {
	url string
}

func NewSolan(u string) *Solana {
	return &Solana{url: u}
}

// 调用 Solana JSON-RPC API
func (s *Solana) callSolanaRPC(method string, params interface{}) ([]byte, error) {
	payload := RequestPayload{
		Jsonrpc: "2.0",
		ID:      1,
		Method:  method,
		Params:  params,
	}
	param := utils.ObjToString(payload)
	body := NewHttp().Post(s.url, "application/json", param)
	return body, nil
}

// 获取最新 slot
func (s *Solana) getLatestSlot() (int, error) {
	body, err := s.callSolanaRPC("getSlot", nil)
	if err != nil {
		return 0, err
	}

	var slot = 0

	s.JsonConvert(body, slot)
	return slot, nil
}

// getMaxAccounts 获取最大账户列表
func (s *Solana) getMaxAccounts() {
	body, err := s.callSolanaRPC("getLargestAccounts", "")
	if err != nil {
		return
	}
	var slot = 0
	s.JsonConvert(body, &slot)
}

// 获取区块数据
func (s *Solana) getBlock(slot int) ([]Transaction, error) {
	params := []interface{}{
		slot,
		map[string]interface{}{
			"encoding":                       "json",
			"maxSupportedTransactionVersion": 0,
			"transactionDetails":             "full",
			"rewards":                        false,
		},
	}

	body, err := s.callSolanaRPC("getBlock", params)
	if err != nil {
		return []Transaction{}, err
	}

	val := make([]Transaction, 0)
	s.JsonConvert(body, val)
	return val, nil
}

// JsonConvert 数据转换
func (s *Solana) JsonConvert(body []byte, val any) {
	var response Response
	if err := json.Unmarshal(body, &response); err != nil {
		log.Fatal(err)
		return
	}

	if err := json.Unmarshal(response.Result, &val); err != nil {
		log.Fatal(err)
		return
	}
}

// 解析交易数据
func (s *Solana) parseTransactions(transactions []Transaction) {
	for idx, tx := range transactions {
		status := "Success"
		if tx.Meta.Err != nil {
			status = "Failed"
		}
		// 获取账户key
		accountKeys := tx.Transaction.Message.AccountKeys

		fmt.Printf("Transaction %d:\n", idx+1)
		fmt.Printf("  Signatures: %v\n", tx.Transaction.Signatures)
		fmt.Printf("  Instructions:\n")
		for _, instr := range tx.Transaction.Message.Instructions {
			fmt.Printf("		ProgramID:%s\n", accountKeys[instr.ProgramIDIndex])
			fmt.Printf("		Accounts:\n")
			fmt.Println(len(instr.Accounts), instr.Accounts)
			//for _, accIdx := range instr.Accounts {
			//	fmt.Printf("    		  - Accounts:%s \n", accountKeys[accIdx])
			//}
			fmt.Printf("		Data: %s\n", instr.Data)
		}
		fmt.Printf("  Status: %s\n\n", status)
		fmt.Println("		----------------------------------------------------------")
	}
}
