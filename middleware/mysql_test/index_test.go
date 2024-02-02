package mysql_test_test

import (
	"fmt"
	"gorm.io/gorm/utils"
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

// TestGormMapping  测试gorm是否能映射string 为[]string 类型
func TestGormMapping(t *testing.T) {
	mysql_test.TestGormMapping()
}

// TestUpdated 测试gorm update 回不会自动更新时间
func TestTestUpdated(t *testing.T) {
	mysql_test.TestUpdated()
}

//

// TestGormToString 测试gorm 中ToString方法
func TestGormToString(t *testing.T) {
	fmt.Println("data:", utils.ToString([]string{"/assets/currency/ada.png"}))
}

// 验证使用Model和不实用的区别
func TestGormFind(t *testing.T) {
	mysql_test.TestGormFind()
}

// TestSelectClient 测试子查询
func TestSelectClient(t *testing.T) {
	mysql_test.TestSelectClient()
}

// TestGormInsert 添加测试数据
func TestGormInsert(t *testing.T) {
	mysql_test.TestGormInsert()
}

// TestWhereEqConvIn 测试添加等于条件会不会转换为in条件
func TestWhereEqConvIn(t *testing.T) {
	mysql_test.TestWhereEqConvIn()
}

// TestWhere  测试where各种写法
func TestWhere(t *testing.T) {
	mysql_test.TestWhere()
}

// TestInsertMysql 插入10000 数据
func TestInsertMysql(t *testing.T) {
	mysql_test.TestInsertMysql()
}

// TestInsert 测试插入数据
func TestInsert(t *testing.T) {
	mysql_test.InsertData()
}
