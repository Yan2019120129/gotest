package models

// AdminMenu 管理员菜单
type AdminMenu struct {
	Id        int    `gorm:"type:int unsigned primary key auto_increment;comment:主键;"`
	AdminId   int    `gorm:"type:int unsigned not null;comment:管理ID"`
	ParentId  int    `gorm:"type:int unsigned not null;comment:父级ID"`
	Name      string `gorm:"type:varchar(50) not null;comment:名称"`
	Type      int    `gorm:"type:tinyint not null;comment:类型 1后台菜单 11前台菜单 12用户菜单 13快捷菜单"`
	Route     string `gorm:"type:varchar(50) not null;comment:路由"`
	Sort      int    `gorm:"type:tinyint not null;comment:排序"`
	Status    int    `gorm:"type:tinyint not null default 10;default:10;comment:状态 -2删除 -1禁用 10开启 "`
	Data      string `gorm:"type:text;comment:数据"`
	CreatedAt int    `gorm:"type:int unsigned not null;autoCreateTime;comment:创建时间"`
}

const (
	// AdminMenuTypeAdminMenu 后台菜单
	AdminMenuTypeAdminMenu = 1

	// AdminMenuTypeHomeMenu 前台菜单
	AdminMenuTypeHomeMenu = 11

	// AdminMenuTypeUserMenu 用户菜单
	AdminMenuTypeUserMenu = 12

	// AdminMenuTypeQuickMenu 快捷菜单
	AdminMenuTypeQuickMenu = 13

	// AdminMenuTypeTabbarMenu H5 Tabbar菜单
	AdminMenuTypeTabbarMenu = 15

	// AdminMenuStatusActive 开启
	AdminMenuStatusActive = 10

	// AdminMenuStatusDisable 禁用
	AdminMenuStatusDisable = -1

	// AdminMenuStatusDelete 删除
	AdminMenuStatusDelete = -2
)
