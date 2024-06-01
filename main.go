package main

var sqlPaths = []string{
	"./common/file/productsql/product.sql",
	"./common/file/productsql/product_attrs_key.sql",
	"./common/file/productsql/product_attrs_sku.sql",
	"./common/file/productsql/product_attrs_val.sql"}

func main() {
	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	groupAnagrams(strs)
}

// 字母异位词分组
func groupAnagrams(arr []string) [][]string {

	return nil
}

// binarySorting 二分排序
func binarySorting(s string) string {
	sLen := len(s)
	if sLen == 1 || sLen == 0 {
		return s
	}

	return s
}
