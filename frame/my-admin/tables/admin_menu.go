package tables

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template/types/form"
)

func GetAdminMenuTable(ctx *context.Context) table.Table {
	adminMenu := table.NewDefaultTable(table.DefaultConfigWithDriver("mysql").SetPrimaryKey("id", db.Bigint))

	info := adminMenu.GetInfo().HideFilterArea()

	info.AddField("Id", "id", db.Bigint).
		FieldFilterable()
	info.AddField("Created_at", "created_at", db.Datetime)
	info.AddField("Updated_at", "updated_at", db.Datetime)
	info.AddField("Deleted_at", "deleted_at", db.Datetime)
	info.AddField("Parent_id", "parent_id", db.Int)
	info.AddField("Name", "name", db.Varchar)
	info.AddField("Route", "route", db.Varchar)
	info.AddField("Sort", "sort", db.Tinyint)
	info.AddField("Status", "status", db.Tinyint)
	info.AddField("Data", "data", db.Text)

	info.SetTable("admin_menu").SetTitle("AdminMenu").SetDescription("AdminMenu")

	formList := adminMenu.GetForm()
	formList.AddField("Id", "id", db.Bigint, form.Default)
	formList.AddField("Created_at", "created_at", db.Datetime, form.Datetime)
	formList.AddField("Updated_at", "updated_at", db.Datetime, form.Datetime)
	formList.AddField("Deleted_at", "deleted_at", db.Datetime, form.Datetime)
	formList.AddField("Parent_id", "parent_id", db.Int, form.Number)
	formList.AddField("Name", "name", db.Varchar, form.Text)
	formList.AddField("Route", "route", db.Varchar, form.Text)
	formList.AddField("Sort", "sort", db.Tinyint, form.Number)
	formList.AddField("Status", "status", db.Tinyint, form.Number)
	formList.AddField("Data", "data", db.Text, form.RichText)

	formList.SetTable("admin_menu").SetTitle("AdminMenu").SetDescription("AdminMenu")

	return adminMenu
}
