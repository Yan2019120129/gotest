package map_test

import (
	"fmt"
	"github.com/brianvoe/gofakeit/v6"
	"maps"
	"time"
)

// Map 测试map键值输入输出
func Map() {
	data := map[int]int{2: 1, 5: 7}
	fmt.Println("data:", data[2])
	if p, ok := data[5]; ok {
		fmt.Println("p:", p)
		fmt.Println("ok:", ok)
	}
}

// ForMap 测试map for 遍历
func ForMap() {
	mapData := map[string]interface{}{"age": 18, "name": "yan", "sex": "man"}
	// k为键，v为值
	for k, v := range mapData {
		fmt.Println("k:", k)
		fmt.Println("v:", v)
	}

	// i 输出的是键
	for i := range mapData {
		fmt.Println("i:", i)
	}
}

// IfMap 测试map for 遍历
func IfMap() {
	mapData := map[string]interface{}{"age": 18, "name": "yan", "sex": "man"}
	// v为值，ok判断是否存在，存在为true，否为false
	if v, ok := mapData["age"]; ok {
		fmt.Println("ok:", ok)
		fmt.Println("v:", v)
	}
	fmt.Println(len(mapData))
}

// CopyMap 测试复制map
func CopyMap() {
	mapData := map[string]interface{}{gofakeit.LastName(): gofakeit.Name(), gofakeit.LastName(): "yan", gofakeit.LastName(): "man"}
	mapDataTow := map[string]interface{}{gofakeit.LastName(): 18, gofakeit.LastName(): gofakeit.Name(), gofakeit.LastName(): "man"}
	// v为值，ok判断是否存在，存在为true，否为false
	tempMap := map[string]interface{}{}
	maps.Copy(tempMap, mapData)
	maps.Copy(tempMap, mapDataTow)
	fmt.Println(tempMap)
}

// MapGoroutine 测试map 多线程读写问题
func MapGoroutine() {
	mapData := map[int]interface{}{}
	for i := 0; i < 100; i++ {
		mapData[gofakeit.Number(0, 100)] = gofakeit.Name()
	}

	findMap := func() {
		for {
			key := gofakeit.Number(0, 1000)
			//if _, ok := mapData[key]; !ok {
			//	mapData[key] = gofakeit.Name()
			//}
			//if _, ok := mapData[key]; ok {
			//	fmt.Println(mapData[key])
			//}

			if mapData[key] == nil {
				mapData[key] = gofakeit.Name()
			}
			time.Sleep(500 * time.Millisecond)
		}
	}

	for i := 0; i < 50; i++ {
		go findMap()
	}

}
