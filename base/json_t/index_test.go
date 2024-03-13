package json_t_test

import (
	"gotest/base/json_t"
	"testing"
)

// TestJson 测试json 字节转换
func TestTestJson(t *testing.T) {
	json_t.TestJson()
}

// TestStringToJson string转json
func TestStringToJson(t *testing.T) {
	json_t.StringToJson()
}

// TestInterfaceToObj interface转Obj
func TestInterfaceToObj(t *testing.T) {
	json_t.InterfaceToObj()
}
