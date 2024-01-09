package models

// Notify 系统通知表
type Notify struct {
	Id        int    `gorm:"type:int unsigned primary key auto_increment;comment:主键;"`
	AdminId   int    `gorm:"type:int unsigned not null;comment:管理ID"`
	UserId    int    `gorm:"type:int unsigned not null;comment:用户ID"`
	Type      int    `gorm:"type:smallint not null default 11;default:11;comment:类型 11前台通知"`
	Name      string `gorm:"type:varchar(60) not null;comment:标题"`
	Content   string `gorm:"type:text;comment:内容"`
	Status    int    `gorm:"type:smallint not null default 10;default:10;comment:状态"`
	Data      string `gorm:"type:text;comment:数据"`
	UpdatedAt int    `gorm:"type:int unsigned not null;autoUpdateTime;comment:更新时间"`
	CreatedAt int    `gorm:"type:int unsigned not null;autoCreateTime;comment:创建时间"`
}

const (
	// SystemNotifyAdminMessage 管理通知消息
	//SystemNotifyAdminMessage = 1

	// SystemNotifyHomeMessage 用户通知消息
	SystemNotifyHomeMessage = 11

	// SystemNotifyAdminSystemMessage 管理系统通知消息
	SystemNotifyAdminSystemMessage = 21

	// SystemNotifyHomeSystemMessage 用户系统通知消息
	SystemNotifyHomeSystemMessage = 31

	// SystemNotifyStatusDelete 消息删除
	SystemNotifyStatusDelete = -2

	// SystemNotifyStatusUnread 未读消息
	SystemNotifyStatusUnread = 10

	// SystemNotifyStatusRead 已读消息
	SystemNotifyStatusRead = 20
)
