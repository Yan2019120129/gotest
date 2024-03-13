package tmap

import (
	"fmt"
	"github.com/brianvoe/gofakeit/v6"
	"testing"
	"time"
)

// TestMap 测试map键值存储
func TestMap(t *testing.T) {
	//Map()
	//HB := []string{"EUR", "USD", "USD", "JPY", "GBP", "USD", "AUD", "USD", "USD", "CAD", "USD", "CHF", "NZD", "USD", "USD", "SGD"}
	//one := make(map[string]string)
	//for _, v := range HB {
	//	one[v] = v
	//}

	bb := []string{"EUR", "USD", "GBP", "USD", "AUD", "USD", "NZD", "USD", "USD", "JPY", "CHF", "JPY", "GBP", "JPY", "SGD", "JPY", "EUR", "SGD", "USD", "SGD", "AUD",
		"SGD", "SGD"}
	temp := make(map[string]int)
	for _, s := range bb {
		if _, ok := temp[s]; ok {
			temp[s] += 1
			continue
		}
		temp[s] = 0
	}
	for k, _ := range temp {
		fmt.Println(k)
	}
}

// TestMap 测试map键值存储
func TestForMap(t *testing.T) {
	ForMap()
}

// TestMap 测试map键值存储
func TestIfMap(t *testing.T) {
	IfMap()
}

// TestCopyMap 测试map键值存储
func TestCopyMap(t *testing.T) {
	CopyMap()
}

// TestMapGoroutine 测试map多线程模式下的读写
func TestMapGoroutine(t *testing.T) {
	MapGoroutine()
	time.Sleep(50 * time.Second)
}

// TestSetMap 测试同一实例写入数据会不会存在读写错误
func TestSetMap(t *testing.T) {
	Instance.SetMapValue(gofakeit.Name(), gofakeit.LastName()).SetMapValue(gofakeit.Name(), gofakeit.LastName()).SetMapValue(gofakeit.Name(), gofakeit.LastName()).SetMapValue(gofakeit.Name(), gofakeit.LastName())
}
