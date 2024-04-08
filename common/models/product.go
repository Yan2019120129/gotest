package models

import "github.com/brianvoe/gofakeit/v6"

// Product 产品表
type Product struct {
	Id                int `gorm:" type:int unsigned primary key auto_increment; comment: 主键"`
	AdminUserId       int `gorm:" type:int unsigned not null; comment: 管理员ID"`
	AdminUser         *AdminUser
	ProductCategoryId int `gorm:" type:int unsigned not null; comment: 类目ID"`
	ProductCategory   *ProductCategory
	WalletAssetsId    int `gorm:" type:int unsigned not null; comment: 资产ID"`
	WalletAssets      *WalletAssets
	Name              string  `gorm:" type:varchar(64) not null; comment: 标题"`
	Images            string  `gorm:" type:varchar(2048) not null; comment: 图片列表"`
	Money             float64 `gorm:" type:decimal(12,2) not null; comment: 金额"`
	Type              int     `gorm:" type:tinyint not null; default: 1; comment: 类型 1默认"`
	Sort              int     `gorm:" type:int unsigned not null; comment: 排序"`
	Status            int     `gorm:" type:tinyint not null; default: 10; comment: 状态 -2删除 -1禁用 10启用"`
	Recommend         int     `gorm:" type:tinyint not null; default: -1; comment: 推荐 -1关闭 10推荐"`
	Sales             int     `gorm:" type:int unsigned not null; comment: 销售量"`
	Nums              int     `gorm:" type:tinyint not null; default: -1; comment: 限购 -1无限"`
	Used              int     `gorm:" type:int unsigned not null; comment: 已使用"`
	Total             int     `gorm:" type:int unsigned not null; comment: 总数"`
	Data              string  `gorm:" type:text; comment: 数据"`
	Describes         string  `gorm:" type:text; comment: 描述"`
	UpdatedAt         int     `gorm:" type:int unsigned not null; autoUpdateTime; comment: 更新时间"`
	CreatedAt         int     `gorm:" type:int unsigned not null; autoCreateTime; comment: 创建时间"`
}

// GetProductDefault 获取默认数据
func GetProductDefault() *Product {
	return &Product{
		Id:        0,
		Name:      gofakeit.Name(),
		Images:    gofakeit.ImageURL(200, 400),
		Money:     gofakeit.Float64Range(10, 30),
		Type:      gofakeit.RandomInt([]int{1, 3}),
		Sort:      gofakeit.Number(1, 100),
		Status:    gofakeit.RandomInt([]int{-2, -1, 20}),
		Recommend: gofakeit.RandomInt([]int{-1, 20}),
		Sales:     gofakeit.Number(1, 500),
		Nums:      -1,
		Used:      gofakeit.Number(1, 500),
		Total:     gofakeit.Number(500, 600),
		Data:      gofakeit.Letter(),
		Describes: gofakeit.Letter(),
	}
}

// SetAdminId 设置管理员Id
func (p *Product) SetAdminId(id int) *Product {
	p.AdminUserId = id
	return p
}

// SetProductCategory 设置分类Id
func (p *Product) SetProductCategory(id int) *Product {
	p.ProductCategoryId = id
	return p
}

// SetWalletAssetsId 设置钱包资产Id
func (p *Product) SetWalletAssetsId(id int) *Product {
	p.WalletAssetsId = id
	return p
}
