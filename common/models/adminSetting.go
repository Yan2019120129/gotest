package models

// AdminSetting 管理设置表
type AdminSetting struct {
	Id        int    `gorm:"type:int unsigned primary key auto_increment;comment:主键;"`
	AdminId   int    `gorm:"type:int unsigned not null;comment:管理ID"`
	GroupId   int    `gorm:"type:int unsigned not null;comment:分组ID"`
	Name      string `gorm:"type:varchar(60) not null;comment:设置名称"`
	Type      int    `gorm:"type:tinyint not null default 1;default:1;comment:input类型 1文本类型 2文本域类型 3富文本类型 4数字类型 5密码类型 10选择类型 11单选类型 12多选类型 13开关类型 21文件类型 22图片类型 23图片组类型 31时间类型 32时间范围类型 41对象类型 42子集对象类型 52编辑文本类型 53编辑富文本类型 54编辑开关类型 61翻译类型"`
	Field     string `gorm:"type:varchar(60) not null;comment:建铭"`
	Value     string `gorm:"type:text;comment:键值"`
	Data      string `gorm:"type:text;comment:input配置"`
	UpdatedAt int    `gorm:"type:int unsigned not null;autoUpdateTime;comment:更新时间"`
	CreatedAt int    `gorm:"type:int unsigned not null;autoCreateTime;comment:创建时间"`
}

const (
	// AdminSettingGroupBasic 基础分组
	AdminSettingGroupBasic = 1

	// AdminSettingGroupWallet 钱包分组
	AdminSettingGroupWallet = 2

	// AdminSettingGroupEarnings 收益分组
	AdminSettingGroupEarnings = 3

	// AdminSettingGroupTemplate 模版配置
	AdminSettingGroupTemplate = 4

	// AdminSettingTypeText 文本类型
	AdminSettingTypeText = 1

	// AdminSettingTypeTextArea 文本域类型
	AdminSettingTypeTextArea = 2

	// AdminSettingTypeEditor 富文本类型
	AdminSettingTypeEditor = 3

	// AdminSettingTypeNumber 数字类型
	AdminSettingTypeNumber = 4

	// AdminSettingTypePassword 密码类型
	AdminSettingTypePassword = 5

	// AdminSettingTypeSelect 选择类型
	AdminSettingTypeSelect = 10

	// AdminSettingTypeRadio 单选类型
	AdminSettingTypeRadio = 11

	// AdminSettingTypeCheckbox 多选类型
	AdminSettingTypeCheckbox = 12

	// AdminSettingTypeToggle 开关类型
	AdminSettingTypeToggle = 13

	// AdminSettingTypeFile 文件类型
	AdminSettingTypeFile = 21

	// AdminSettingTypeImage 图片类型
	AdminSettingTypeImage = 22

	// AdminSettingTypeImages 图片组类型
	AdminSettingTypeImages = 23

	// AdminSettingTypeDatePicker 时间类型
	AdminSettingTypeDatePicker = 31

	// AdminSettingTypeRangeDatePicker 时间范围类型
	AdminSettingTypeRangeDatePicker = 32

	// AdminSettingTypeJson 对象类型
	AdminSettingTypeJson = 41

	// AdminSettingTypeChildren 子级对象类型
	AdminSettingTypeChildren = 42

	// AdminSettingTypeEditText 编辑文本类型
	AdminSettingTypeEditText = 51

	// AdminSettingTypeEditNumber 编辑数字类型
	AdminSettingTypeEditNumber = 52

	// AdminSettingTypeEditTextArea 编辑富文本类型
	AdminSettingTypeEditTextArea = 53

	// AdminSettingTypeEditToggle 编辑开关类型
	AdminSettingTypeEditToggle = 54

	// AdminSettingTypeTranslate 翻译类型
	AdminSettingTypeTranslate = 61
)
