//go:build ignore

package examples_disabled

// 以下三个指令属于 Go runtime 内部写屏障控制。
// 普通包中使用会被编译器拒绝，业务代码不应该使用。

// go:nowritebarrier 禁止函数中出现写屏障。
//
//go:nowritebarrier
func runtimeOnlyNoWriteBarrier() {
}

// go:nowritebarrierrec 禁止当前函数及其递归调用链出现写屏障。
//
//go:nowritebarrierrec
func runtimeOnlyNoWriteBarrierRec() {
}

// go:yeswritebarrierrec 用于结束 runtime 中 nowritebarrierrec 的递归约束范围。
//
//go:yeswritebarrierrec
func runtimeOnlyYesWriteBarrierRec() {
}
