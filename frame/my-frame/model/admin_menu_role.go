package model

func init() {
	TableManage.addTable(&AdminMenuRole{}, "管理员菜单关联")
}

// AdminMenuRole 管理员菜单关联
type AdminMenuRole struct {
	Base
	RoleID      uint `json:"roleID" gorm:"comment:权限ID"`
	AdminManuID uint `json:"adminManuID" gorm:"comment:管理员菜单ID"`
}
