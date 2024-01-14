package models

// ProductCollect 产品收藏表
type ProductCollect struct {
	Id        int    `gorm:"type: int unsigned primary key auto_increment; comment: 主键"`
	AdminId   int    `gorm:"type: int unsigned not null; comment: 管理员ID"`
	ProductId int    `gorm:"type: int unsigned not null; comment: 产品ID"`
	UserId    int    `gorm:"type: int unsigned not null; comment: 用户ID"`
	Data      string `gorm:"type: text; comment: 数据"`
	Status    int    `gorm:"type: tinyint not null; default: 10; comment: 状态 -2删除 10收藏"`
	UpdatedAt int    `gorm:"type: int unsigned not null; autoUpdateTime; comment: 收盘时间"`
	CreatedAt int    `gorm:"type: int unsigned not null; autoCreateTime; comment: 开盘时间"`
}

const (
	ProductCollectStatusActivate = 10 //	收藏
	ProductCollectStatusDelete   = -2 //	删除
)
