package tables

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template/types/form"
)

func GetAccessTable(ctx *context.Context) table.Table {

	access := table.NewDefaultTable(table.DefaultConfigWithDriver("mysql").SetPrimaryKey("id", db.Bigint))

	info := access.GetInfo().HideFilterArea()

	info.AddField("Id", "id", db.Bigint).
		FieldFilterable()
	info.AddField("Admin_id", "admin_id", db.Int)
	info.AddField("User_id", "user_id", db.Int)
	info.AddField("Source_id", "source_id", db.Int)
	info.AddField("Name", "name", db.Varchar)
	info.AddField("Type", "type", db.Tinyint)
	info.AddField("Ip", "ip", db.Int)
	info.AddField("Route", "route", db.Varchar)
	info.AddField("Headers", "headers", db.Text)
	info.AddField("Data", "data", db.Text)
	info.AddField("Created_at", "created_at", db.Datetime)
	info.AddField("Updated_at", "updated_at", db.Datetime)
	info.SetTable("access").SetTitle("Access").SetDescription("Access")

	formList := access.GetForm()
	formList.AddField("Id", "id", db.Bigint, form.Default)

	formList.AddField("Admin_id", "admin_id", db.Int, form.Number)
	formList.AddField("User_id", "user_id", db.Int, form.Number)
	formList.AddField("Source_id", "source_id", db.Int, form.Number)
	formList.AddField("Name", "name", db.Varchar, form.Text)
	formList.AddField("Type", "type", db.Tinyint, form.Number)
	formList.AddField("Ip", "ip", db.Int, form.Ip)
	formList.AddField("Route", "route", db.Varchar, form.Text)
	formList.AddField("Headers", "headers", db.Text, form.RichText)
	formList.AddField("Data", "data", db.Text, form.RichText)
	formList.SetTable("access").SetTitle("Access").SetDescription("Access")

	return access
}
