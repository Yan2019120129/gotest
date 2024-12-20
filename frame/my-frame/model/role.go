package model

func init() {
	TableManage.addTable(&Role{}, "权限")
}

// Role 权限
type Role struct {
	Base
	Name string `json:"name" gorm:"type:varchar(64);comment:名称"`
	Desc string `json:"desc" gorm:"type:varchar(128);comment:描述"`
}
