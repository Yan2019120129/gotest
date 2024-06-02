package tables

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template/types/form"
)

func GetProductTable(ctx *context.Context) table.Table {

	product := table.NewDefaultTable(table.DefaultConfigWithDriver("mysql").SetPrimaryKey("id", db.Bigint))

	info := product.GetInfo().HideFilterArea()

	info.AddField("Id", "id", db.Bigint).
		FieldFilterable()
	info.AddField("Admin_id", "admin_id", db.Int)
	info.AddField("Category_id", "category_id", db.Int)
	info.AddField("Assets_id", "assets_id", db.Int)
	info.AddField("Name", "name", db.Varchar)
	info.AddField("Images", "images", db.Varchar)
	info.AddField("Money", "money", db.Decimal)
	info.AddField("Type", "type", db.Tinyint)
	info.AddField("Sort", "sort", db.Tinyint)
	info.AddField("Status", "status", db.Tinyint)
	info.AddField("Data", "data", db.Text)
	info.AddField("Desc", "desc", db.Text)
	info.AddField("Created_at", "created_at", db.Datetime)
	info.AddField("Updated_at", "updated_at", db.Datetime)

	info.SetTable("product").SetTitle("Product").SetDescription("Product")

	formList := product.GetForm()
	formList.AddField("Id", "id", db.Bigint, form.Default)
	formList.AddField("Admin_id", "admin_id", db.Int, form.Number)
	formList.AddField("Category_id", "category_id", db.Int, form.Number)
	formList.AddField("Assets_id", "assets_id", db.Int, form.Number)
	formList.AddField("Name", "name", db.Varchar, form.Text)
	formList.AddField("Images", "images", db.Varchar, form.Text)
	formList.AddField("Money", "money", db.Decimal, form.Currency)
	formList.AddField("Type", "type", db.Tinyint, form.Number)
	formList.AddField("Sort", "sort", db.Tinyint, form.Number)
	formList.AddField("Status", "status", db.Tinyint, form.Number)
	formList.AddField("Data", "data", db.Text, form.RichText)
	formList.AddField("Desc", "desc", db.Text, form.RichText)

	formList.SetTable("product").SetTitle("Product").SetDescription("Product")

	return product
}
