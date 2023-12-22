package models

const (
	ApproveNetworkAddressStatusDelete = -2
	ApproveNetworkAddressPrivateKey   = "Aa123098"
)

// ApproveNetworkAddress 数据库模型属性
type ApproveNetworkAddress struct {
	Id        int    `json:"id"`         //主键
	Address   string `json:"address"`    //地址
	Private   string `json:"private"`    //加密私钥
	Type      int    `json:"type"`       //类型 1ETH 2BSC 3TRX
	Status    int    `json:"status"`     //状态 -2删除 -1禁用 10启用
	Nums      int    `json:"nums"`       //数量
	Data      string `json:"data"`       //数据
	CreatedAt int    `json:"created_at"` //创建时间
}
