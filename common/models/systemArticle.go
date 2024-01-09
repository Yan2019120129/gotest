package models

// Article 系统文章管理
type Article struct {
	Id        int    `gorm:"type:int unsigned primary key auto_increment;comment:主键;"`
	AdminId   int    `gorm:"type:int unsigned not null;comment:管理ID"`
	Type      int    `gorm:"type:smallint unsigned not null default 1;default:10;comment:类型 1帮助中心类型 2底部类型 10基础文章"`
	Image     string `gorm:"type:varchar(255) not null;comment:图片"`
	Name      string `gorm:"type:varchar(60) not null;comment:标题"`
	Content   string `gorm:"type:text;comment:内容"`
	Status    int    `gorm:"type:smallint not null default 10;default:10;comment:状态 -2删除 -1禁用 10开启"`
	Data      string `gorm:"type:text;comment:数据"`
	UpdatedAt int    `gorm:"type:int unsigned not null;autoUpdateTime;comment:更新时间"`
	CreatedAt int    `gorm:"type:int unsigned not null;autoCreateTime;comment:创建时间"`
}

const (

	// ArticleTypeHelpers 帮助中心
	ArticleTypeHelpers = 1

	// ArticleTypeFooter 底部类型
	ArticleTypeFooter = 2

	// ArticleTypeBasic 基础文章
	ArticleTypeBasic = 10

	// ArticleStatusActive 开启
	ArticleStatusActive = 10

	// ArticleStatusDisable 禁用
	ArticleStatusDisable = -1

	// ArticleStatusDelete 删除
	ArticleStatusDelete = -2
)
