package models

const (
	ProductOrderStatusDelete   = -2 //	订单状态删除
	ProductOrderStatusCancel   = -1 //	订单状态取消
	ProductOrderStatusPayment  = 10 //	订单状态备货
	ProductOrderStatusPending  = 11 //	订单状态出库
	ProductOrderStatusReceipt  = 12 //	订单状态收货
	ProductOrderStatusRefund   = 16 //	订单状态退货
	ProductOrderStatusRecycle  = 17 //	订单状态入库
	ProductOrderStatusComplete = 20 //	订单状态完成
)

// ProductOrder 数据库模型属性
type ProductOrder struct {
	Id          int64   `json:"id"`            //主键
	ShopOrderId int64   `json:"shop_order_id"` //店铺订单ID
	AdminId     int64   `json:"admin_id"`      //管理员ID
	UserId      int64   `json:"user_id"`       //用户ID
	SkuId       int64   `json:"sku_id"`        //skuID
	ProductId   int64   `json:"product_id"`    //产品ID
	OrderSn     string  `json:"order_sn"`      //订单编号
	Money       float64 `json:"money"`         //原价
	FinalMoney  float64 `json:"final_money"`   //成交价
	Nums        int64   `json:"nums"`          //数量
	Type        int64   `json:"type"`          //类型 1默认订单
	Status      int64   `json:"status"`        //状态 -2删除 -1取消 [10备货 11出库] 12已收货 [16退货 17入库] 20完成
	Data        string  `json:"data"`          //数据
	ExpiredAt   int64   `json:"expired_at"`    //过期时间
	UpdatedAt   int64   `json:"updated_at"`    //更新时间
	CreatedAt   int64   `json:"created_at"`    //创建时间
}
