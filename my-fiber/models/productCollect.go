package models

const (
	ProductCollectStatusActivate = 10 //	收藏
	ProductCollectStatusDelete   = -2 //	删除
)

// ProductCollect 数据库模型属性
type ProductCollect struct {
	Id        int    `json:"id"`         //主键
	AdminId   int    `json:"admin_id"`   //管理员ID
	ProductId int    `json:"product_id"` //产品ID
	UserId    int    `json:"user_id"`    //用户ID
	Data      string `json:"data"`       //数据
	Status    int    `json:"status"`     //状态 -2删除 10收藏
	UpdatedAt int    `json:"updated_at"` //收盘时间
	CreatedAt int    `json:"created_at"` //开盘时间
}
