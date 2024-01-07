package models

// ProductOrder 数据库模型属性
type ProductOrder struct {
	Id        int     `json:"id"`         //主键
	AdminId   int     `json:"admin_id"`   //管理员ID
	UserId    int     `json:"user_id"`    //用户ID
	ProductId int     `json:"product_id"` //产品ID
	OrderSn   string  `json:"order_sn"`   //订单编号
	Money     float64 `json:"money"`      //	购买金额
	Nums      float64 `json:"nums"`       //	购买数量
	Price     float64 `json:"price"`      //	买入单价
	Fee       float64 `json:"fee"`        //	手续费
	Side      int     `json:"side"`       //	方向	1买 2卖
	Mode      int     `json:"mode"`       //	模式	1币币 2合约 3期权 13机器人期权
	Type      int     `json:"type"`       //	类型	1市价 2限价
	Status    int     `json:"status"`     //	状态 	-2删除 -1取消 10待定 11运行中 20完成
	Data      string  `json:"data"`       //数据
	ExpiredAt int     `json:"expired_at"` //过期时间
	UpdatedAt int     `json:"updated_at"` //更新时间
	CreatedAt int     `json:"created_at"` //创建时间
}
