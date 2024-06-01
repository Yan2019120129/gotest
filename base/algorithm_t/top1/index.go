package top1

// TwoSum 两数之和
func TwoSum(nums []int, target int) []int {
	sumMap := map[int]int{}
	for i, num := range nums {
		if p, ok := sumMap[target-num]; ok {
			return []int{p, i}
		}
		sumMap[num] = i
	}
	return nil
}
