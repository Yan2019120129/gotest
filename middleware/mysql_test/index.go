package mysql_test

import (
	"encoding/json"
	"fmt"
	"github.com/brianvoe/gofakeit/v6"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"gotest/common/models"
	"gotest/common/module/gorm/database"
	"gotest/common/module/logs"
	"gotest/common/utils"
	"gotest/middleware/mysql_test/dto"
	"time"
)

// TestWhereId 测试where 条件
func TestWhereId() {
	id := 24587
	userInfo := &models.User{}
	if result := database.DB.Where(id).Find(userInfo); result.Error != nil {
		logs.Logger.Warn("错误信息", zap.Error(result.Error))
		return
	}
	logs.Logger.Info("信息", zap.String("data", utils.ObjToString(userInfo)))
}

// TestWhereIds 测试where 条件
func TestWhereIds() {
	ids := []int{24587, 24588}
	userInfo := &models.User{}
	if result := database.DB.Where(ids).Find(userInfo); result.Error != nil {
		logs.Logger.Warn("错误信息", zap.Error(result.Error))
		return
	}
	logs.Logger.Info("信息", zap.String("data", utils.ObjToString(userInfo)))
}

// TestWhereOtherIds 测试where 条件，测试结果不可行，会自动映射为主键ID
func TestWhereOtherIds() {
	adminIds := []int{1, 2}
	userInfo := &models.User{}
	if result := database.DB.Where(adminIds).Find(userInfo); result.Error != nil {
		logs.Logger.Warn("错误信息", zap.Error(result.Error))
		return
	}
	logs.Logger.Info("信息", zap.String("data", utils.ObjToString(userInfo)))
}

// TestGormMapping 测试Gorm能否自动映射string 类型为[]string
func TestGormMapping() {
	password := gofakeit.Password(true, true, true, false, false, 10)
	logs.Logger.Info("数据", zap.Reflect("密码", password))
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
		logs.Logger.Warn("错误信息", zap.Error(result.Error))
	}
	logs.Logger.Info("数据", zap.Reflect("创建成功", userInfo))

	userData := &dto.UserData{}
	if result := database.DB.Model(&models.User{}).Where(userInfo.Id).Take(userData); result.Error != nil {
		logs.Logger.Warn("错误信息", zap.Error(result.Error))
	}
	logs.Logger.Info("数据", zap.Reflect("用户信息", userData))
}

// TestUpdated 测试不穿入字段会不会更新时间
func TestUpdated() {
	userInfo := &models.User{}
	if result := database.DB.Where(24591).Find(userInfo); result.Error != nil {
		logs.Logger.Warn("错误信息", zap.Error(result.Error))
	}
	if result := database.DB.Where(userInfo.Id).Updates(userInfo); result.Error != nil {
		logs.Logger.Warn("错误信息", zap.Error(result.Error))
	}
}

// TestGormFind 测试携带Model和不携带Model 的区别
func TestGormFind() {
	userList := []*models.User{}
	if result := database.DB.Where(0).Take(&userList); result.Error != nil {
		logs.Logger.Error("Gorm", zap.Error(result.Error))
	}
	logs.Logger.Info("Gorm", zap.Reflect("data", userList))

	userList1 := []*models.User{}
	if result := database.DB.Model(userList1).Find(&userList1); result.Error != nil {
		logs.Logger.Error("Gorm", zap.Error(result.Error))
	}
	logs.Logger.Info("Gorm", zap.Reflect("data", userList))
}

// TestSelectClient 测试子查询,使用sql条件判断
func TestSelectClient() {
	//model := database.DB.Table("product AS p").
	//	Select("p.id", "p.category_id",
	//		"CASE WHEN (SELECT pc.product_id FROM product_collect AS pc WHERE p.id=pc.product_id) IS NOT NULL THEN true ELSE false END AS isCollect",
	//		"p.images", "p.name", "p.money", "p.type", "p.sales", "p.nums", "p.used", "p.total", "p.data")

	productList := make([]map[string]interface{}, 0)
	//if result := model.Where("admin_id = ?", 1).
	//	Scan(&productList); result.Error != nil {
	//}

	database.DB.Model(&models.Product{}).Find(&productList)

	//data := []*dto.IndexData{}
	//fmt.Println("data:", productList)
	for _, v := range productList {
		//productInfo := &dto.IndexData{}
		productInfo := &models.Product{}
		dataByteArray, err := json.Marshal(v)
		if err != nil {
			panic(err)
		}
		if err = json.Unmarshal(dataByteArray, productInfo); err != nil {
			panic(err)
		}
		//if result := database.DB.Where("product_id = ?", productInfo.Id).
		//	Where("user_id = ?", 1).
		//	Where("admin_id = ?", 1).
		//	Find(&models.ProductCollect{}); result.Error != nil {
		//} else if result.RowsAffected > 0 {
		//	productInfo.IsCollect = true
		//}
		fmt.Println("data:", productInfo)
		//data = append(data, productInfo)
	}
}

// TestGormInsert 插入数据
func TestGormInsert() {
	productList := []*models.Product{}
	if result := database.DB.Find(&productList); result.Error != nil {
		logs.Logger.Warn("错误信息", zap.Error(result.Error))
	}
	for _, v := range productList {
		data := &dto.ProductData{
			InstId:    gofakeit.MonthString(),
			Last:      gofakeit.Float64Range(1000, 10000),
			LastSz:    gofakeit.Float64Range(10, 100),
			Open24h:   gofakeit.Float64Range(500, 1000),
			High24h:   gofakeit.Float64Range(500, 1000),
			Low24h:    gofakeit.Float64Range(500, 1000),
			Vol24h:    gofakeit.Float64Range(500, 1000),
			Amount24h: gofakeit.Float64Range(500, 1000),
			Ts:        time.Now().Unix(),
		}
		dataBytes, err := json.Marshal(data)
		if err != nil {
			logs.Logger.Warn("错误信息", zap.Error(err))
		}
		v.Data = string(dataBytes)
		if result := database.DB.Updates(v); result.Error != nil {
			logs.Logger.Warn("错误信息", zap.Error(result.Error))
		}
	}
}

// TestWhereEqConvIn 测试等于是否会转换为in
func TestWhereEqConvIn() {
	adminIds := []int{1, 2}
	userList := &models.User{}
	if result := database.DB.Where(adminIds).Find(userList); result.Error != nil {
		logs.Logger.Error("gorm", zap.Error(result.Error))
	}
	logs.Logger.Info("gorm", zap.Reflect("data", userList))
}

// TestInsertMysql 100w 插入一百万的数据
func TestInsertMysql() {
	for i := 0; i < 100000; i++ {
		if err := database.DB.Transaction(func(tx *gorm.DB) error {
			// 插入管理员数据
			adminInfo := &models.AdminUser{
				ParentId:    gofakeit.Number(1, 100),
				UserName:    gofakeit.Name(),
				NickName:    gofakeit.Name(),
				Email:       gofakeit.Email(),
				Avatar:      gofakeit.ImageURL(200, 100),
				Password:    gofakeit.Password(true, true, true, false, false, 10),
				SecurityKey: gofakeit.Password(true, true, true, false, false, 10),
				Money:       gofakeit.Float64Range(100, 100000),
				Status:      gofakeit.RandomInt([]int{-2, -1, 10}),
				Data:        gofakeit.Letter(),
				Domains:     gofakeit.Letter(),
			}
			if result := tx.Create(adminInfo); result.Error != nil {
				logs.Logger.Error("mysql", zap.Error(result.Error))
				return result.Error
			}

			// 插入用户数据
			userInfo := &models.User{
				AdminId:     gofakeit.Number(1, 100),
				ParentId:    gofakeit.Number(1, 100),
				UserName:    gofakeit.Name(),
				NickName:    gofakeit.Name(),
				Email:       gofakeit.Email(),
				Telephone:   gofakeit.Phone(),
				Avatar:      gofakeit.ImageURL(200, 100),
				Sex:         gofakeit.RandomInt([]int{1, 2}),
				Password:    gofakeit.Password(true, true, true, false, false, 10),
				SecurityKey: gofakeit.Password(true, true, true, false, false, 10),
				Money:       gofakeit.Float64Range(100, 100000),
				Type:        gofakeit.RandomInt([]int{1, 11, 21}),
				Status:      gofakeit.RandomInt([]int{-2, -1, 10}),
				Data:        gofakeit.Sentence(10),
				Desc:        gofakeit.Sentence(20),
			}

			if result := tx.Create(userInfo); result.Error != nil {
			}

			// 插入用户资产
			userAssetsInfo := &models.WalletUserAssets{
				AdminId:        gofakeit.Number(1, 100),
				UserId:         gofakeit.Number(1, 100),
				WalletAssetsId: gofakeit.Number(1, 100),
				Money:          gofakeit.Float64Range(100, 100000),
				Status:         gofakeit.RandomInt([]int{-2, -1, 10}),
				Data:           gofakeit.Sentence(10),
			}
			if result := tx.Create(userAssetsInfo); result.Error != nil {
				return result.Error
			}

			// 插入钱包订单
			walletOrderInfo := &models.WalletOrder{
				AdminId:  gofakeit.Number(1, 100),
				UserId:   gofakeit.Number(1, 100),
				AssetsId: gofakeit.Number(1, 100),
				SourceId: gofakeit.RandomInt([]int{1, 2, 11, 12, 20}),
				Type:     gofakeit.RandomInt([]int{1, 2, 11, 12}),
				OrderSn:  utils.NewRandom().OrderSn(),
				Money:    gofakeit.Float64Range(100, 100000),
				Voucher:  gofakeit.ImageURL(200, 100),
				Fee:      gofakeit.Float64Range(0.1, 1),
				Status:   gofakeit.RandomInt([]int{-2, -1, 10}),
				Data:     gofakeit.Sentence(10),
			}
			if result := tx.Create(walletOrderInfo); result.Error != nil {
				return result.Error
			}
			logs.Logger.Info("mysql", zap.String("result", "Finish"))
			return nil
		}); err != nil {
			logs.Logger.Error("mysql", zap.String("result", "fail"), zap.Error(err))
		}

	}
}
