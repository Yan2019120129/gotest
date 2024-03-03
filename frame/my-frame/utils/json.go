package utils

import (
	"encoding/json"
	"log"
)

// ObjToByteList 结构体转换byte数组
func ObjToByteList(params interface{}) []byte {
	if data, err := json.Marshal(params); err != nil {
		return nil
	} else {
		return data
	}
}

// ByteListToObj byte数组转换结构体
func ByteListToObj(params []byte, obj interface{}) {
	if err := json.Unmarshal(params, obj); err != nil {
		log.Println(err)
	}
}
