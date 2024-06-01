package tables

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template/types/form"
)

func GetLangTable(ctx *context.Context) table.Table {

	lang := table.NewDefaultTable(table.DefaultConfigWithDriver("mysql"))

	info := lang.GetInfo()

	info.AddField("Id", "id", db.Int)
	info.AddField("Admin_id", "admin_id", db.Int)
	info.AddField("Name", "name", db.Varchar)
	info.AddField("Alias", "alias", db.Varchar)
	info.AddField("Icon", "icon", db.Varchar)
	info.AddField("Sort", "sort", db.Tinyint)
	info.AddField("Status", "status", db.Tinyint)
	info.AddField("Data", "data", db.Varchar)
	info.AddField("Created_at", "created_at", db.Int)

	info.SetTable("lang").SetTitle("Lang").SetDescription("Lang")

	formList := lang.GetForm()
	formList.AddField("Id", "id", db.Int, form.Default).
		FieldDisableWhenCreate()
	formList.AddField("Admin_id", "admin_id", db.Int, form.Number)
	formList.AddField("Name", "name", db.Varchar, form.Text)
	formList.AddField("Alias", "alias", db.Varchar, form.Text)
	formList.AddField("Icon", "icon", db.Varchar, form.Text)
	formList.AddField("Sort", "sort", db.Tinyint, form.Number)
	formList.AddField("Status", "status", db.Tinyint, form.Number)
	formList.AddField("Data", "data", db.Varchar, form.Text)
	formList.AddField("Created_at", "created_at", db.Int, form.Number).
		FieldHide().FieldNowWhenInsert()

	formList.SetTable("lang").SetTitle("Lang").SetDescription("Lang")

	return lang
}
