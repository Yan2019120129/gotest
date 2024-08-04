package top14

import "math"

// 判断一个数是否为质数
func isPrime(n int) bool {
	// 小于等于1的数不是质数
	if n <= 1 {
		return false
	}
	// 2和3是质数
	if n <= 3 {
		return true
	}
	// 能被2或3整除的数不是质数
	if n%2 == 0 || n%3 == 0 {
		return false
	}
	// 检查是否能被形如6k±1的数整除
	for i := 5; i*i <= n; i += 6 {
		if n%i == 0 || n%(i+2) == 0 {
			return false
		}
	}
	return true
}

// 判断区间特殊数字外的数字总数
func nonSpecialCount(l, r int) int {
	startSqrt := int(math.Sqrt(float64(l)))
	endSrqt := int(math.Sqrt(float64(r)))

	count := 0
	if startSqrt == 0 || endSrqt == 0 {
		if isPrime(startSqrt) {
			count++
		}
		if isPrime(endSrqt) {
			count++
		}
	} else {
		for i := startSqrt; i <= endSrqt; i++ {
			if isPrime(i) {
				count++
			}
		}
	}

	return r - l + 1 - count
}

//// nonSpecialCount 优质题解,根据条件1 <= l <= r <= 10^9
//const mx = 31622 // 通过计算得int(math.Sqrt(math.Pow10(9)))
//
//// 具体思路是生成质数表，通过埃筛的方式生成
//var pi = [mx + 1]int{}
//
//func init() {
//	for i := 2; i <= mx; i++ {
//		if pi[i] == 0 { // i 是质数
//			pi[i] = pi[i-1] + 1
//			for j := i * i; j <= mx; j += i {
//				pi[j] = -1 // 标记 i 的倍数为合数
//			}
//		} else {
//			pi[i] = pi[i-1]
//		}
//	}
//}
//
//func nonSpecialCount(l, r int) int {
//	cntR := pi[int(math.Sqrt(float64(r)))]
//	cntL := pi[int(math.Sqrt(float64(l-1)))]
//	return r - l + 1 - (cntR - cntL)
//}
