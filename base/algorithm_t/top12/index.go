package top12

// 49. 字母异位词分组
//func groupAnagrams(strs []string) [][]string {
//	tmpInt := make(map[string]int)
//	data := make([][]string, 0)
//	j := 0
//	for _, str := range strs {
//		bytes := []byte(str)
//		sort.Slice(bytes, func(i, j int) bool {
//			return bytes[i] < bytes[j]
//		})
//		key := string(bytes)
//		if _, ok := tmpInt[key]; !ok {
//			data = append(data, []string{str})
//			tmpInt[key] = j
//			j++
//			continue
//		}
//		data[tmpInt[key]] = append(data[tmpInt[key]], str)
//	}
//	return data
//}

// 49. 字母异位词分组 优化
func groupAnagrams(strs []string) [][]string {
	mp := map[[26]int][]string{}
	for _, value := range strs {
		a := [26]int{}
		for _, val := range value {
			a[val-'a']++
		}
		mp[a] = append(mp[a], value)
	}
	ans := [][]string{}
	for _, value := range mp {
		ans = append(ans, value)
	}
	return ans
}
