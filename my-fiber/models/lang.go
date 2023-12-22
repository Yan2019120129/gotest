package models

const (
	LangStatusActivate = 10
	LangStatusDisabled = -1
)

// Lang 数据库模型属性
type Lang struct {
	Id        int    `json:"id"`         //主键
	AdminId   int    `json:"admin_id"`   //管理员ID
	Name      string `json:"name"`       //名称
	Alias     string `json:"alias"`      //别名
	Icon      string `json:"icon"`       //图标
	Sort      int    `json:"sort"`       //排序
	Status    int    `json:"status"`     //状态 1禁用｜10启用
	Data      string `json:"data"`       //数据
	CreatedAt int    `json:"created_at"` //创建时间
}
