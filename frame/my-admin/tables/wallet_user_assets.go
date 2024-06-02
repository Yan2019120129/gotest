package tables

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template/types/form"
)

func GetWalletUserAssetsTable(ctx *context.Context) table.Table {

	walletUserAssets := table.NewDefaultTable(table.DefaultConfigWithDriver("mysql").SetPrimaryKey("id", db.Bigint))

	info := walletUserAssets.GetInfo().HideFilterArea()

	info.AddField("Id", "id", db.Bigint).
		FieldFilterable()
	info.AddField("Created_at", "created_at", db.Datetime)
	info.AddField("Updated_at", "updated_at", db.Datetime)
	info.AddField("Deleted_at", "deleted_at", db.Datetime)
	info.AddField("Admin_id", "admin_id", db.Int)
	info.AddField("User_id", "user_id", db.Int)
	info.AddField("Assets_id", "assets_id", db.Int)
	info.AddField("Money", "money", db.Decimal)
	info.AddField("Status", "status", db.Smallint)
	info.AddField("Data", "data", db.Text)

	info.SetTable("wallet_user_assets").SetTitle("WalletUserAssets").SetDescription("WalletUserAssets")

	formList := walletUserAssets.GetForm()
	formList.AddField("Id", "id", db.Bigint, form.Default)
	formList.AddField("Created_at", "created_at", db.Datetime, form.Datetime)
	formList.AddField("Updated_at", "updated_at", db.Datetime, form.Datetime)
	formList.AddField("Deleted_at", "deleted_at", db.Datetime, form.Datetime)
	formList.AddField("Admin_id", "admin_id", db.Int, form.Number)
	formList.AddField("User_id", "user_id", db.Int, form.Number)
	formList.AddField("Assets_id", "assets_id", db.Int, form.Number)
	formList.AddField("Money", "money", db.Decimal, form.Currency)
	formList.AddField("Status", "status", db.Smallint, form.Number)
	formList.AddField("Data", "data", db.Text, form.RichText)

	formList.SetTable("wallet_user_assets").SetTitle("WalletUserAssets").SetDescription("WalletUserAssets")

	return walletUserAssets
}
