package top9

// RemoveElement 移除元素
func RemoveElement(nums []int, val int) int {
	length := len(nums)
	number := 0
	for i := 0; i < length; i++ {
		if nums[i] == val {
			nums = append(nums[:i], nums[i+1:]...)
			number++
		}
	}
	return length - number
}
