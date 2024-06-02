package tables

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template/types/form"
)

func GetArticleTable(ctx *context.Context) table.Table {

	article := table.NewDefaultTable(table.DefaultConfigWithDriver("mysql").SetPrimaryKey("id", db.Bigint))

	info := article.GetInfo().HideFilterArea()

	info.AddField("Id", "id", db.Bigint).
		FieldFilterable()
	info.AddField("Created_at", "created_at", db.Datetime)
	info.AddField("Updated_at", "updated_at", db.Datetime)
	info.AddField("Deleted_at", "deleted_at", db.Datetime)
	info.AddField("Admin_id", "admin_id", db.Int)
	info.AddField("Image", "image", db.Varchar)
	info.AddField("Name", "name", db.Varchar)
	info.AddField("Content", "content", db.Varchar)
	info.AddField("Link", "link", db.Varchar)
	info.AddField("Type", "type", db.Smallint)
	info.AddField("Status", "status", db.Smallint)
	info.AddField("Data", "data", db.Text)

	info.SetTable("article").SetTitle("Article").SetDescription("Article")

	formList := article.GetForm()
	formList.AddField("Id", "id", db.Bigint, form.Default)
	formList.AddField("Created_at", "created_at", db.Datetime, form.Datetime)
	formList.AddField("Updated_at", "updated_at", db.Datetime, form.Datetime)
	formList.AddField("Deleted_at", "deleted_at", db.Datetime, form.Datetime)
	formList.AddField("Admin_id", "admin_id", db.Int, form.Number)
	formList.AddField("Image", "image", db.Varchar, form.Text)
	formList.AddField("Name", "name", db.Varchar, form.Text)
	formList.AddField("Content", "content", db.Varchar, form.Text)
	formList.AddField("Link", "link", db.Varchar, form.Text)
	formList.AddField("Type", "type", db.Smallint, form.Number)
	formList.AddField("Status", "status", db.Smallint, form.Number)
	formList.AddField("Data", "data", db.Text, form.RichText)

	formList.SetTable("article").SetTitle("Article").SetDescription("Article")

	return article
}
