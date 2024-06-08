package language

import "github.com/GoAdminGroup/go-admin/modules/language"

// InitCN 初始化中文翻译
func InitCN() {
	lang := map[string]string{
		// 公共
		"level": "等级",

		"system_manage": "系统管理",
		"translate":     "翻译",
		"notify":        "消息通知",
		"lang":          "语言",
		"country":       "国家",
		"article":       "文章",
		"admin_setting": "管理设置",

		"users_manage": "用户管理",
		"setting":      "设置",
		"access":       "访问记录",
		"invite":       "邀请",
		"auth":         "验证",

		"wallet":               "钱包",
		"wallet_assets":        "资产",
		"wallet_payment":       "支付",
		"wallet_user_account":  "账户",
		"wallet_user_assets":   "用户资产",
		"wallet_user_bill":     "支付账单",
		"wallet_user_convert":  "转账",
		"wallet_user_order":    "支付订单",
		"wallet_user_transfer": "划转",
	}
	language.AppendTo(language.CN, lang)
}
