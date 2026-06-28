package sort_t

// MergeSort returns a sorted copy of nums using merge sort.
func MergeSort(nums []int) []int {
	if len(nums) <= 1 {
		result := make([]int, len(nums))
		copy(result, nums)
		return result
	}

	mid := len(nums) / 2
	left := MergeSort(nums[:mid])
	right := MergeSort(nums[mid:])

	return merge(left, right)
}

func merge(left, right []int) []int {
	result := make([]int, 0, len(left)+len(right))
	i, j := 0, 0

	for i < len(left) && j < len(right) {
		if left[i] <= right[j] {
			result = append(result, left[i])
			i++
			continue
		}

		result = append(result, right[j])
		j++
	}

	result = append(result, left[i:]...)
	result = append(result, right[j:]...)

	return result
}
