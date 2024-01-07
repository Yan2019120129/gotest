package models

const (
	CountryStatusActivate = 10
	CountryStatusDisabled = -1
)

// Country 数据库模型属性
type Country struct {
	Id        int    `json:"id"`         //主键
	AdminId   int    `json:"admin_id"`   //管理员ID
	LangId    int    `json:"lang_id"`    //语言ID
	Name      string `json:"name"`       //名称
	Alias     string `json:"alias"`      //别名
	Iso1      string `json:"iso1"`       //	ISO3166-1
	Icon      string `json:"icon"`       //图标
	Code      string `json:"code"`       //区号
	Sort      int    `json:"sort"`       //排序
	Status    int    `json:"status"`     //状态 -1禁用｜10启用
	Data      string `json:"data"`       //数据
	CreatedAt int    `json:"created_at"` //创建时间
}
