package slice_test

import "fmt"

// TestSliceForDelete 测试切片循环删除是否会出现问题
func TestSliceForDelete() {
	fields := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}
	for i, field := range fields {
		if field%3 == 0 {
			fields = append(fields[:i], fields[i+1:]...)
		}
	}
	fmt.Println("fields:", fields)
}
