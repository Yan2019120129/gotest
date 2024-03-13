package mysql_t

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
	"gotest/middleware/mysql_t/dto"
	"reflect"
	"time"
)

// WhereId 测试where 条件
func WhereId() {
	id := 24587
	userInfo := &models.User{}
	if result := database.DB.Where(id).Find(userInfo); result.Error != nil {
		logs.Logger.Warn("错误信息", zap.Error(result.Error))
		return
	}
	logs.Logger.Info("信息", zap.String("data", utils.ObjToString(userInfo)))
}

// WhereIds 测试where 条件
func WhereIds() {
	ids := []int{24587, 24588}
	userInfo := &models.User{}
	if result := database.DB.Where(ids).Find(userInfo); result.Error != nil {
		logs.Logger.Warn("错误信息", zap.Error(result.Error))
		return
	}
	logs.Logger.Info("信息", zap.String("data", utils.ObjToString(userInfo)))
}

// WhereOtherIds 测试where 条件，测试结果不可行，会自动映射为主键ID
func WhereOtherIds() {
	adminIds := []int{1, 2}
	userInfo := &models.User{}
	if result := database.DB.Where(adminIds).Find(userInfo); result.Error != nil {
		logs.Logger.Warn("错误信息", zap.Error(result.Error))
		return
	}
	logs.Logger.Info("信息", zap.String("data", utils.ObjToString(userInfo)))
}

// GormMapping 测试Gorm能否自动映射string 类型为[]string
func GormMapping() {
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

// Updated 测试不穿入字段会不会更新时间
func Updated() {
	userInfo := &models.User{}
	if result := database.DB.Where(24591).Find(userInfo); result.Error != nil {
		logs.Logger.Warn("错误信息", zap.Error(result.Error))
	}
	if result := database.DB.Where(userInfo.Id).Updates(userInfo); result.Error != nil {
		logs.Logger.Warn("错误信息", zap.Error(result.Error))
	}
}

// GormFind 测试携带Model和不携带Model 的区别
func GormFind() {
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

// SelectClient 测试子查询,使用sql条件判断
func SelectClient() {
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

// GormInsert 插入数据
func GormInsert() {
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

// WhereEqConvIn 测试等于是否会转换为in
func WhereEqConvIn() {
	adminIds := []int{1, 2}
	userList := &models.User{}
	if result := database.DB.Where(adminIds).Find(userList); result.Error != nil {
		logs.Logger.Error("gorm", zap.Error(result.Error))
	}
	logs.Logger.Info("gorm", zap.Reflect("data", userList))
}

// Where 测试where 各种写法
func Where() {
	var userIds []int
	database.DB.Model(&models.User{}).Pluck("id", &userIds)
	logs.Logger.Info("mysql", zap.Reflect("userIds", userIds))
	logs.Logger.Info("mysql", zap.Reflect("Len(userIds)", len(userIds)))

	userInfo := make([]*models.User, 0)
	database.DB.Where("username = ? OR telephone = ?", "ceshi1", "15577098754").Find(&userInfo)
	logs.Logger.Info("mysql", zap.Reflect("userInfo", userInfo))

	userInfoTwo := &models.User{}
	database.DB.Where("username = ? AND telephone = ?", "ceshi2", "15577098754").Take(userInfoTwo)
	logs.Logger.Info("mysql", zap.Reflect("userInfoTwo", userInfoTwo))

	var userId int
	result := database.DB.Model(&models.User{}).Select("id").Where(1).Scan(&userId)
	logs.Logger.Info("mysql", zap.Reflect("Scan.Error", result.Error))
	logs.Logger.Info("mysql", zap.Reflect("Scan.RowsAffected", result.RowsAffected))
	logs.Logger.Info("mysql", zap.Reflect("userId", userId))

	userInfoThree := &models.User{}
	database.DB.Where(models.User{UserName: "ceshi2", Telephone: ""}).Take(userInfoThree)
	logs.Logger.Info("mysql", zap.Reflect("userInfoThree", userInfoThree))

	userInfoFourth := &models.User{}
	type Test struct {
		Model *gorm.DB
	}

	test := Test{database.DB.Where(24588)}
	database.DB.Where(test.Model).Where("admin_id = ?", 2).Take(userInfoTwo)
	logs.Logger.Info("mysql", zap.Reflect("userInfoFourth", userInfoFourth))

	var userReflect interface{}
	userReflect = &models.User{}
	structType := reflect.TypeOf(userReflect)
	structInstance := reflect.New(structType)
	userInstance := structInstance.Type()
	database.DB.Find(&userInstance)
	logs.Logger.Info("mysql", zap.Reflect("structType", structType))

	logs.Logger.Fatal("fatal", zap.String(logs.LogMsgApp, "test"))
}

// InsertMysql 100w 插入一百万的数据
func InsertMysql() {
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

// InsertData 插入数据
func InsertData() {
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
	database.DB.Create(userInfo)

	userMap := map[string]interface{}{}
	database.DB.Model(&models.User{}).Where(24590).Take(&userMap)
	logs.Logger.Info("mysql", zap.Reflect("userMap", userMap))

	delete(userMap, "")
	userMap["id"] = 0
	userMap["username"] = gofakeit.Name()
	userMap["email"] = gofakeit.Email()
	userMap["telephone"] = gofakeit.Phone()
	database.DB.Model(&models.User{}).Create(&userMap)
	logs.Logger.Info("mysql", zap.Reflect("userMap", userMap))
	logs.Logger.Info("mysql", zap.Reflect("userId", userMap["id"]))
}

// Select 测试select
func Select() {
	type Test struct {
		Id  int    `json:"id"`
		Key string `json:"key"`
	}
	test := &Test{}
	database.DB.Model(&models.User{}).Select("id", "username as 'Key'").Where(24587).Find(test)
	logs.Logger.Info("mysql", zap.Reflect("user", test))
}

// GoroutineGorm 测试多协程情况下连接问题
func GoroutineGorm() {
	for i := 0; i < 50; i++ {
		go CrontabSelect(1 * time.Second)
	}
	time.Sleep(100 * time.Second)
}

// CrontabSelect 测试select
func CrontabSelect(second time.Duration) {
	ch := time.NewTicker(second)
	for {
		type Test struct {
			Id  int    `json:"id"`
			Key string `json:"key"`
		}
		test := make([]*Test, 0)
		database.DB.Model(&models.User{}).Select("id", "username as 'Key'").Find(&test)
		logs.Logger.Info("mysql", zap.Reflect("user", test))
		<-ch.C
	}

}

// Insert 测试插入
func Insert() {
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
	//userMap := make(map[string]interface{})
	//userMap["admin_id"] = gofakeit.Number(1, 100)
	//userMap["parent_id"] = gofakeit.Number(1, 100)
	//userMap["username"] = gofakeit.Name()
	//userMap["nickname"] = gofakeit.Name()
	//userMap["email"] = gofakeit.Email()
	//userMap["telephone"] = gofakeit.Phone()
	//userMap["avatar"] = gofakeit.ImageURL(200, 100)
	//userMap["sex"] = gofakeit.RandomInt([]int{1, 2})
	//userMap["password"] = gofakeit.Password(true, true, true, false, false, 10)
	//userMap["security_key"] = gofakeit.Password(true, true, true, false, false, 10)
	//userMap["money"] = gofakeit.Float64Range(100, 100000)
	//userMap["birthday"] = gofakeit.Number(100000000, 1000000000)
	//userMap["type"] = gofakeit.RandomInt([]int{1, 11, 21})
	//userMap["status"] = gofakeit.RandomInt([]int{-2, -1, 10})
	//userMap["data"] = gofakeit.Sentence(10)
	//database.DB.Table("user").Create(userMap)
	database.DB.Table("user").Create(userInfo)
}
