package models

import "github.com/brianvoe/gofakeit/v6"

// ProductCategory 产品分类表
type ProductCategory struct {
	Id          int    `gorm:"type: int unsigned primary key auto_increment; comment: 主键"`
	ParentId    int    `gorm:"type:int unsigned not null; comment: 分类上级ID"`
	AdminUserId int    `gorm:"type:int unsigned not null; comment: 管理员ID"`
	Type        int    `gorm:"type:tinyint not null; default: 1; comment: 类型 1数字货币,2外汇，3期权"`
	Name        string `gorm:"type:varchar(60) not null; comment: 标题"`
	Icon        string `gorm:"type:varchar(60) not null; comment: 封面"`
	Sort        int    `gorm:"type:tinyint not null;default: 99; comment:排序"`
	Recommend   int    `gorm:"type:tinyint not null; default: -1; comment: 推荐 -1关闭 10推荐"`
	Status      int    `gorm:"type:tinyint not null; default: 10; comment: 状态 -2删除 -1禁用 10启用"`
	Data        string `gorm:"type:text; comment: 数据"`
	UpdatedAt   int    `gorm:"type:int unsigned not null;autoUpdateTime; comment: 更新时间"`
	CreatedAt   int    `gorm:"type:int unsigned not null;autoCreateTime; comment: 创建时间"`
}

// GetProductCategoryDefault 获取产品默认数据
func GetProductCategoryDefault() *ProductCategory {
	return &ProductCategory{
		Type:      gofakeit.RandomInt([]int{1, 2, 3}),
		Name:      gofakeit.BookTitle(),
		Icon:      gofakeit.ImageURL(200, 400),
		Sort:      gofakeit.Number(1, 100),
		Recommend: gofakeit.RandomInt([]int{-1, 10}),
		Status:    gofakeit.RandomInt([]int{-2, -1, 10}),
		Data:      gofakeit.Letter(),
	}
}

// SetParentId 设置上级Id
func (p *ProductCategory) SetParentId(id int) *ProductCategory {
	p.ParentId = id
	return p
}

// SetAdminId 设置管理员Id
func (p *ProductCategory) SetAdminId(id int) *ProductCategory {
	p.AdminUserId = id
	return p
}

const (
	ProductCategoryStatusActivate = 10 //	分类状态启用
	ProductCategoryStatusDisabled = -1 //	分类状态禁用
	ProductCategoryStatusDelete   = -2 //	分类状态删除
	ProductCategoryRecommendOff   = -1 //	分类推荐关闭
	ProductCategoryRecommendOpen  = 10 //	分类推荐开启
	ProductCategoryTypeDefault    = 1  //	数字货币
	ProductCategoryTypeForex      = 2  //	外汇
	ProductCategoryTypeOptions    = 3  //	期权
)
