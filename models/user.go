package models

// User 用户表
type User struct {
	Id          int     `gorm:"type:int unsigned primary key auto_increment;comment:主键;"`
	AdminId     int     `gorm:"type:int unsigned not null;comment:管理ID"`
	ParentId    int     `gorm:"type:int unsigned not null;comment:父级ID"`
	UserName    string  `gorm:"type:varchar(60) not null;comment:用户名"`
	NickName    string  `gorm:"type:varchar(60) not null;comment:昵称"`
	Email       string  `gorm:"type:varchar(60) not null;comment:邮箱"`
	Telephone   string  `gorm:"type:varchar(50) not null;comment:手机号码"`
	Avatar      string  `gorm:"type:varchar(120) not null;comment:头像"`
	Sex         int     `gorm:"type:tinyint not null;comment:性别"`
	Birthday    int     `gorm:"type:int unsigned not null;comment:生日"`
	Password    string  `gorm:"type:varchar(120) not null;comment:密码"`
	SecurityKey string  `gorm:"type:varchar(120) not null;comment:密钥"`
	Money       float64 `gorm:"type:decimal(12,2) not null;comment:金额"`
	Type        int     `gorm:"type:tinyint not null default 11;comment:类型"`
	Status      int     `gorm:"type:smallint not null default 10;comment:状态"`
	Data        string  `gorm:"type:text;comment:数据"`
	Desc        string  `gorm:"type:text;comment:详情"`
	UpdatedAt   int     `gorm:"type:int unsigned not null;comment:更新时间"`
	CreatedAt   int     `gorm:"type:int unsigned not null;comment:创建时间"`
}

const (
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

// UserInfo 用户信息
type UserInfo struct {
	Id       int    `json:"id"`       //	用户ID
	ParentId int    `json:"parentId"` //	上级ID
	UserName string `json:"username"` //	用户名
	NickName string `json:"nickname"` //	用户昵称
}
