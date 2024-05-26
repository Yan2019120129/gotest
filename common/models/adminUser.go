package models

import "github.com/brianvoe/gofakeit/v6"

// AdminUser 管理表
type AdminUser struct {
	Model
	//Users       []*User `gorm:"foreignKey:AdminId"`
	ParentId    int     `gorm:"type:int unsigned not null;comment:上级ID"`
	UserName    string  `gorm:"column:username;type:varchar(60) not null;comment:用户名"`
	NickName    string  `gorm:"column:nickname;type:varchar(60) not null;comment:昵称"`
	Email       string  `gorm:"type:varchar(60) not null;comment:邮箱"`
	Avatar      string  `gorm:"type:varchar(120) not null;comment:头像"`
	Password    string  `gorm:"type:varchar(120) not null;comment:密码"`
	SecurityKey string  `gorm:"type:varchar(120) not null;comment:密钥"`
	Money       float64 `gorm:"type:decimal(12,2) not null;comment:金额"`
	Status      int     `gorm:"type:smallint not null default 10;default:10;comment:状态 10激活 -1冻结 -2删除"`
	Data        string  `gorm:"type:text;comment:数据"`
	Domains     string  `gorm:"type:varchar(1020) not null;comment:绑定域名"`
	ExpiredAt   int     `gorm:"type:int unsigned not null;comment:过期时间"`
}

// GetDefaultAdminUser 获取默认管理实例
func GetDefaultAdminUser() *AdminUser {
	return &AdminUser{
		ParentId:    0,
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
}

// SetUser  设置用户数据
func (a *AdminUser) SetUser(users ...*User) *AdminUser {
	//a.Users = append(a.Users, users...)
	return a
}

// SetParentId  设置上级Id
func (a *AdminUser) SetParentId(parentId int) *AdminUser {
	a.ParentId = parentId
	return a
}

const (
	// LangStatusActive 激活
	AdminUserStatusActive = 10

	// LangStatusDisable 冻结
	AdminUserStatusDisable = -1

	// LangStatusDelete 删除
	AdminUserStatusDelete = -2
)
