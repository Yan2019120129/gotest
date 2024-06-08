package language

import "github.com/GoAdminGroup/go-admin/modules/language"

// InitEN 初始化英文翻译
func InitEN() {
	lang := map[string]string{
		// 公共
		"level": "Level",

		"system_manage": "System manage",
		"translate":     "Translate",
		"notify":        "Notify",
		"lang":          "Lang",
		"country":       "Country",
		"article":       "Article",

		"users_manage": "Users manage",
		"access":       "Access",
		"invite":       "Invite",
		"auth":         "Auth",

		"wallet":               "Wallet",
		"wallet_assets":        "Assets",
		"wallet_payment":       "Payment",
		"wallet_user_account":  "Account",
		"wallet_user_assets":   "Assets",
		"wallet_user_bill":     "Bill",
		"wallet_user_convert":  "Convert",
		"wallet_user_order":    "Order",
		"wallet_user_transfer": "Transfer",
	}
	language.AppendTo(language.EN, lang)
}
