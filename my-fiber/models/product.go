package models

const (
	ProductNumsUnlimited  = -1 //	没有限制购买
	ProductStatusActivate = 10 //	状态激活
	ProductStatusDisabled = -1 //	状态禁用
	ProductStatusDelete   = -2 //	状态删除
	ProductRecommend      = 10 //	产品推荐
	ProductTypeOkex       = 1  //	okex平台币
	ProductTypeSystem     = 2  //	系统币
)

// Product 数据库模型属性
type Product struct {
	Id             int     `json:"id"`               //主键
	AdminId        int     `json:"admin_id"`         //管理员ID
	CategoryId     int     `json:"category_id"`      //类目ID
	AssetsId       int     `json:"assets_id"`        //资产ID
	SymbolAssetsId int     `json:"symbol_assets_id"` //标识资产ID
	Name           string  `json:"name"`             //标题
	Symbol         string  `json:"symbol"`           //标识
	Images         string  `json:"images"`           //图片列表
	Money          float64 `json:"money"`            //金额
	Type           int     `json:"type"`             //类型 1默认
	Sort           int     `json:"sort"`             //排序
	Status         int     `json:"status"`           //状态 -2删除 -1禁用 10启用
	Recommend      int     `json:"recommend"`        //推荐 -1关闭 10推荐
	Sales          int     `json:"sales"`            //销售量
	Nums           int     `json:"nums"`             //限购 -1无限
	Used           int     `json:"used"`             //已使用
	Total          int     `json:"total"`            //总数
	Data           string  `json:"data"`             //数据
	Describes      string  `json:"describes"`        //数据
	UpdatedAt      int     `json:"updated_at"`       //更新时间
	CreatedAt      int     `json:"created_at"`       //创建时间
}
