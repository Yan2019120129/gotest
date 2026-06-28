package test_go_t

import (
	"runtime"
	"sync"
)

// ForceGC 演示运行时层的 GC：分配一批对象后主动触发一次垃圾回收。
func ForceGC() uint32 {
	before := currentGCCount()
	values := make([][]byte, 0, 128)
	for i := 0; i < 128; i++ {
		values = append(values, make([]byte, 1024))
	}
	runtime.KeepAlive(values)
	runtime.GC()
	return currentGCCount() - before
}

func currentGCCount() uint32 {
	var stats runtime.MemStats
	runtime.ReadMemStats(&stats)
	return stats.NumGC
}

// SchedulerYield 演示 scheduler：Gosched 主动让出当前 goroutine 的执行权。
func SchedulerYield() string {
	runtime.Gosched()
	return "yield"
}

// CountByMap 演示 map：统计字符串出现次数。
func CountByMap(values []string) map[string]int {
	counts := make(map[string]int)
	for _, value := range values {
		counts[value]++
	}
	return counts
}

// SumByGoroutine 演示 goroutine：多个 goroutine 并发计算后汇总结果。
func SumByGoroutine(values []int) int {
	results := make(chan int, len(values))
	var wg sync.WaitGroup
	for _, value := range values {
		value := value
		wg.Add(1)
		go func() {
			defer wg.Done()
			results <- value
		}()
	}

	wg.Wait()
	close(results)

	total := 0
	for value := range results {
		total += value
	}
	return total
}
