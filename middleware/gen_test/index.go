package gen_test

import (
	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gorm"
)

// Gen 测试gen模型生成
func Gen() {
	g := gen.NewGenerator(gen.Config{
		//  设置输出路径
		OutPath: "/Users/taozi/Documents/Golang/gotest/gen/query",
		Mode:    gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface, // 选择生成模式
	})
	//  建立数据库连接
	db, _ := gorm.Open(mysql.Open("root:Aa123098..@tcp(127.0.0.1:3306)/exchange?charset=utf8mb4&parseTime=True&loc=Local"))
	g.UseDB(db) // 选择数据库连接

	// 为结构模型生成基本类型安全的 DAO API。用户的以下约定

	g.ApplyBasic(
		// Generate struct `User` based on table `users`。
		g.GenerateModel("product"),

		// Generate struct `Employee` based on table `users`。
		g.GenerateModelAs("product_category", "Employee"),

		// Generate struct `User` based on table `users` and generating options。
		g.GenerateModel("user", gen.FieldType("id", "int")),
	)
	g.ApplyBasic(
		// 从当前数据库的所有表生成结构
		g.GenerateAllTable()...,
	)
	// 生成代码
	g.Execute()
}
