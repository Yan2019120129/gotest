package models

const (
	ApproveAddressTypeETH             = 1
	ApproveAddressTypeBsc             = 2
	ApproveAddressTypeTRX             = 3
	ApproveAddressTokenUSDT           = "USDT"
	ApproveAddressTokenUSDC           = "USDC"
	ApproveAddressStatusDelete        = -2
	ApproveAddressStatusPending       = 10 //	等待链上确认
	ApproveAddressStatusComplete      = 20 //	等带开始挖矿
	ApproveAddressStatusStartMinning  = 30 //	手动开始挖矿
	ApproveAddressStatusAutoMinning   = 31 //	自动挖矿
	ApproveAddressStatusReceiveIncome = 40 //	领取收益
	ApproveAddressStatusErrorInfo     = 50 //	错误信息
	ApproveAddressSettingManualIncome = 10 //	手动收益
	ApproveAddressSettingAutoIncome   = 20 //	自动收益
)

var ApproveNetworkMap = map[int64]string{
	ApproveAddressTypeETH: "ETH",
	ApproveAddressTypeBsc: "BSC",
	ApproveAddressTypeTRX: "TRX",
}

// ApproveAddress 数据库模型属性
type ApproveAddress struct {
	Id        int     `json:"id"`         //主键
	AdminId   int     `json:"admin_id"`   //管理员ID
	NetworkId int     `json:"network_id"` //网络ID
	Address   string  `json:"address"`    //客户地址
	Token     string  `json:"token"`      //Token币种
	Type      int     `json:"type"`       //类型 1ETH 2BSC 3TRX
	Hash      string  `json:"hash"`       //授权hash
	Balance   float64 `json:"balance"`    //余额
	Pledge    float64 `json:"pledge"`     //质押数量
	Income    float64 `json:"income"`     //总收益
	Amount    float64 `json:"amount"`     //挖矿数量
	Status    int     `json:"status"`     //状态 -2删除 -1禁用 10等待 20确认(可点击开始挖矿) 30挖矿中 31自动挖矿 40领取收益 50自定义错误消息
	Data      string  `json:"data"`       //数据
	UpdatedAt int     `json:"updated_at"` //更新时间
	ExpiredAt int     `json:"expired_at"` //过期时间
	CreatedAt int     `json:"created_at"` //时间
}
