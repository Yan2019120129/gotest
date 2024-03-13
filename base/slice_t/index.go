package slice_t

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

// SliceForDeleteOne 删除数组中的元素
func SliceForDeleteOne() {
	fields := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}
	for i, field := range fields {
		if field%3 == 0 {
			copy(fields[i:], fields[i+1:])
			fields = fields[:len(fields)-1]
		}
	}
	fmt.Println("fields:", fields)
}
