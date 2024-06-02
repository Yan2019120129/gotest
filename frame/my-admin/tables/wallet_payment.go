package tables

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template/types/form"
)

func GetWalletPaymentTable(ctx *context.Context) table.Table {

	walletPayment := table.NewDefaultTable(table.DefaultConfigWithDriver("mysql").SetPrimaryKey("id", db.Bigint))

	info := walletPayment.GetInfo().HideFilterArea()

	info.AddField("Id", "id", db.Bigint).
		FieldFilterable()
	info.AddField("Created_at", "created_at", db.Datetime)
	info.AddField("Updated_at", "updated_at", db.Datetime)
	info.AddField("Deleted_at", "deleted_at", db.Datetime)
	info.AddField("Admin_id", "admin_id", db.Int)
	info.AddField("Assets_id", "assets_id", db.Int)
	info.AddField("Name", "name", db.Varchar)
	info.AddField("Icon", "icon", db.Varchar)
	info.AddField("Type", "type", db.Tinyint)
	info.AddField("Mode", "mode", db.Tinyint)
	info.AddField("Rate", "rate", db.Decimal)
	info.AddField("Is_voucher", "is_voucher", db.Tinyint)
	info.AddField("Level", "level", db.Tinyint)
	info.AddField("Status", "status", db.Smallint)
	info.AddField("Data", "data", db.Text)
	info.AddField("Desc", "desc", db.Text)

	info.SetTable("wallet_payment").SetTitle("WalletPayment").SetDescription("WalletPayment")

	formList := walletPayment.GetForm()
	formList.AddField("Id", "id", db.Bigint, form.Default)
	formList.AddField("Created_at", "created_at", db.Datetime, form.Datetime)
	formList.AddField("Updated_at", "updated_at", db.Datetime, form.Datetime)
	formList.AddField("Deleted_at", "deleted_at", db.Datetime, form.Datetime)
	formList.AddField("Admin_id", "admin_id", db.Int, form.Number)
	formList.AddField("Assets_id", "assets_id", db.Int, form.Number)
	formList.AddField("Name", "name", db.Varchar, form.Text)
	formList.AddField("Icon", "icon", db.Varchar, form.Text)
	formList.AddField("Type", "type", db.Tinyint, form.Number)
	formList.AddField("Mode", "mode", db.Tinyint, form.Number)
	formList.AddField("Rate", "rate", db.Decimal, form.Text)
	formList.AddField("Is_voucher", "is_voucher", db.Tinyint, form.Number)
	formList.AddField("Level", "level", db.Tinyint, form.Number)
	formList.AddField("Status", "status", db.Smallint, form.Number)
	formList.AddField("Data", "data", db.Text, form.RichText)
	formList.AddField("Desc", "desc", db.Text, form.RichText)

	formList.SetTable("wallet_payment").SetTitle("WalletPayment").SetDescription("WalletPayment")

	return walletPayment
}
