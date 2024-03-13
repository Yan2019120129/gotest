package database

// Database 抽象产品接口
type Database interface {
	// NewDatabase 创建数据库
	NewDatabase()

	// Use 使用数据库
	Use()
}
