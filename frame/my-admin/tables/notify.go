package tables

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template/types/form"
)

func GetNotifyTable(ctx *context.Context) table.Table {

	notify := table.NewDefaultTable(table.DefaultConfigWithDriver("mysql").SetPrimaryKey("id", db.Bigint))

	info := notify.GetInfo().HideFilterArea()

	info.AddField("Id", "id", db.Bigint).
		FieldFilterable()
	info.AddField("Created_at", "created_at", db.Datetime)
	info.AddField("Updated_at", "updated_at", db.Datetime)
	info.AddField("Deleted_at", "deleted_at", db.Datetime)
	info.AddField("Admin_id", "admin_id", db.Int)
	info.AddField("User_id", "user_id", db.Int)
	info.AddField("Mode", "mode", db.Smallint)
	info.AddField("Type", "type", db.Smallint)
	info.AddField("Name", "name", db.Varchar)
	info.AddField("Content", "content", db.Text)
	info.AddField("Status", "status", db.Smallint)
	info.AddField("Data", "data", db.Text)

	info.SetTable("notify").SetTitle("Notify").SetDescription("Notify")

	formList := notify.GetForm()
	formList.AddField("Id", "id", db.Bigint, form.Default)
	formList.AddField("Created_at", "created_at", db.Datetime, form.Datetime)
	formList.AddField("Updated_at", "updated_at", db.Datetime, form.Datetime)
	formList.AddField("Deleted_at", "deleted_at", db.Datetime, form.Datetime)
	formList.AddField("Admin_id", "admin_id", db.Int, form.Number)
	formList.AddField("User_id", "user_id", db.Int, form.Number)
	formList.AddField("Mode", "mode", db.Smallint, form.Number)
	formList.AddField("Type", "type", db.Smallint, form.Number)
	formList.AddField("Name", "name", db.Varchar, form.Text)
	formList.AddField("Content", "content", db.Text, form.RichText)
	formList.AddField("Status", "status", db.Smallint, form.Number)
	formList.AddField("Data", "data", db.Text, form.RichText)

	formList.SetTable("notify").SetTitle("Notify").SetDescription("Notify")

	return notify
}
