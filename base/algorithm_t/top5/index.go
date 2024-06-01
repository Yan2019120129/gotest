package top5

import "math"

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
