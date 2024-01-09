package models

// User 用户表
type User struct {
	Id          int     `gorm:"type:int unsigned primary key;autoincrement:1000;comment:主键;"`
	AdminId     int     `gorm:"type:int unsigned not null;default:1;comment:管理ID"`
	ParentId    int     `gorm:"type:int unsigned not null;comment:父级ID"`
	UserName    string  `gorm:"column:username;uniqueIndex;type:varchar(60) not null;comment:用户名"`
	NickName    string  `gorm:"column:nickname;type:varchar(60) not null;comment:昵称"`
	Email       string  `gorm:"uniqueIndex;type:varchar(60);default:null;comment:邮箱"`
	Telephone   string  `gorm:"uniqueIndex;type:varchar(50);default:null;comment:手机号码"`
	Avatar      string  `gorm:"type:varchar(120) not null;comment:头像"`
	Sex         int     `gorm:"type:tinyint not null;comment:性别 1男 2女"`
	Birthday    int     `gorm:"type:int unsigned not null;comment:生日"`
	Password    string  `gorm:"type:varchar(120) not null;comment:密码"`
	SecurityKey string  `gorm:"type:varchar(120) not null;comment:密钥"`
	Money       float64 `gorm:"type:decimal(12,2) not null;comment:金额"`
	Type        int     `gorm:"type:tinyint not null default 11;default:11;comment:类型 1虚拟用户 11默认用户 21会员用户"`
	Status      int     `gorm:"type:smallint not null default 10;default:10;comment:状态 -2删除 -1冻结 10激活"`
	Data        string  `gorm:"type:text;comment:数据"`
	Desc        string  `gorm:"type:text;comment:详情"`
	UpdatedAt   int     `gorm:"type:int unsigned not null;autoUpdateTime;comment:更新时间"`
	CreatedAt   int     `gorm:"type:int unsigned not null;autoCreateTime;comment:创建时间"`
}

const (
	//  UserSexMale 男
	UserSexMale = 1

	// UserSexWoman 女
	UserSexWoman = 2

	// UserStatusActive 激活
	UserStatusActive = 10

	// UserStatusDisable 冻结
	UserStatusDisable = -1

	// UserStatusDelete 删除
	UserStatusDelete = -2

	// UserTypeVirtual 虚拟用户
	UserTypeVirtual = 1

	// UserTypeDefault 默认用户
	UserTypeDefault = 11

	// UserTypeLevel 会员用户
	UserTypeLevel = 21
)
