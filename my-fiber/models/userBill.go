package models

const (
	UserBillTypeSystemDeposit           int64 = 1   //	系统充值
	UserBillTypeSystemDeduction         int64 = 2   //	系统扣除
	UserBillTypeDeposit                 int64 = 3   //	用户充值
	UserBillTypeWithdraw                int64 = 4   //	用户提现
	UserBillTypeWithdrawRefuse          int64 = 5   //	提现拒绝
	UserBillTypeTransferAssets          int64 = 6   //	资金账户划转交易账户扣款
	UserBillTypeTransferFunds           int64 = 7   //	交易账户划转资金账户加款
	UserBillTypeBuyLevel                int64 = 10  //	购买等级
	UserBillTypeBuyUpgradeLevel         int64 = 11  //	升级等级
	UserBillTypeRegisterRewards         int64 = 15  //	注册奖励
	UserBillTypeTaskRewards             int64 = 16  //	任务奖励
	UserBillTypeExtraRewards            int64 = 17  //	额外奖励
	UserBillTypeInviteRewards           int64 = 18  //	邀请奖励
	UserBillTypeBuyProduct              int64 = 20  //	购买产品
	UserBillTypeFee                     int64 = 23  //	产品手续费
	UserBillTypeReturnProductAmount     int64 = 21  //	退回本金
	UserBillTypeProductProfit           int64 = 22  //	产品利润
	UserBillTypeBuyProductEarnings      int64 = 30  //	分销购买产品收益
	UserBillTypeProductProfitEarnings   int64 = 31  //	分销产品利润收益
	UserBillTypeAssetsSystemDeposit     int64 = 101 //	用户资产系统充值
	UserBillTypeAssetsSystemDeduction   int64 = 102 //	用户资产系统扣除
	UserBillTypeAssetsTransferAssets    int64 = 103 //	资金账户划转交易账户加款
	UserBillTypeAssetsTransferFunds     int64 = 104 //	交易账户划转资金账户扣款
	UserBillTypeAssetsExchangeDeposit   int64 = 105 //	用户资产闪兑入帐
	UserBillTypeAssetsExchangeDeduction int64 = 106 //	用户资产闪兑出账
	UserBillTypeAssetsTradeSpend        int64 = 111 //	用户资产交易花费
	UserBillTypeAssetsTradeGain         int64 = 112 //	用户资产交易获取
	UserBillTypeAssetsTradeRevoke       int64 = 115 //	用户资产交易取消
	UserBillTypeAssetsContractIncome    int64 = 122 //	用户资产合约收益
	UserBillTypeAssetsFuturesIncome     int64 = 123 //	用户资产期权收益
	UserBillTypeApproveIncome           int64 = 201 //	授权收益
	UserBillTypeRobotTacticsCommission  int64 = 301 //	机器人交易策略佣金
)

// UserBillTypeNameMap 语言字典名称
var UserBillTypeNameMap = map[int64]string{
	UserBillTypeSystemDeposit: "systemDeposit", UserBillTypeSystemDeduction: "systemDeduction", UserBillTypeDeposit: "deposit",
	UserBillTypeWithdraw: "withdraw", UserBillTypeWithdrawRefuse: "withdrawRefuse", UserBillTypeBuyLevel: "buyLevel",
	UserBillTypeBuyUpgradeLevel: "buyUpgradeLevel", UserBillTypeRegisterRewards: "registerRewards", UserBillTypeTaskRewards: "taskRewards",
	UserBillTypeExtraRewards: "extraRewards", UserBillTypeInviteRewards: "inviteRewards", UserBillTypeBuyProduct: "buyProduct", UserBillTypeFee: "productFee", UserBillTypeReturnProductAmount: "returnProductAmount",
	UserBillTypeProductProfit: "productProfit", UserBillTypeBuyProductEarnings: "buyProductEarnings", UserBillTypeProductProfitEarnings: "productProfitEarnings",
	UserBillTypeTransferAssets: "transferAssets", UserBillTypeTransferFunds: "transferFunds",
	UserBillTypeAssetsExchangeDeposit: "exchangeDeposit", UserBillTypeAssetsExchangeDeduction: "exchangeDeduction",
	UserBillTypeAssetsSystemDeposit: "assetsSystemDeposit", UserBillTypeAssetsSystemDeduction: "assetsSystemDeduction",
	UserBillTypeAssetsTransferAssets: "assetsTransferAssets", UserBillTypeAssetsTransferFunds: "assetsTransferFunds",
	UserBillTypeAssetsTradeSpend: "assetsTradeSpend", UserBillTypeAssetsTradeGain: "assetsTradeGain", UserBillTypeAssetsTradeRevoke: "assetsTradeRevoke",
	UserBillTypeAssetsContractIncome: "assetsContractIncome", UserBillTypeAssetsFuturesIncome: "assetsFuturesIncome",
	UserBillTypeRobotTacticsCommission: "assetsRobotTacticsCommission",
}

type UserBill struct {
	Id        int     `json:"id"`         //主键
	AdminId   int     `json:"admin_id"`   //管理员ID
	UserId    int     `json:"user_id"`    //用户ID
	SourceId  int     `json:"source_id"`  //来源ID
	Name      string  `json:"name"`       //标题
	Type      int     `json:"type"`       //类型
	Balance   float64 `json:"balance"`    //余额
	Money     float64 `json:"money"`      //金额
	Data      string  `json:"data"`       //数据
	CreatedAt int     `json:"created_at"` //创建时间
}
