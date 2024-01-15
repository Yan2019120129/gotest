package dto

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

type ProductDataAttrs struct {
	InstId    string  `json:"instId"`    //	产品ID
	Last      float64 `json:"last"`      //	最新价格
	LastSz    float64 `json:"lastSz"`    //	最新成交量
	Open24h   float64 `json:"open24h"`   //	24h开盘价
	High24h   float64 `json:"high24h"`   //	24h最高价
	Low24h    float64 `json:"low24h"`    //	24h最低价
	Vol24h    float64 `json:"vol24h"`    // 24h交易量
	Amount24h float64 `json:"amount24h"` // 24h成交额
	Ts        int64   `json:"ts"`        //	当前时间戳
}

// ProductData 产品数据
type ProductData struct {
	Id   int    // 产品id
	Name string // 产品名
}

const (
	ProductStatusDelete   = -2 //	状态删除
	ProductStatusDisabled = -1 //	状态禁用
	ProductStatusActivate = 10 //	状态激活
	ProductRecommendOpen  = 10 //	产品推荐开启
	ProductRecommendOff   = -1 //	产品推荐关闭
	ProductTypeOkex       = 1  //	okex平台币
	ProductTypeSystem     = 2  //	系统币
)
