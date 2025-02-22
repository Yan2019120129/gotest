package main

import (
	"encoding/json"
)

// RequestPayload 定义请求结构
type RequestPayload struct {
	Jsonrpc string      `json:"jsonrpc"`
	ID      int         `json:"id"`
	Method  string      `json:"method"`
	Params  interface{} `json:"params,omitempty"`
}

// Transaction 定义交易结构
type Transaction struct {
	Transaction struct {
		Signatures []string `json:"signatures"`
		Message    struct {
			AccountKeys  []string `json:"accountKeys"`
			Instructions []struct {
				ProgramIDIndex int    `json:"programIdIndex"`
				Data           string `json:"data"`
				Accounts       []int  `json:"accounts"`
			} `json:"instructions"`
		} `json:"message"`
	} `json:"transaction"`
	Meta struct {
		Err interface{} `json:"err"`
	} `json:"meta"`
}

// BlockResponse 定义区块响应结构
type BlockResponse struct {
	Result struct {
		Transactions []Transaction `json:"transactions"`
	} `json:"result"`
}

type Response struct {
	Result json.RawMessage `json:"result"`
}

//
