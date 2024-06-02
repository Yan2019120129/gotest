package tables

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template/types/form"
)

func GetWalletUserTransferTable(ctx *context.Context) table.Table {

	walletUserTransfer := table.NewDefaultTable(table.DefaultConfigWithDriver("mysql").SetPrimaryKey("id", db.Bigint))

	info := walletUserTransfer.GetInfo().HideFilterArea()

	info.AddField("Id", "id", db.Bigint).
		FieldFilterable()
	info.AddField("Created_at", "created_at", db.Datetime)
	info.AddField("Updated_at", "updated_at", db.Datetime)
	info.AddField("Deleted_at", "deleted_at", db.Datetime)
	info.AddField("Admin_id", "admin_id", db.Int)
	info.AddField("Sender_id", "sender_id", db.Int)
	info.AddField("Receiver_id", "receiver_id", db.Int)
	info.AddField("Type", "type", db.Tinyint)
	info.AddField("Assets_id", "assets_id", db.Int)
	info.AddField("Money", "money", db.Decimal)
	info.AddField("Fee", "fee", db.Decimal)
	info.AddField("Status", "status", db.Smallint)
	info.AddField("Data", "data", db.Text)

	info.SetTable("wallet_user_transfer").SetTitle("WalletUserTransfer").SetDescription("WalletUserTransfer")

	formList := walletUserTransfer.GetForm()
	formList.AddField("Id", "id", db.Bigint, form.Default)
	formList.AddField("Created_at", "created_at", db.Datetime, form.Datetime)
	formList.AddField("Updated_at", "updated_at", db.Datetime, form.Datetime)
	formList.AddField("Deleted_at", "deleted_at", db.Datetime, form.Datetime)
	formList.AddField("Admin_id", "admin_id", db.Int, form.Number)
	formList.AddField("Sender_id", "sender_id", db.Int, form.Number)
	formList.AddField("Receiver_id", "receiver_id", db.Int, form.Number)
	formList.AddField("Type", "type", db.Tinyint, form.Number)
	formList.AddField("Assets_id", "assets_id", db.Int, form.Number)
	formList.AddField("Money", "money", db.Decimal, form.Currency)
	formList.AddField("Fee", "fee", db.Decimal, form.Text)
	formList.AddField("Status", "status", db.Smallint, form.Number)
	formList.AddField("Data", "data", db.Text, form.RichText)

	formList.SetTable("wallet_user_transfer").SetTitle("WalletUserTransfer").SetDescription("WalletUserTransfer")

	return walletUserTransfer
}
