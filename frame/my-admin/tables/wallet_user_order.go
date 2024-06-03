package tables

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template/types/form"
)

func GetWalletUserOrderTable(ctx *context.Context) table.Table {

	walletUserOrder := table.NewDefaultTable(table.DefaultConfigWithDriver("mysql").SetPrimaryKey("id", db.Bigint))

	info := walletUserOrder.GetInfo().HideFilterArea()

	info.AddField("Id", "id", db.Bigint).
		FieldFilterable()
	info.AddField("Admin_id", "admin_id", db.Int)
	info.AddField("User_id", "user_id", db.Int)
	info.AddField("Assets_id", "assets_id", db.Bigint)
	info.AddField("Source_id", "source_id", db.Int)
	info.AddField("Type", "type", db.Tinyint)
	info.AddField("Order_sn", "order_sn", db.Varchar)
	info.AddField("Money", "money", db.Decimal)
	info.AddField("Fee", "fee", db.Decimal)
	info.AddField("Voucher", "voucher", db.Varchar)
	info.AddField("Status", "status", db.Smallint)
	info.AddField("Data", "data", db.Text)
	info.AddField("Created_at", "created_at", db.Datetime)
	info.AddField("Updated_at", "updated_at", db.Datetime)

	info.SetTable("wallet_user_order").SetTitle("WalletUserOrder").SetDescription("WalletUserOrder")

	formList := walletUserOrder.GetForm()
	formList.AddField("Id", "id", db.Bigint, form.Default)
	formList.AddField("Admin_id", "admin_id", db.Int, form.Number)
	formList.AddField("User_id", "user_id", db.Int, form.Number)
	formList.AddField("Assets_id", "assets_id", db.Bigint, form.Number)
	formList.AddField("Source_id", "source_id", db.Int, form.Number)
	formList.AddField("Type", "type", db.Tinyint, form.Number)
	formList.AddField("Order_sn", "order_sn", db.Varchar, form.Text)
	formList.AddField("Money", "money", db.Decimal, form.Currency)
	formList.AddField("Fee", "fee", db.Decimal, form.Text)
	formList.AddField("Voucher", "voucher", db.Varchar, form.Text)
	formList.AddField("Status", "status", db.Smallint, form.Number)
	formList.AddField("Data", "data", db.Text, form.RichText)

	formList.SetTable("wallet_user_order").SetTitle("WalletUserOrder").SetDescription("WalletUserOrder")

	return walletUserOrder
}
