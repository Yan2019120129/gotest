package tables

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template/types/form"
)

func GetLevelTable(ctx *context.Context) table.Table {

	level := table.NewDefaultTable(table.DefaultConfigWithDriver("mysql").SetPrimaryKey("id", db.Bigint))

	info := level.GetInfo().HideFilterArea()

	info.AddField("Id", "id", db.Bigint).
		FieldFilterable()
	info.AddField("Admin_id", "admin_id", db.Int)
	info.AddField("Name", "name", db.Varchar)
	info.AddField("Icon", "icon", db.Varchar)
	info.AddField("Symbol", "symbol", db.Tinyint)
	info.AddField("Money", "money", db.Decimal)
	info.AddField("Days", "days", db.Tinyint)
	info.AddField("Type", "type", db.Smallint)
	info.AddField("Status", "status", db.Smallint)
	info.AddField("Data", "data", db.Text)
	info.AddField("Desc", "desc", db.Text)
	info.AddField("Created_at", "created_at", db.Datetime)
	info.AddField("Updated_at", "updated_at", db.Datetime)

	info.SetTable("level").SetTitle("Level").SetDescription("Level")

	formList := level.GetForm()
	formList.AddField("Id", "id", db.Bigint, form.Default)
	formList.AddField("Admin_id", "admin_id", db.Int, form.Number)
	formList.AddField("Name", "name", db.Varchar, form.Text)
	formList.AddField("Icon", "icon", db.Varchar, form.Text)
	formList.AddField("Symbol", "symbol", db.Tinyint, form.Number)
	formList.AddField("Money", "money", db.Decimal, form.Currency)
	formList.AddField("Days", "days", db.Tinyint, form.Number)
	formList.AddField("Type", "type", db.Smallint, form.Number)
	formList.AddField("Status", "status", db.Smallint, form.Number)
	formList.AddField("Data", "data", db.Text, form.RichText)
	formList.AddField("Desc", "desc", db.Text, form.RichText)

	formList.SetTable("level").SetTitle("Level").SetDescription("Level")

	return level
}
