package models

const (
	UserAssetsStatusActivate = 10
	UserAssetsStatusDelete   = -2
)

// UserAssets 数据库模型属性
type UserAssets struct {
	Id          int     `json:"id"`           //主键
	AdminId     int     `json:"admin_id"`     //管理员ID
	UserId      int     `json:"user_id"`      //用户ID
	AssetsId    int     `json:"assets_id"`    //资产ID
	Money       float64 `json:"money"`        //金额
	FreezeMoney float64 `json:"freeze_money"` //冻结金额
	Data        string  `json:"data"`         //数据
	Status      int     `json:"status"`       //状态 -2删除｜-1禁用｜10启用
	CreatedAt   int     `json:"created_at"`   //创建时间
	UpdatedAt   int     `json:"updated_at"`   //更新时间
}
