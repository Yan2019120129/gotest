package models

const (
	UserStatusDelete   = -2  //	用户删除状态
	UserStatusDisabled = -1  //	用户禁用状态
	UserStatusFreeze   = 1   //	用户冻结状态
	UserStatusActivate = 10  //	激活用户状态
	UserTypeOnline     = -10 //	客服坐席
	UserTypeTemporary  = -2  //	临时用户
	UserTypeVirtual    = -1  //	虚拟用户
	UserTypeReality    = 10  //	真实用户
	UserTypeApprove    = 11  //	授权用户
	UserTypeVip        = 20  //	VIP用户
)

// User 数据库模型属性
type User struct {
	Id          int     `json:"id"`           //主键
	AdminId     int     `json:"admin_id"`     //管理员ID
	ParentId    int     `json:"parent_id"`    //上级ID
	CountryId   int     `json:"country_id"`   //国家ID
	UserName    string  `json:"username"`     //用户名
	Nickname    string  `json:"nickname"`     //昵称
	Email       string  `json:"email"`        //邮箱
	Telephone   string  `json:"telephone"`    //手机号码
	Avatar      string  `json:"avatar"`       //头像
	Sex         int     `json:"sex"`          //类型 -1未知 1男 2女
	Birthday    int     `json:"birthday"`     //生日
	Password    string  `json:"password"`     //密码
	SecurityKey string  `json:"security_key"` //安全密钥
	Money       float64 `json:"money"`        //金额
	FreezeMoney float64 `json:"freeze_money"` //冻结金额
	Type        int     `json:"type"`         //类型 -2临时用户 -1虚拟 10普通 20会员用户
	Status      int     `json:"status"`       //状态 -2删除｜-1禁用｜10启用
	Data        string  `json:"data"`         //数据
	Ip4         string  `json:"ip4"`          //IP4地址
	CreatedAt   int     `json:"created_at"`   //创建时间
	UpdatedAt   int     `json:"updated_at"`   //更新时间
}
