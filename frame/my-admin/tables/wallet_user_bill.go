package tables

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template/types/form"
)

func GetWalletUserBillTable(ctx *context.Context) table.Table {

	walletUserBill := table.NewDefaultTable(table.DefaultConfigWithDriver("mysql").SetPrimaryKey("id", db.Bigint))

	info := walletUserBill.GetInfo().HideFilterArea()

	info.AddField("Id", "id", db.Bigint).
		FieldFilterable()
	info.AddField("Created_at", "created_at", db.Datetime)
	info.AddField("Updated_at", "updated_at", db.Datetime)
	info.AddField("Deleted_at", "deleted_at", db.Datetime)
	info.AddField("Admin_id", "admin_id", db.Int)
	info.AddField("User_id", "user_id", db.Int)
	info.AddField("Assets_id", "assets_id", db.Int)
	info.AddField("Source_id", "source_id", db.Int)
	info.AddField("Type", "type", db.Smallint)
	info.AddField("Name", "name", db.Varchar)
	info.AddField("Money", "money", db.Decimal)
	info.AddField("Balance", "balance", db.Decimal)
	info.AddField("Data", "data", db.Text)

	info.SetTable("wallet_user_bill").SetTitle("WalletUserBill").SetDescription("WalletUserBill")

	formList := walletUserBill.GetForm()
	formList.AddField("Id", "id", db.Bigint, form.Default)
	formList.AddField("Created_at", "created_at", db.Datetime, form.Datetime)
	formList.AddField("Updated_at", "updated_at", db.Datetime, form.Datetime)
	formList.AddField("Deleted_at", "deleted_at", db.Datetime, form.Datetime)
	formList.AddField("Admin_id", "admin_id", db.Int, form.Number)
	formList.AddField("User_id", "user_id", db.Int, form.Number)
	formList.AddField("Assets_id", "assets_id", db.Int, form.Number)
	formList.AddField("Source_id", "source_id", db.Int, form.Number)
	formList.AddField("Type", "type", db.Smallint, form.Number)
	formList.AddField("Name", "name", db.Varchar, form.Text)
	formList.AddField("Money", "money", db.Decimal, form.Currency)
	formList.AddField("Balance", "balance", db.Decimal, form.Text)
	formList.AddField("Data", "data", db.Text, form.RichText)

	formList.SetTable("wallet_user_bill").SetTitle("WalletUserBill").SetDescription("WalletUserBill")

	return walletUserBill
}
