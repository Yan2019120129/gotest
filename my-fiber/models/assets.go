package models

const (
	AssetsStatusDelete   = -2
	AssetsStatusDisabled = -1
	AssetsStatusActivate = 10
	AssetsTypeDigital    = 1      //	数字货币
	AssetsTypeBank       = 10     //	银行货币
	AssetsTypeSystem     = 20     //	系统货币
	AssetsTypeAssets     = 30     //	资产货币
	MainAssetsUSDT       = "USDT" //	主资产名称
)

// AssetsTypeList 资产列表
var AssetsTypeList = map[int64]string{
	AssetsTypeDigital: "数字货币", AssetsTypeBank: "银行货币", AssetsTypeSystem: "系统货币",
}

// Assets 数据库模型属性
type Assets struct {
	Id        int    `json:"id"`         //主键
	AdminId   int    `json:"admin_id"`   //管理员ID
	Name      string `json:"name"`       //名称
	Icon      string `json:"icon"`       //图标
	Type      int    `json:"type"`       //类型 1ETH 2BSC 3TRX
	Data      string `json:"data"`       //数据
	Status    int    `json:"status"`     //状态 -2删除｜-1禁用｜10启用
	CreatedAt int    `json:"created_at"` //创建时间
	UpdatedAt int    `json:"updated_at"` //更新时间
}
