package top13

// 128.最长连续序列
func longestConsecutive(nums []int) int {
	sortMap := make(map[int]bool)
	for _, num := range nums {
		sortMap[num] = true
	}
	numsLen := 0
	for k := range sortMap {
		if !sortMap[k-1] {
			j := k
			i := 1
			for sortMap[j+1] {
				i++
				j++
			}
			if numsLen < i {
				numsLen = i
			}
		}
	}
	return numsLen
}
