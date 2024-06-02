package tables

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template/types/form"
)

func GetMenuTable(ctx *context.Context) table.Table {

	menu := table.NewDefaultTable(table.DefaultConfigWithDriver("mysql").SetPrimaryKey("id", db.Bigint))

	info := menu.GetInfo().HideFilterArea()

	info.AddField("Id", "id", db.Bigint).
		FieldFilterable()
	info.AddField("Admin_id", "admin_id", db.Int)
	info.AddField("Parent_id", "parent_id", db.Int)
	info.AddField("Name", "name", db.Varchar)
	info.AddField("Route", "route", db.Varchar)
	info.AddField("Sort", "sort", db.Tinyint)
	info.AddField("Icon", "icon", db.Varchar)
	info.AddField("Active_icon", "active_icon", db.Varchar)
	info.AddField("Is_desktop", "is_desktop", db.Tinyint)
	info.AddField("Is_mobile", "is_mobile", db.Tinyint)
	info.AddField("Type", "type", db.Tinyint)
	info.AddField("Status", "status", db.Tinyint)
	info.AddField("Data", "data", db.Text)
	info.AddField("Created_at", "created_at", db.Datetime)
	info.AddField("Updated_at", "updated_at", db.Datetime)

	info.SetTable("menu").SetTitle("Menu").SetDescription("Menu")

	formList := menu.GetForm()
	formList.AddField("Id", "id", db.Bigint, form.Default)
	formList.AddField("Admin_id", "admin_id", db.Int, form.Number)
	formList.AddField("Parent_id", "parent_id", db.Int, form.Number)
	formList.AddField("Name", "name", db.Varchar, form.Text)
	formList.AddField("Route", "route", db.Varchar, form.Text)
	formList.AddField("Sort", "sort", db.Tinyint, form.Number)
	formList.AddField("Icon", "icon", db.Varchar, form.Text)
	formList.AddField("Active_icon", "active_icon", db.Varchar, form.Text)
	formList.AddField("Is_desktop", "is_desktop", db.Tinyint, form.Number)
	formList.AddField("Is_mobile", "is_mobile", db.Tinyint, form.Number)
	formList.AddField("Type", "type", db.Tinyint, form.Number)
	formList.AddField("Status", "status", db.Tinyint, form.Number)
	formList.AddField("Data", "data", db.Text, form.RichText)

	formList.SetTable("menu").SetTitle("Menu").SetDescription("Menu")

	return menu
}
