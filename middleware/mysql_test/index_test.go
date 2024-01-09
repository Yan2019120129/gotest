package mysql_test_test

import (
	"gotest/middleware/mysql_test"
	"testing"
)

// TestWhereId 测试where 条件
func TestWhereId(t *testing.T) {
	mysql_test.TestWhereId()
}

// TestWhereIds 测试where 自动映射id数组条件
func TestWhereIds(t *testing.T) {
	mysql_test.TestWhereIds()
}

// TestWhereOtherIds 测试where 其他数组条件
func TestWhereOtherIds(t *testing.T) {
	mysql_test.TestWhereOtherIds()
}
