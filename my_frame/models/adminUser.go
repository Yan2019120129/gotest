package models

// AdminUser 管理表
type AdminUser struct {
	Id          int     `gorm:"type:int unsigned primary key auto_increment;comment:主键;"`
	ParentId    int     `gorm:"type:int unsigned not null;comment:上级ID"`
	UserName    string  `gorm:"column:username;type:varchar(60) not null;comment:用户名"`
	NickName    string  `gorm:"column:nickname;type:varchar(60) not null;comment:昵称"`
	Email       string  `gorm:"type:varchar(60) not null;comment:邮箱"`
	Avatar      string  `gorm:"type:varchar(120) not null;comment:头像"`
	Password    string  `gorm:"type:varchar(120) not null;comment:密码"`
	SecurityKey string  `gorm:"type:varchar(120) not null;comment:密钥"`
	Money       float64 `gorm:"type:decimal(12,2) not null;comment:金额"`
	Status      int     `gorm:"type:smallint not null default 10;default:10;comment:状态 10激活 -1冻结 -2删除"`
	Data        string  `gorm:"type:text;comment:数据"`
	Domains     string  `gorm:"type:varchar(1020) not null;comment:绑定域名"`
	ExpiredAt   int     `gorm:"type:int unsigned not null;comment:过期时间"`
	UpdatedAt   int     `gorm:"type:int unsigned not null;autoUpdateTime;comment:更新时间"`
	CreatedAt   int     `gorm:"type:int unsigned not null;autoCreateTime;comment:创建时间"`
}

const (
	// LangStatusActive 激活
	AdminUserStatusActive = 10

	// LangStatusDisable 冻结
	AdminUserStatusDisable = -1

	// LangStatusDelete 删除
	AdminUserStatusDelete = -2
)
