package models

const (
	AccessLogsTypeAdmin = 1 //	后端日志
	AccessLogsTypeHome  = 2 //	前端日志
)

type AccessLogs struct {
	Id        int    `gorm:"type:int unsigned primary key auto_increment;comment:主键"` // 主键
	AdminId   int    `gorm:"type:int unsigned not null;comment:管理员ID"`                // 管理员ID
	UserId    int    `gorm:"type:int unsigned not null;comment:用户ID"`                 // 用户ID
	Type      int    `gorm:"type:int unsigned not null;comment:日志类型"`                 // 日志类型
	Name      string `gorm:"type:varchar(50) not null;comment:标题"`                    // 标题
	Ip4       string `gorm:"type:varchar(120);not null;comment:IP4地址"`                // IP4地址
	UserAgent string `gorm:"type:varchar(255) not null;comment:请求头ua信息"`              // ua信息
	Lang      string `gorm:"type:varchar(255) not null;comment:语言信息"`                 // 语言信息
	Route     string `gorm:"type:varchar(255) not null;comment:操作路由"`                 // 操作路由
	Data      string `gorm:"type:varchar(255) not null;comment:数据"`                   // 数据
	CreatedAt int    `gorm:"type:int unsigned not null;comment:时间"`                   // 时间
}
