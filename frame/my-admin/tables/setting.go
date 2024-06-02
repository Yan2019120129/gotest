package tables

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template/types/form"
)

func GetSettingTable(ctx *context.Context) table.Table {

	setting := table.NewDefaultTable(table.DefaultConfigWithDriver("mysql").SetPrimaryKey("id", db.Bigint))

	info := setting.GetInfo().HideFilterArea()

	info.AddField("Id", "id", db.Bigint).
		FieldFilterable()
	info.AddField("Created_at", "created_at", db.Datetime)
	info.AddField("Updated_at", "updated_at", db.Datetime)
	info.AddField("Deleted_at", "deleted_at", db.Datetime)
	info.AddField("Admin_id", "admin_id", db.Int)
	info.AddField("User_id", "user_id", db.Int)
	info.AddField("Name", "name", db.Varchar)
	info.AddField("Type", "type", db.Tinyint)
	info.AddField("Field", "field", db.Varchar)
	info.AddField("Value", "value", db.Text)
	info.AddField("Data", "data", db.Text)

	info.SetTable("setting").SetTitle("Setting").SetDescription("Setting")

	formList := setting.GetForm()
	formList.AddField("Id", "id", db.Bigint, form.Default)
	formList.AddField("Created_at", "created_at", db.Datetime, form.Datetime)
	formList.AddField("Updated_at", "updated_at", db.Datetime, form.Datetime)
	formList.AddField("Deleted_at", "deleted_at", db.Datetime, form.Datetime)
	formList.AddField("Admin_id", "admin_id", db.Int, form.Number)
	formList.AddField("User_id", "user_id", db.Int, form.Number)
	formList.AddField("Name", "name", db.Varchar, form.Text)
	formList.AddField("Type", "type", db.Tinyint, form.Number)
	formList.AddField("Field", "field", db.Varchar, form.Text)
	formList.AddField("Value", "value", db.Text, form.RichText)
	formList.AddField("Data", "data", db.Text, form.RichText)

	formList.SetTable("setting").SetTitle("Setting").SetDescription("Setting")

	return setting
}
