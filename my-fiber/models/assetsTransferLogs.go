package models

const (
	AssetsTransferLogsTypeTransfer = 1 //	划转
	AssetsTransferLogsTypeExchange = 2 //	闪兑
)

// AssetsTransferLogs 数据库模型属性
type AssetsTransferLogs struct {
	Id             int     `json:"id"`               //主键
	AdminId        int     `json:"admin_id"`         //管理员ID
	UserId         int     `json:"user_id"`          //用户ID
	Type           int     `json:"type"`             //类型 1划转 2闪兑
	AssetsId       int     `json:"assets_id"`        //资产ID
	SymbolAssetsId int     `json:"symbol_assets_id"` //转化资产ID
	Money          float64 `json:"money"`            //金额
	Amount         float64 `json:"amount"`           //数量
	Data           string  `json:"data"`             //数据
	CreatedAt      int     `json:"created_at"`       //创建时间
}
