package tables

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template/types/form"
)

func GetCategoryTable(ctx *context.Context) table.Table {

	category := table.NewDefaultTable(table.DefaultConfigWithDriver("mysql").SetPrimaryKey("id", db.Bigint))

	info := category.GetInfo().HideFilterArea()

	info.AddField("Id", "id", db.Bigint).
		FieldFilterable()
	info.AddField("Parent_id", "parent_id", db.Int)
	info.AddField("Admin_id", "admin_id", db.Int)
	info.AddField("Type", "type", db.Tinyint)
	info.AddField("Name", "name", db.Varchar)
	info.AddField("Icon", "icon", db.Varchar)
	info.AddField("Sort", "sort", db.Tinyint)
	info.AddField("Status", "status", db.Tinyint)
	info.AddField("Data", "data", db.Text)
	info.AddField("Created_at", "created_at", db.Datetime)
	info.AddField("Updated_at", "updated_at", db.Datetime)
	info.SetTable("category").SetTitle("Category").SetDescription("Category")

	formList := category.GetForm()
	formList.AddField("Id", "id", db.Bigint, form.Default)
	formList.AddField("Parent_id", "parent_id", db.Int, form.Number)
	formList.AddField("Admin_id", "admin_id", db.Int, form.Number)
	formList.AddField("Type", "type", db.Tinyint, form.Number)
	formList.AddField("Name", "name", db.Varchar, form.Text)
	formList.AddField("Icon", "icon", db.Varchar, form.Text)
	formList.AddField("Sort", "sort", db.Tinyint, form.Number)
	formList.AddField("Status", "status", db.Tinyint, form.Number)
	formList.AddField("Data", "data", db.Text, form.RichText)

	formList.SetTable("category").SetTitle("Category").SetDescription("Category")

	return category
}
