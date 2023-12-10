package models

const (
	ShopOrderStatusDelete   = -2 //	店铺订单删除
	ShopOrderStatusCancel   = -1 //	店铺订单取消
	ShopOrderStatusPayment  = 10 //	店铺订单待付款
	ShopOrderStatusPending  = 20 //	店铺订单待发货
	ShopOrderStatusShipped  = 21 //	店铺订单已发货
	ShopOrderStatusComplete = 40 //	店铺订单完成
	ShopOrderTypeWholesale  = 2  //	店铺订单批发订单
)

// ShopOrder 数据库模型属性
type ShopOrder struct {
	Id         int64   `json:"id"`          //主键
	AdminId    int64   `json:"admin_id"`    //管理员ID
	UserId     int64   `json:"user_id"`     //用户ID
	ShopId     int64   `json:"shop_id"`     //店铺ID
	OrderSn    string  `json:"order_sn"`    //订单编号
	Money      float64 `json:"money"`       //原价
	FinalMoney float64 `json:"final_money"` //成交价
	Earnings   float64 `json:"earnings"`    //订单收益
	Type       int64   `json:"type"`        //类型 1自营订单 2批发订单
	Status     int64   `json:"status"`      //状态 -2删除 -1取消订单 10待付款 20待发货 21已发货 30确认收货 40完成
	Data       string  `json:"data"`        //数据
	PaymentId  int64   `json:"payment_id"`  //支付id
	Address    string  `json:"address"`     //地址信息
	ExpiredAt  int64   `json:"expired_at"`  //过期时间
	UpdatedAt  int64   `json:"updated_at"`  //付款时间
	CreatedAt  int64   `json:"created_at"`  //创建时间
}
