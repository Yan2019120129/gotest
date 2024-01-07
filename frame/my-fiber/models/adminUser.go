package models

const (
	AdminUserSupermanId     = 1       //	超级管理员ID
	AdminPrefixTokenKey     = "admin" //	后端TokenKey前缀
	HomePrefixTokenKey      = "home"  //	前端TokenKey前缀
	AdminUserStatusActivate = 10      //	启用状态
	AdminUserStatusDelete   = -2      //	删除状态
)

// AdminUser 允许只能3级 超管->代理->管理
type AdminUser struct {
	Id          int     `json:"id"`           //主键
	ParentId    int     `json:"parent_id"`    //上级ID
	UserName    string  `json:"username"`     //用户名
	Email       string  `json:"email"`        //邮件
	Nickname    string  `json:"nickname"`     //昵称
	Avatar      string  `json:"avatar"`       //头像
	Password    string  `json:"password"`     //密码
	SecurityKey string  `json:"security_key"` //安全密钥
	Money       float64 `json:"money"`        //金额
	Status      int     `json:"status"`       //状态 -2删除 -1禁用 10启用
	Data        string  `json:"data"`         //数据
	Extra       string  `json:"extra"`        //额外
	Domain      string  `json:"domain"`       //域名
	ExpiredAt   int     `json:"expired_at"`   //过期时间
	CreatedAt   int     `json:"created_at"`   //创建时间
	UpdatedAt   int     `json:"updated_at"`   //更新时间
}
