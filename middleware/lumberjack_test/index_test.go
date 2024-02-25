package lumberjack_test_test

import (
	"gotest/middleware/lumberjack_test"
	"testing"
)

// TestNewLumberjack 测试lumberjack包
func TestNewLumberjack(t *testing.T) {
	lumberjack_test.NewLumberjack()
}

// TestNewLumberjackAndZap 测试lumberjack包
func TestNewLumberjackAndZap(t *testing.T) {
	lumberjack_test.NewLumberjackAndZap()
}
