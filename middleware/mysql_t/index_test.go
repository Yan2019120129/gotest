package mysql_t

import (
	"fmt"
	"gorm.io/gorm/utils"
	"testing"
)

// TestWhereId 测试where 条件
func TestWhereId(t *testing.T) {
	WhereId()
}

// TestWhereIds 测试where 自动映射id数组条件
func TestWhereIds(t *testing.T) {
	WhereIds()
}

// TestWhereOtherIds 测试where 其他数组条件
func TestWhereOtherIds(t *testing.T) {
	WhereOtherIds()
}

// TestGormMapping  测试gorm是否能映射string 为[]string 类型
func TestGormMapping(t *testing.T) {
	GormMapping()
}

// TestUpdated 测试gorm update 回不会自动更新时间
func TestTestUpdated(t *testing.T) {
	Updated()
}

//

// TestGormToString 测试gorm 中ToString方法
func TestGormToString(t *testing.T) {
	fmt.Println("data:", utils.ToString([]string{"/assets/currency/ada.png"}))
}

// 验证使用Model和不实用的区别
func TestGormFind(t *testing.T) {
	GormFind()
}

// TestSelectClient 测试子查询
func TestSelectClient(t *testing.T) {
	SelectClient()
}

// TestGormInsert 添加测试数据
func TestGormInsert(t *testing.T) {
	GormInsert()
}

// TestWhereEqConvIn 测试添加等于条件会不会转换为in条件
func TestWhereEqConvIn(t *testing.T) {
	WhereEqConvIn()
}

// TestWhere  测试where各种写法
func TestWhere(t *testing.T) {
	Where()
}

// TestInsertMysql 插入10000 数据
func TestInsertMysql(t *testing.T) {
	InsertMysql()
}

// TestInsert 测试插入数据
func TestInsert(t *testing.T) {
	//InsertData()
	Insert()
}

// TestInsert 测试插入数据
func TestSelect(t *testing.T) {
	Select()
}

// TestGoroutineGorm 测试插入数据
func TestGoroutineGorm(t *testing.T) {
	GoroutineGorm()
}
