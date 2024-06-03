package tables

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template/types/form"
)

func GetWalletAssetsTable(ctx *context.Context) table.Table {

	walletAssets := table.NewDefaultTable(table.DefaultConfigWithDriver("mysql").SetPrimaryKey("id", db.Bigint))

	info := walletAssets.GetInfo().HideFilterArea()

	info.AddField("Id", "id", db.Bigint).
		FieldFilterable()
	info.AddField("Admin_id", "admin_id", db.Int)
	info.AddField("Name", "name", db.Varchar)
	info.AddField("Symbol", "symbol", db.Varchar)
	info.AddField("Icon", "icon", db.Varchar)
	info.AddField("Type", "type", db.Tinyint)
	info.AddField("Rate", "rate", db.Decimal)
	info.AddField("Status", "status", db.Smallint)
	info.AddField("Data", "data", db.Text)
	info.AddField("Created_at", "created_at", db.Datetime)
	info.AddField("Updated_at", "updated_at", db.Datetime)

	info.SetTable("wallet_assets").SetTitle("WalletAssets").SetDescription("WalletAssets")

	formList := walletAssets.GetForm()
	formList.AddField("Id", "id", db.Bigint, form.Default)
	formList.AddField("Admin_id", "admin_id", db.Int, form.Number)
	formList.AddField("Name", "name", db.Varchar, form.Text)
	formList.AddField("Symbol", "symbol", db.Varchar, form.Text)
	formList.AddField("Icon", "icon", db.Varchar, form.Text)
	formList.AddField("Type", "type", db.Tinyint, form.Number)
	formList.AddField("Rate", "rate", db.Decimal, form.Text)
	formList.AddField("Status", "status", db.Smallint, form.Number)
	formList.AddField("Data", "data", db.Text, form.RichText)

	formList.SetTable("wallet_assets").SetTitle("WalletAssets").SetDescription("WalletAssets")

	return walletAssets
}
