package top6

import "fmt"

// r^n = r^(n-1) + r^(n-2)

// MaximumNumberOfStringPairs 最大字符串匹配问题
func MaximumNumberOfStringPairs(words []string) int {
	tempMap := map[string]string{}
	count := 0
	for _, v := range words {
		temp := ""
		for _, s := range v {
			temp = fmt.Sprintf("%c", s) + temp
		}
		tempMap[v] = temp
		if v == temp {
			continue
		}
		if _, ok := tempMap[temp]; ok {
			count++
		}
	}
	return count
}

// MaximumNumberOfStringPairsTow 最大字符串匹配问题
func MaximumNumberOfStringPairsTow(words []string) int {
	length := len(words)
	count := 0
	for i := 0; i < length-1; i++ {
		for j := i + 1; j < length; j++ {
			if words[i][0] == words[j][1] && words[i][1] == words[j][0] {
				count++
			}
		}
	}
	return count
}

// MaximumNumberOfStringPairsFrid 最大字符串匹配问题
func MaximumNumberOfStringPairsFrid(words []string) int {
	ans := 0
	seen := map[int]int{}
	for _, word := range words {
		fmt.Println("v:", int(word[1])*100+int(word[0]))
		ans += seen[int(word[1])*100+int(word[0])]
		seen[int(word[0])*100+int(word[1])] = 1
	}
	return ans
}
