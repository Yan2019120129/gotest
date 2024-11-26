package main

import (
	"fmt"
)

func main() {
	value := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	fmt.Println(groupAnagrams(value))

}

func groupAnagrams(strs []string) [][]string {
	res := make([][]string, 0)
	mp := map[[26]int]int{}
	for _, v := range strs {
		var word = [26]int{}
		for _, s := range v {
			word[s-'a']++
		}
		if i, ok := mp[word]; ok {
			res[i] = append(res[i], v)
		} else {
			mp[word] = len(res)
			res = append(res, []string{v})
		}
	}

	return res
}
