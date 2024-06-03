package tables

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template/types/form"
)

func GetWalletUserConvertTable(ctx *context.Context) table.Table {

	walletUserConvert := table.NewDefaultTable(table.DefaultConfigWithDriver("mysql").SetPrimaryKey("id", db.Bigint))

	info := walletUserConvert.GetInfo().HideFilterArea()

	info.AddField("Id", "id", db.Bigint).
		FieldFilterable()
	info.AddField("Admin_id", "admin_id", db.Int)
	info.AddField("User_id", "user_id", db.Int)
	info.AddField("Type", "type", db.Tinyint)
	info.AddField("Assets_id", "assets_id", db.Int)
	info.AddField("To_assets_id", "to_assets_id", db.Int)
	info.AddField("Money", "money", db.Decimal)
	info.AddField("Nums", "nums", db.Decimal)
	info.AddField("Fee", "fee", db.Decimal)
	info.AddField("Status", "status", db.Smallint)
	info.AddField("Data", "data", db.Text)
	info.AddField("Created_at", "created_at", db.Datetime)
	info.AddField("Updated_at", "updated_at", db.Datetime)

	info.SetTable("wallet_user_convert").SetTitle("WalletUserConvert").SetDescription("WalletUserConvert")

	formList := walletUserConvert.GetForm()
	formList.AddField("Id", "id", db.Bigint, form.Default)
	formList.AddField("Admin_id", "admin_id", db.Int, form.Number)
	formList.AddField("User_id", "user_id", db.Int, form.Number)
	formList.AddField("Type", "type", db.Tinyint, form.Number)
	formList.AddField("Assets_id", "assets_id", db.Int, form.Number)
	formList.AddField("To_assets_id", "to_assets_id", db.Int, form.Number)
	formList.AddField("Money", "money", db.Decimal, form.Currency)
	formList.AddField("Nums", "nums", db.Decimal, form.Text)
	formList.AddField("Fee", "fee", db.Decimal, form.Text)
	formList.AddField("Status", "status", db.Smallint, form.Number)
	formList.AddField("Data", "data", db.Text, form.RichText)

	formList.SetTable("wallet_user_convert").SetTitle("WalletUserConvert").SetDescription("WalletUserConvert")

	return walletUserConvert
}
