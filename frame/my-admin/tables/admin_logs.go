package tables

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template/types/form"
)

func GetAdminLogsTable(ctx *context.Context) table.Table {

	adminLogs := table.NewDefaultTable(table.DefaultConfigWithDriver("mysql").SetPrimaryKey("id", db.Bigint))

	info := adminLogs.GetInfo().HideFilterArea()

	info.AddField("Id", "id", db.Bigint).
		FieldFilterable()
	info.AddField("Admin_id", "admin_id", db.Int)
	info.AddField("Ip", "ip", db.Int)
	info.AddField("Headers", "headers", db.Text)
	info.AddField("Name", "name", db.Varchar)
	info.AddField("Route", "route", db.Varchar)
	info.AddField("Body", "body", db.Text)
	info.AddField("Data", "data", db.Text)
	info.AddField("Created_at", "created_at", db.Datetime)
	info.AddField("Updated_at", "updated_at", db.Datetime)
	info.SetTable("admin_logs").SetTitle("AdminLogs").SetDescription("AdminLogs")

	formList := adminLogs.GetForm()
	formList.AddField("Id", "id", db.Bigint, form.Default)
	formList.AddField("Admin_id", "admin_id", db.Int, form.Number)
	formList.AddField("Ip", "ip", db.Int, form.Ip)
	formList.AddField("Headers", "headers", db.Text, form.RichText)
	formList.AddField("Name", "name", db.Varchar, form.Text)
	formList.AddField("Route", "route", db.Varchar, form.Text)
	formList.AddField("Body", "body", db.Text, form.RichText)
	formList.AddField("Data", "data", db.Text, form.RichText)
	formList.SetTable("admin_logs").SetTitle("AdminLogs").SetDescription("AdminLogs")

	return adminLogs
}
