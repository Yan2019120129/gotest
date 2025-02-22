package model

func init() {
	TableManage.addTable(&User{}, "用户")
}

// User 用户
type User struct {
	Base
	Name  string `json:"name" gorm:"type:varchar(64);comment:用户名"`
	Pwd   string `json:"pwd" gorm:"type:varchar(64);comment:密码"`
	Sex   int8   `json:"sex"`
	EMail string `json:"eMail" gorm:"type:varchar(64);comment:电子邮箱"`
}
