package models

const (
	ProductCategoryStatusActivate = 10 //	分类状态启用
	ProductCategoryStatusDisabled = -1 //	分类状态禁用
	ProductCategoryStatusDelete   = -2 //	分类状态删除
	ProductCategoryRecommend      = 10 //	分类推荐
	ProductCategoryTypeDefault    = 1  //	数字货币
)

var ProductCategoryTypeList map[int64]string = map[int64]string{
	ProductCategoryTypeDefault: "数字货币",
}

// ProductCategory 数据库模型属性
type ProductCategory struct {
	Id        int    `json:"id"`         //主键
	ParentId  int    `json:"parent_id"`  //分类上级ID
	AdminId   int    `json:"admin_id"`   //管理员ID
	Type      int    `json:"type"`       //类型 1数字货币
	Name      string `json:"name"`       //标题
	Image     string `json:"image"`      //封面
	Sort      int    `json:"sort"`       //排序
	Status    int    `json:"status"`     //状态 -2删除 -1禁用 10启用
	Recommend int    `json:"recommend"`  //推荐 -1关闭 10推荐
	Data      string `json:"data"`       //数据
	UpdatedAt int    `json:"updated_at"` //更新时间
	CreatedAt int    `json:"created_at"` //创建时间
}
