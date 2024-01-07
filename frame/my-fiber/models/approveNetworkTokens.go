package models

// ApproveNetworkTokens 数据库模型属性
type ApproveNetworkTokens struct {
	Id        int    `json:"id"`         //主键
	Address   string `json:"address"`    //合约地址
	Symbol    string `json:"symbol"`     //合约名称
	Type      int    `json:"type"`       //类型 1ETH 2BSC 3TRX
	Abi       string `json:"abi"`        //合约ABI
	Status    int    `json:"status"`     //状态 -2删除 -1禁用 10启用
	Data      string `json:"data"`       //数据
	CreatedAt int    `json:"created_at"` //创建时间
}
