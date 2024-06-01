package top2

// MoveZeroes 移动零值
func MoveZeroes(nums []int) {
	length := len(nums)
	j := 0
	for i := 0; i < length; i++ {
		if nums[i] != 0 {
			temp := nums[i]
			nums[i] = nums[j]
			nums[j] = temp
			j++
		}
	}
}

// MoveZeroesMy 移动零值
func MoveZeroesMy(nums []int) {
	length := len(nums)
	for i := 0; i < length; i++ {
		if nums[i] == 0 {
			for j := i; j < length; j++ {
				if nums[j] != 0 {
					nums[i] = nums[j]
					nums[j] = 0
					break
				}
			}
		}
	}
}
