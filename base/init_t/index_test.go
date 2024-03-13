package init_t

import (
	"gotest/base/init_t/first_naming"
	"testing"
)

// TestNaming 测试加载init 加载顺序
func TestNaming(t *testing.T) {
	Naming()
}

// TestFirstNaming 测试init加载顺序
func TestFirstNaming(t *testing.T) {
	first_naming.Naming()
}
