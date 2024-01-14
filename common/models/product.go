package models

const (
	ProductNumsUnlimited  = -1 //	没有限制购买
	ProductStatusActivate = 10 //	状态上架
	ProductStatusDisabled = -1 //	状态下架
	ProductStatusDelete   = -2 //	状态删除
	ProductTypeDefault    = 1  //	自营商品
	ProductTypeWholesale  = 2  //	批发商品
)

// ProductShop 数据库模型属性
type ProductShop struct {
	Id            int64   `json:"id"`             //主键
	ParentId      int64   `json:"parent_id"`      //上级ID
	AdminId       int64   `json:"admin_id"`       //管理员ID
	CategoryId    int64   `json:"category_id"`    //类目ID
	ShopId        int64   `json:"shop_id"`        //店铺ID 0批发中心
	AssetsId      int64   `json:"assets_id"`      //资产ID
	Name          string  `json:"name"`           //标题
	Images        string  `json:"images"`         //图片列表
	OriginalMoney float64 `json:"original_money"` //原价
	Money         float64 `json:"money"`          //现价
	Increase      float64 `json:"increase"`       //涨幅
	Type          int64   `json:"type"`           //类型 1自营商品 2批发商品
	Sort          int64   `json:"sort"`           //排序
	Status        int64   `json:"status"`         //状态 -2删除 -1禁用 10启用
	Sales         int64   `json:"sales"`          //销售量
	Rating        float64 `json:"rating"`         //评分
	Nums          int64   `json:"nums"`           //限购 -1 无限制
	Used          int64   `json:"used"`           //已用
	Total         int64   `json:"total"`          //总数
	Data          string  `json:"data"`           //数据
	Describes     string  `json:"describes"`      //描述
	UpdatedAt     int64   `json:"updated_at"`     //更新时间
	CreatedAt     int64   `json:"created_at"`     //创建时间
}

// HomeProductInfo 前台产品信息
type HomeProductInfo struct {
	ProductInfo
	IsFollow bool `json:"isFollow"` // 是否关注产品
}

// StoreProductInfo 商家批发产品信息
type StoreProductInfo struct {
	ProductInfo
	Status   int64   `json:"status"`   //	产品状态
	Stock    int64   `json:"stock"`    //	产品库存
	Increase float64 `json:"increase"` //	涨幅
}

// ProductInfo 产品信息
type ProductInfo struct {
	Id            int64   `json:"id"`            //	产品ID
	ShopId        int64   `json:"shopId"`        //	店铺ID
	Image         string  `json:"image"`         //	产品图片
	Name          string  `json:"name"`          //	产品名称
	Money         float64 `json:"money"`         //	产品现价
	OriginalMoney float64 `json:"originalMoney"` //	产品原价
	Type          int64   `json:"type"`          //	产品类型
	Sales         int64   `json:"sales"`         //	产品销量
	CreatedAt     int64   `json:"createdAt"`     //	创建时间
}

// ProductImage 产品图片
type ProductImage struct {
	Label string `json:"label"` //	图片说明
	Value string `json:"value"` //	图片地址
}
