package map_test

import "fmt"

// TestMap 测试map键值输入输出
func TestMap() {
	data := map[int]int{2: 1, 5: 7}
	fmt.Println("data:", data[2])
	if p, ok := data[5]; ok {
		fmt.Println("p:", p)
		fmt.Println("ok:", ok)
	}
}

// TestForMap 测试map for 遍历
func TestForMap() {
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

// TestIfMap 测试map for 遍历
func TestIfMap() {
	mapData := map[string]interface{}{"age": 18, "name": "yan", "sex": "man"}
	// v为值，ok判断是否存在，存在为true，否为false
	if v, ok := mapData["age"]; ok {
		fmt.Println("ok:", ok)
		fmt.Println("v:", v)
	}
}
