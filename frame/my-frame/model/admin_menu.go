package model

func init() {
	TableManage.addTable(&AdminMenu{}, "管理员菜单")
}

// AdminMenu 管理员菜单
type AdminMenu struct {
	Base
	Route    string `json:"route" gorm:"comment:路由"`
	ParentID string `json:"parentID" gorm:"comment:上级路由ID"`
	Desc     string `json:"desc" gorm:"comment:详情"`
}
