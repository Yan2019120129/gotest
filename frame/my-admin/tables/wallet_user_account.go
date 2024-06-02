package tables

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template/types/form"
)

func GetWalletUserAccountTable(ctx *context.Context) table.Table {

	walletUserAccount := table.NewDefaultTable(table.DefaultConfigWithDriver("mysql").SetPrimaryKey("id", db.Bigint))

	info := walletUserAccount.GetInfo().HideFilterArea()

	info.AddField("Id", "id", db.Bigint).
		FieldFilterable()
	info.AddField("Admin_id", "admin_id", db.Int)
	info.AddField("User_id", "user_id", db.Int)
	info.AddField("Payment_id", "payment_id", db.Int)
	info.AddField("Name", "name", db.Varchar)
	info.AddField("Real_name", "real_name", db.Varchar)
	info.AddField("Number", "number", db.Varchar)
	info.AddField("Code", "code", db.Varchar)
	info.AddField("Remark", "remark", db.Varchar)
	info.AddField("Status", "status", db.Smallint)
	info.AddField("Data", "data", db.Text)
	info.AddField("Created_at", "created_at", db.Datetime)
	info.AddField("Updated_at", "updated_at", db.Datetime)

	info.SetTable("wallet_user_account").SetTitle("WalletUserAccount").SetDescription("WalletUserAccount")

	formList := walletUserAccount.GetForm()
	formList.AddField("Id", "id", db.Bigint, form.Default)
	formList.AddField("Admin_id", "admin_id", db.Int, form.Number)
	formList.AddField("User_id", "user_id", db.Int, form.Number)
	formList.AddField("Payment_id", "payment_id", db.Int, form.Number)
	formList.AddField("Name", "name", db.Varchar, form.Text)
	formList.AddField("Real_name", "real_name", db.Varchar, form.Text)
	formList.AddField("Number", "number", db.Varchar, form.Text)
	formList.AddField("Code", "code", db.Varchar, form.Text)
	formList.AddField("Remark", "remark", db.Varchar, form.Text)
	formList.AddField("Status", "status", db.Smallint, form.Number)
	formList.AddField("Data", "data", db.Text, form.RichText)

	formList.SetTable("wallet_user_account").SetTitle("WalletUserAccount").SetDescription("WalletUserAccount")

	return walletUserAccount
}
