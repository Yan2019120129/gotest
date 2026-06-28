package test_go_t

import "testing"

func TestRuntimeGC(t *testing.T) {
	if got := ForceGC(); got == 0 {
		t.Fatal("GC 示例没有观察到 NumGC 增长")
	}
}

func TestRuntimeScheduler(t *testing.T) {
	if got := SchedulerYield(); got != "yield" {
		t.Fatalf("scheduler 示例结果错误: %s", got)
	}
}

func TestRuntimeMap(t *testing.T) {
	counts := CountByMap([]string{"paid", "pending", "paid"})
	if counts["paid"] != 2 || counts["pending"] != 1 {
		t.Fatalf("map 统计结果错误: %#v", counts)
	}
}

func TestRuntimeGoroutine(t *testing.T) {
	if got := SumByGoroutine([]int{1, 2, 3, 4}); got != 10 {
		t.Fatalf("goroutine 汇总结果错误: %d", got)
	}
}
