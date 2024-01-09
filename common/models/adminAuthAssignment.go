package models

// AuthAssignment 权限分配表
type AuthAssignment struct {
	Name      string `gorm:"type:varchar(50) not null;index:idx_admin_name;comment:名称"`
	AdminId   int    `gorm:"type:int unsigned not null;index:idx_admin_name;comment:管理ID"`
	CreatedAt int    `gorm:"type:int unsigned not null;autoCreateTime;comment:创建时间"`
}
