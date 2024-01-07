package models

const (
	UserWalletAccountStatusActivate = 10
	UserWalletAccountStatusDelete   = -2
	UserWalletAccountStatusDisabled = -1
)

type UserWalletAccount struct {
	Id         int    `json:"id"`          //主键
	AdminId    int    `json:"admin_id"`    //管理员ID
	UserId     int    `json:"user_id"`     //用户ID
	PaymentId  int    `json:"payment_id"`  //提现方式ID
	Name       string `json:"name"`        //名称
	RealName   string `json:"real_name"`   //真实名字
	CardNumber string `json:"card_number"` //卡号
	Address    string `json:"address"`     //地址
	Status     int    `json:"status"`      //状态 -2删除 -1禁用 10启用
	Sort       int    `json:"sort"`        //排序
	Data       string `json:"data"`        //数据
	UpdatedAt  int    `json:"updated_at"`  //更新时间
	CreatedAt  int    `json:"created_at"`  //创建时间
}
