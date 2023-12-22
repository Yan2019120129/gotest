package models

const (
	UserLevelOrderStatusActivate = 10
)

type UserLevelOrder struct {
	Id        int    `json:"id"`         //主键
	AdminId   int    `json:"admin_id"`   //管理员ID
	UserId    int    `json:"user_id"`    //用户ID
	LevelId   int    `json:"level_id"`   //等级ID
	Data      string `json:"data"`       //数据
	Status    int    `json:"status"`     //状态 -2删除 -1禁用 10启用
	CreatedAt int    `json:"created_at"` //创建时间
	UpdatedAt int    `json:"updated_at"` //更新时间
}
