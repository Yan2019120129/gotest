package algorithm

import (
	"fmt"
	"math"
)

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

// -------------------------------------
// 假设你正在爬楼梯。需要 n 阶你才能到达楼顶。
// 每次你可以爬 1 或 2 个台阶。你有多少种不同的方法可以爬到楼顶呢？
// usedStairs 存储计算过的结果
var usedStairs = make(map[int]int)

// ClimbStairsRecursion 爬楼梯 递归解法
func ClimbStairsRecursion(n int) int {
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	if v, ok := usedStairs[n]; ok {
		return v
	} else {
		result := ClimbStairsRecursion(n-1) + ClimbStairsRecursion(n-2)
		usedStairs[n] = result
		return result
	}
}

// -------------------------------------

// ClimbStairs 使用滚动数组的方式求解
func ClimbStairs(n int) int {
	p, q, r := 0, 0, 1
	for i := 1; i <= n; i++ {
		p = q
		q = r
		r = q + p
	}
	return r
}

// -------------------------------------
type matrix [2][2]int

// ClimbStarMatrix 矩阵快速幂方式
func ClimbStarMatrix(n int) int {
	res := pow(matrix{{1, 1}, {1, 0}}, n)
	return res[0][0]
}

func mul(a, b matrix) (c matrix) {
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			c[i][j] = a[i][0]*b[0][j] + a[i][1]*b[1][j]
		}
	}
	return c
}

func pow(a matrix, n int) matrix {
	res := matrix{{1, 0}, {0, 1}}
	for ; n > 0; n >>= 1 {
		if n&1 == 1 {
			res = mul(res, a)
		}
		a = mul(a, a)
	}
	return res
}

// -------------------------------------

// ClimbStairMath 数学公式解法
func ClimbStairMath(n int) int {
	sqrt5 := math.Sqrt(5)
	pow1 := math.Pow((1+sqrt5)/2, float64(n+1))
	pow2 := math.Pow((1-sqrt5)/2, float64(n+1))
	return int(math.Round((pow1 - pow2) / sqrt5))
}

// -------------------------------------

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

// 输入: nums = [0,1,0,3,12]
// 输出: [1,3,12,0,0]

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

// Merge 二路归并算法
func Merge(array, arrayTow []int, low, m, height int) {
	//i := low
	//j := m + 1
	//k := low
	//for (i <= m) && (j <= height) {
	//	if array[i]
	//}
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func GetIntersectionNode(headA, headB *ListNode) *ListNode {
	values, lenght, tempMap := searchValues(headA)
	values1, lenght1, _ := searchValues(headB)
	tempLen := lenght
	if lenght > lenght1 {
		tempLen = lenght1
	}
	index := ""
	for i := 1; i <= tempLen; i++ {
		if values1[lenght1-i] == values[lenght-i] {
			index = values1[lenght1-i]
			continue
		}
	}
	return tempMap[index]
}

func searchValues(headB *ListNode) ([]string, int, map[string]*ListNode) {
	mapNode := make(map[string]*ListNode)
	values := make([]string, 0)
	for headB != nil {
		key := fmt.Sprintf("%p", headB)
		mapNode[key] = headB
		values = append(values, key)
		headB = headB.Next
	}

	return values, len(values), mapNode
}
