package models

// ProductKline 数据库模型属性
type ProductKline struct {
	Id         int     `json:"id"`          //主键
	AdminId    int     `json:"admin_id"`    //管理员ID
	ProductId  int     `json:"product_id"`  //产品ID
	OpenPrice  float64 `json:"open_price"`  //开盘价格
	HighPrice  float64 `json:"high_price"`  //最高价格
	LowsPrice  float64 `json:"lows_price"`  //最低价格
	ClosePrice float64 `json:"close_price"` //收盘价格
	Vol        float64 `json:"vol"`         //交易量
	Amount     float64 `json:"amount"`      //成交额
	UpdatedAt  int     `json:"updated_at"`  //收盘时间
	CreatedAt  int     `json:"created_at"`  //开盘时间
}
