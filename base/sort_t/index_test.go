package sort_t

import (
	"fmt"
	"sort"
	"testing"
)

// sortString 字符串排序
func TestSortString(t *testing.T) {
	strings := []string{"yan", "jia", "jie", "yang", "chuan", "song"}
	sort.Strings(strings)
	for _, s := range strings {
		fmt.Println(s)
	}
}
