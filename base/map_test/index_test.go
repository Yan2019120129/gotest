package map_test_test

import (
	"github.com/brianvoe/gofakeit/v6"
	"gotest/base/map_test"
	"testing"
	"time"
)

// TestMap 测试map键值存储
func TestMap(t *testing.T) {
	map_test.Map()
}

// TestMap 测试map键值存储
func TestForMap(t *testing.T) {
	map_test.ForMap()
}

// TestMap 测试map键值存储
func TestIfMap(t *testing.T) {
	map_test.IfMap()
}

// TestCopyMap 测试map键值存储
func TestCopyMap(t *testing.T) {
	map_test.CopyMap()
}

// TestMapGoroutine 测试map多线程模式下的读写
func TestMapGoroutine(t *testing.T) {
	map_test.MapGoroutine()
	time.Sleep(50 * time.Second)
}

// TestSetMap 测试同一实例写入数据会不会存在读写错误
func TestSetMap(t *testing.T) {
	map_test.Instance.SetMapValue(gofakeit.Name(), gofakeit.LastName()).SetMapValue(gofakeit.Name(), gofakeit.LastName()).SetMapValue(gofakeit.Name(), gofakeit.LastName()).SetMapValue(gofakeit.Name(), gofakeit.LastName())
}
