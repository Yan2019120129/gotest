package models

// UserSetting 数据库模型属性
type UserSetting struct {
	Id      int    `json:"id"`       //主键
	AdminId int    `json:"admin_id"` //管理ID
	UserId  int    `json:"user_id"`  //用户ID
	GroupId int    `json:"group_id"` //组ID
	Name    string `json:"name"`     //名称
	Type    string `json:"type"`     //类型
	Field   string `json:"field"`    //健名
	Value   string `json:"value"`    //健值
	Data    string `json:"data"`     //数据
}
