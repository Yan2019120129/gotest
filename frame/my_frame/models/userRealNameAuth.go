package models

// RealNameAuth 实名认证
type RealNameAuth struct {
	Id        int    `gorm:"type:int unsigned primary key auto_increment;comment:主键;"`
	AdminId   int    `gorm:"type:int unsigned not null;comment:管理ID"`
	UserId    int    `gorm:"type:int unsigned not null;comment:用户ID"`
	RealName  string `gorm:"type:varchar(50) not null;comment:真实姓名"`
	Number    string `gorm:"type:varchar(50) not null;comment:卡号"`
	Photo1    string `gorm:"type:varchar(120) not null;comment:证件照1"`
	Photo2    string `gorm:"type:varchar(120) not null;comment:证件照2"`
	Photo3    string `gorm:"type:varchar(120) not null;comment:证件照3"`
	Telephone string `gorm:"type:varchar(50) not null;comment:手机号码"`
	Address   string `gorm:"type:varchar(255) not null;default:'';comment:详细地址"`
	Type      int    `gorm:"type:tinyint not null default 1;default:1;comment:类型 1身份证"`
	Status    int    `gorm:"type:tinyint not null default 10;default:10;comment:状态 -2删除 -1拒绝 10审核 20完成"`
	Data      string `gorm:"type:text;comment:input配置"`
	UpdatedAt int    `gorm:"type:int unsigned not null;autoUpdateTime;comment:更新时间"`
	CreatedAt int    `gorm:"type:int unsigned not null;autoCreateTime;comment:创建时间"`
}

const (
	// UserRealNameTypeIdCard 身份证
	UserRealNameTypeIdCard = 1

	// UserRealNameStatusComplete 完成
	UserRealNameStatusComplete = 20

	// UserRealNameStatusActive 审核
	UserRealNameStatusActive = 10

	// UserRealNameStatusRefuse 拒绝
	UserRealNameStatusRefuse = -1

	// UserRealNameStatusDelete 删除
	UserRealNameStatusDelete = -2
)
