package models

// Invite 用户邀请码
type Invite struct {
	Id        int    `gorm:"type:int unsigned primary key auto_increment;comment:主键;"`
	AdminId   int    `gorm:"type:int unsigned not null;comment:管理ID"`
	UserId    int    `gorm:"type:int unsigned not null;comment:用户ID"`
	Code      string `gorm:"type:varchar(50) not null;uniqueIndex;comment:邀请码"`
	Status    int    `gorm:"type:tinyint not null default 10;default:10;comment:状态 -2删除 -1禁用 10开启"`
	Data      string `gorm:"type:text;comment:数据"`
	UpdatedAt int    `gorm:"type:int unsigned not null;autoUpdateTime;comment:更新时间"`
	CreatedAt int    `gorm:"type:int unsigned not null;autoCreateTime;comment:创建时间"`
}

const (
	// UserInviteTypeAdmin 管理邀请码
	UserInviteTypeAdmin = 1

	// UserInviteTypeUser 用户邀请码
	UserInviteTypeUser = 2

	// UserInviteStatusActive 开启
	UserInviteStatusActive = 10

	// UserInviteStatusDisable 禁用
	UserInviteStatusDisable = -1

	// UserInviteStatusDelete 删除
	UserInviteStatusDelete = -2
)
