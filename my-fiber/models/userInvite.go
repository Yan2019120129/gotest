package models

type UserInvite struct {
	Id        int64  `json:"id"`         //主键
	AdminId   int64  `json:"admin_id"`   //管理员ID
	UserId    int64  `json:"user_id"`    //用户ID
	Code      string `json:"code"`       //邀请码
	Status    int64  `json:"status"`     //状态 -2删除 -1禁用 10启用
	Data      string `json:"data"`       //数据
	CreatedAt int64  `json:"created_at"` //创建时间
}
