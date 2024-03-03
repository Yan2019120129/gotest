package models

// AuthItem 权限目录
type AuthItem struct {
	Name      string `gorm:"type:varchar(50) primary key not null;comment:名称"`
	Type      int    `gorm:"type:tinyint not null;comment:类型 1权限角色 2路由权限 3路由名称"`
	Desc      string `gorm:"type:varchar(255) not null;comment:详情"`
	Rule      string `gorm:"type:varchar(255) not null;comment:规则"`
	Data      string `gorm:"type:varchar(255) not null;comment:数据"`
	UpdatedAt int    `gorm:"type:int unsigned not null;autoUpdateTime;comment:更新时间"`
	CreatedAt int    `gorm:"type:int unsigned not null;autoCreateTime;comment:创建时间"`
}

const (
	//	AuthItemTypeRole 权限角色
	AuthItemTypeRole = 1

	//	AuthItemTypeRoute 权限路由
	AuthItemTypeRoute = 2

	//	AuthItemTypeName 路由名称类型
	AuthItemTypeName = 3

	// AuthRoleSuperManage 超级管理员
	AuthRoleSuperManage = "超级管理员"

	// AuthRoleMerchatManage 商户管理员
	AuthRoleMerchatManage = "商户管理员"

	// AuthRoleAgentManage 代理管理员
	AuthRoleAgentManage = "代理管理员"
)
