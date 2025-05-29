package utils

import "math"

// Round 返回指定精度float,存在四舍五入
func Round(f float64, n int) float64 {
	shift := math.Pow(10, float64(n))
	return math.Round(f*shift) / shift
}
