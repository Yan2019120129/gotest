package model

func init() {
	TableManage.addTable(&AdminUser{}, "管理员用户")
}

// AdminUser 管理员用户
type AdminUser struct {
	Base
	Name  string `json:"name" gorm:"type:varchar(64);comment:管理员用户名"`
	Pwd   string `json:"pwd" gorm:"comment:密码"`
	EMail string `json:"eMail" gorm:"type:varchar(64);comment:电子邮箱"`
	Phone string `json:"phone" gorm:"type:varchar(64);comment:手机号码"`
}
