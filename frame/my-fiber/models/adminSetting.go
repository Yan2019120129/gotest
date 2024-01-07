package models

const (
	SettingTypeText                          = "text"
	SettingTypeNumber                        = "number"
	SettingTypeEditor                        = "editor"
	SettingTypeImage                         = "image"
	SettingTypeImages                        = "images"
	SettingTypeSelect                        = "select"
	SettingTypeCheckbox                      = "checkbox"
	SettingTypeChildren                      = "children"
	SettingTypeJson                          = "json"
	SettingGroupBasic                        = 1                      //	基本设置
	SettingGroupHome                         = 2                      //	首页设置
	SettingGroupFinance                      = 3                      //	财务设置
	SettingGroupTemplate                     = 4                      //	模版设置
	SettingGroupHelpers                      = 5                      //	帮助中心
	SettingGroupExchange                     = 6                      //	交易设置
	SettingGroupApprove                      = 7                      //	授权设置
	SettingGroupRobot                        = 8                      //	机器人设置
	UpdateAdminTokenParamsField              = "site_token"           //	前端Token健铭
	AdminSettingBuyLevelModePremium          = "premium"              //	补价模式
	AdminSettingBuyLevelModeEquivalence      = "equivalence"          //	等价模式
	AdminSettingProductEarningsModeManual    = "manual"               //	产品收益模式【手动】
	AdminSettingProductEarningsModeAutomatic = "automatic"            //	产品收益模式【自动】
	AdminSettingSiteName                     = "site_name"            //	站点名称
	AdminSettingIntroduce                    = "home_introduce"       // 站点介绍
	AdminSettingNotice                       = "home_notice"          //	站点公告
	AdminSettingPrivacyPolicy                = "home_privacy"         //	站点隐私
	AdminSettingServiceAgreement             = "home_protocol"        //	站点协议
	AdminSettingDepositTip                   = "finance_deposit_tip"  //	充值提示
	AdminSettingWithdrawTip                  = "finance_withdraw_tip" // 提现提示
)

// AdminSetting 数据库模型属性
type AdminSetting struct {
	Id      int    `json:"id"`       //主键
	AdminId int    `json:"admin_id"` //管理员ID
	GroupId int    `json:"group_id"` //组ID
	Name    string `json:"name"`     //名称
	Type    string `json:"type"`     //类型
	Field   string `json:"field"`    //健名
	Value   string `json:"value"`    //健值
	Data    string `json:"data"`     //数据
}
