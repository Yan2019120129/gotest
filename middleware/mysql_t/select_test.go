package mysql_t

import (
	"gotest/common/models"
	"gotest/common/module/gorm/database"
	"testing"
)

// SelectAllField 全字段
func BenchmarkSelectAllField(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		userInfo := models.User{}
		database.DB.Where("id = ?", 1).Find(&userInfo)
	}
}

// SelectAllField 全字段
func BenchmarkSelectField(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		username := ""
		database.DB.Model(&models.User{}).Select("user_name").Where("id = ?", 1).Find(&username)
	}
}

// SelectAllField 全字段
func BenchmarkSelectField1(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		username := ""
		database.DB.Model(&models.User{}).Select("user_name").Where("id = ?", 1).Scan(&username)
	}
}

// SelectAllField 全字段
func BenchmarkSelectField2(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		username := ""
		database.DB.Raw("select user_name from user where id = ?", 1).Scan(&username)
	}
}
