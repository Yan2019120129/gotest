package models

import (
	"github.com/brianvoe/gofakeit/v6"
)

// User 用户表
type User struct {
	Model
	AdminUser   *AdminUser
	AdminUserId int     `gorm:"type:int unsigned not null;default:1;comment:管理ID"`
	ParentId    int     `gorm:"type:int unsigned not null;comment:父级ID"`
	UserName    string  `gorm:"column:username;uniqueIndex;type:varchar(60) not null;comment:用户名"`
	NickName    string  `gorm:"column:nickname;type:varchar(60) not null;comment:昵称"`
	Email       string  `gorm:"uniqueIndex;type:varchar(60);default:null;comment:邮箱"`
	Telephone   string  `gorm:"uniqueIndex;type:varchar(50);default:null;comment:手机号码"`
	Avatar      string  `gorm:"type:varchar(120) not null;comment:头像"`
	Sex         int     `gorm:"type:tinyint not null;comment:性别 1男 2女"`
	Birthday    int     `gorm:"type:int unsigned not null;comment:生日"`
	Password    string  `gorm:"type:varchar(120) not null;comment:密码"`
	SecurityKey string  `gorm:"type:varchar(120) not null;comment:密钥"`
	Money       float64 `gorm:"type:decimal(12,2) not null;comment:金额"`
	Type        int     `gorm:"type:tinyint not null default 11;default:11;comment:类型 1虚拟用户 11默认用户 21会员用户"`
	Status      int     `gorm:"type:smallint not null default 10;default:10;comment:状态 -2删除 -1冻结 10激活"`
	Data        string  `gorm:"type:text;comment:数据"`
	Desc        string  `gorm:"type:text;comment:详情"`
}

func GetDefaultUser() *User {
	return &User{
		AdminUserId: 0,
		ParentId:    0,
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
}
func (u *User) SetAdminId(adminUserId int) *User {
	u.AdminUserId = adminUserId
	return u
}

func (u *User) SetAdminUser(adminUser *AdminUser) *User {
	u.AdminUser = adminUser
	return u
}

func (u *User) SetParentId(parentId int) *User {
	u.ParentId = parentId
	return u
}

const (
	//  UserSexMale 男
	UserSexMale = 1

	// UserSexWoman 女
	UserSexWoman = 2

	// UserStatusActive 激活
	UserStatusActive = 10

	// UserStatusDisable 冻结
	UserStatusDisable = -1

	// UserStatusDelete 删除
	UserStatusDelete = -2

	// UserTypeVirtual 虚拟用户
	UserTypeVirtual = 1

	// UserTypeDefault 默认用户
	UserTypeDefault = 11

	// UserTypeLevel 会员用户
	UserTypeLevel = 21
)
