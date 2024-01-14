package models

// Product 产品表
type Product struct {
	Id         int     `gorm:" type:int unsigned primary key auto_increment; comment: 主键"`
	AdminId    int     `gorm:" type:int unsigned not null; comment: 管理员ID"`
	CategoryId int     `gorm:" type:int unsigned not null; comment: 类目ID"`
	AssetsId   int     `gorm:" type:int unsigned not null; comment: 资产ID"`
	Name       string  `gorm:" type:varchar(64) not null; comment: 标题"`
	Images     string  `gorm:" type:varchar(2048) not null; comment: 图片列表"`
	Money      float64 `gorm:" type:decimal(12,2) not null; comment: 金额"`
	Type       int     `gorm:" type:tinyint not null; default: 1; comment: 类型 1默认"`
	Sort       int     `gorm:" type:int unsigned not null; comment: 排序"`
	Status     int     `gorm:" type:tinyint not null; default: 10; comment: 状态 -2删除 -1禁用 10启用"`
	Recommend  int     `gorm:" type:tinyint not null; default: -1; comment: 推荐 -1关闭 10推荐"`
	Sales      int     `gorm:" type:int unsigned not null; comment: 销售量"`
	Nums       int     `gorm:" type:tinyint not null; default: -1; comment: 限购 -1无限"`
	Used       int     `gorm:" type:int unsigned not null; comment: 已使用"`
	Total      int     `gorm:" type:int unsigned not null; comment: 总数"`
	Data       string  `gorm:" type:text; comment: 数据"`
	Describes  string  `gorm:" type:text; comment: 描述"`
	UpdatedAt  int     `gorm:" type:int unsigned not null; autoUpdateTime; comment: 更新时间"`
	CreatedAt  int     `gorm:" type:int unsigned not null; autoCreateTime; comment: 创建时间"`
}
