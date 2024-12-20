package model

func init() {
	TableManage.addTable(&AdminRole{}, "管理员权限关联")
}

// AdminRole 管理员权限关联
type AdminRole struct {
	Base
	RoleID  uint `json:"roleID" gorm:"comment:权限ID"`
	AdminID uint `json:"adminID" gorm:"comment:管理员ID"`
}
