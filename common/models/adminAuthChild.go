package models

// AuthChild 权限目录子级表
type AuthChild struct {
	Parent string `gorm:"type:varchar(50) not null;index:idx_parent_child;comment:父级"`
	Child  string `gorm:"type:varchar(50) not null;index:idx_parent_child;comment:子级"`
	Type   int    `gorm:"type:tinyint not null;comment:类型 2路由名称对应路由 3角色对应路由名"`
}

const (
	// SuperAdminId 超级管理ID
	SuperAdminId = 1

	// AdminMerchat 商户管理
	AdminMerchat = 2

	// AdminAgent 代理管理
	AdminAgent = 3

	// AuthChildTypeRoleParentRole 角色对应角色
	AuthChildTypeRoleParentRole = 1

	// AuthChildTypeRouteNameRoute 路由名称对应路由
	AuthChildTypeRouteNameRoute = 2

	// AuthChildTypeRoleRouteName 角色对应路由名称
	AuthChildTypeRoleRouteName = 3
)
