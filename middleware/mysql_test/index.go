package mysql_test

import (
	"fmt"
	"github.com/brianvoe/gofakeit/v6"
	"go.uber.org/zap"
	"gotest/common/models"
	"gotest/common/module/gorm/database"
	"gotest/common/module/logger"
	"gotest/common/utils"
	"gotest/middleware/mysql_test/dto"
	"time"
)

// TestWhereId 测试where 条件
func TestWhereId() {
	id := 24587
	userInfo := &models.User{}
	if result := database.DB.Where(id).Find(userInfo); result.Error != nil {
		logger.Logger.Warn("错误信息", zap.Error(result.Error))
		return
	}
	logger.Logger.Info("信息", zap.String("data", utils.ObjToString(userInfo)))
}

// TestWhereIds 测试where 条件
func TestWhereIds() {
	ids := []int{24587, 24588}
	userInfo := &models.User{}
	if result := database.DB.Where(ids).Find(userInfo); result.Error != nil {
		logger.Logger.Warn("错误信息", zap.Error(result.Error))
		return
	}
	logger.Logger.Info("信息", zap.String("data", utils.ObjToString(userInfo)))
}

// TestWhereOtherIds 测试where 条件，测试结果不可行，会自动映射为主键ID
func TestWhereOtherIds() {
	adminIds := []int{1, 2}
	userInfo := &models.User{}
	if result := database.DB.Where(adminIds).Find(userInfo); result.Error != nil {
		logger.Logger.Warn("错误信息", zap.Error(result.Error))
		return
	}
	logger.Logger.Info("信息", zap.String("data", utils.ObjToString(userInfo)))
}

// TestGormMapping 测试Gorm能否自动映射string 类型为[]string
func TestGormMapping() {
	password := gofakeit.Password(true, true, true, false, false, 10)
	logger.Logger.Info("数据", zap.Reflect("密码", password))
	avatars := []string{gofakeit.ImageURL(200, 100), gofakeit.ImageURL(200, 100)}
	userInfo := &models.User{
		AdminId:     1,
		ParentId:    1,
		UserName:    gofakeit.Name(),
		NickName:    gofakeit.Name(),
		Email:       gofakeit.Email(),
		Telephone:   gofakeit.Phone(),
		Avatar:      utils.ObjToString(avatars),
		Sex:         1,
		Birthday:    int(time.Now().Unix()),
		Password:    password,
		SecurityKey: password,
		Money:       gofakeit.Float64Range(100, 350),
		Type:        1,
		Status:      10,
		Data:        gofakeit.Letter(),
		Desc:        gofakeit.Letter(),
	}
	if result := database.DB.Create(userInfo); result.Error != nil {
		logger.Logger.Warn("错误信息", zap.Error(result.Error))
	}
	logger.Logger.Info("数据", zap.Reflect("创建成功", userInfo))

	userData := &dto.UserData{}
	if result := database.DB.Model(&models.User{}).Where(userInfo.Id).Take(userData); result.Error != nil {
		logger.Logger.Warn("错误信息", zap.Error(result.Error))
	}
	logger.Logger.Info("数据", zap.Reflect("用户信息", userData))
}

// TestUpdated 测试不穿入字段会不会更新时间
func TestUpdated() {
	userInfo := &models.User{}
	if result := database.DB.Where(24591).Find(userInfo); result.Error != nil {
		logger.Logger.Warn("错误信息", zap.Error(result.Error))
	}
	if result := database.DB.Where(userInfo.Id).Updates(userInfo); result.Error != nil {
		logger.Logger.Warn("错误信息", zap.Error(result.Error))
	}
}

// TestGormFind 测试携带Model和不携带Model 的区别
func TestGormFind() {
	userList := []*models.User{}
	if result := database.DB.Find(&userList); result.Error != nil {
		logger.Logger.Warn("错误信息", zap.Error(result.Error))
	}
	logger.Logger.Info("信息", zap.Reflect("用户数据", userList))

	userList1 := []*models.User{}
	if result := database.DB.Model(userList1).Find(&userList1); result.Error != nil {
		logger.Logger.Warn("错误信息", zap.Error(result.Error))
	}
	logger.Logger.Info("信息", zap.Reflect("用户数据", userList))
}

// TestSelectClient 测试子查询,使用sql条件判断
func TestSelectClient() {
	model := database.DB.Table("product AS p").
		Select("p.id", "p.category_id",
			"CASE WHEN (SELECT pc.product_id FROM product_collect AS pc WHERE p.id=pc.product_id) IS NOT NULL THEN true ELSE false END AS IsCollect",
			"p.images", "p.name", "p.money", "p.type", "p.sales", "p.nums", "p.used", "p.total", "p.data")

	productList := []*dto.IndexData{}
	if result := model.Where("admin_id = ?", 1).
		Scan(&productList); result.Error != nil {
	}
	for _, v := range productList {
		fmt.Println("data:", v)
	}
}
