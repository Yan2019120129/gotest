package models

const (
	UserLevelStatusActivate = 10
	UserLevelStatusDisabled = -1
	UserLevelStatusDelete   = -2
)

type UserLevel struct {
	Id        int     `json:"id"`         //主键
	AdminId   int     `json:"admin_id"`   //管理员ID
	Name      string  `json:"name"`       //名称
	Icon      string  `json:"icon"`       //图标
	Level     int     `json:"level"`      //等级
	Money     float64 `json:"money"`      //购买金额
	Days      int     `json:"days"`       //购买天数 -1无限时间
	Status    int     `json:"status"`     //状态 -2删除 -1禁用 10启用
	Data      string  `json:"data"`       //数据
	CreatedAt int     `json:"created_at"` //创建时间
	UpdatedAt int     `json:"updated_at"` //更新时间
}
