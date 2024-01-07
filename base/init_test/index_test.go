package init_test_test

import (
	"gotest/base/init_test"
	"gotest/base/init_test/first_naming"
	"testing"
)

// TestNaming 测试加载init 加载顺序
func TestNaming(t *testing.T) {
	init_test.Naming()
}

// TestFirstNaming 测试init加载顺序
func TestFirstNaming(t *testing.T) {
	first_naming.Naming()
}
