package test_go_t

import (
	"runtime"
	"testing"
	"unsafe"
)

func TestNoInlineAndNoSplit(t *testing.T) {
	if got := NoInlineAdd(1, 2); got != 3 {
		t.Fatalf("go:noinline 示例结果错误: %d", got)
	}

	if got := NoSplitAdd(3, 4); got != 7 {
		t.Fatalf("go:nosplit 示例结果错误: %d", got)
	}

	if got := NoEscapeAdd(5, 6); got != 11 {
		t.Fatalf("go:noescape 示例结果错误: %d", got)
	}

	value := 100
	addr := UintptrEscapesAddress(uintptr(unsafe.Pointer(&value)))
	runtime.KeepAlive(&value)
	if addr == 0 {
		t.Fatal("go:uintptrescapes 示例返回了空地址")
	}
}

func TestLinknameHiddenToken(t *testing.T) {
	if got := LinknameHiddenToken("u1001"); got != "token:u1001" {
		t.Fatalf("go:linkname 示例结果错误: %s", got)
	}
}
