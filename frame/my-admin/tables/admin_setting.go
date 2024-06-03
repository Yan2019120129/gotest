package tables

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template/types/form"
)

func GetAdminSettingTable(ctx *context.Context) table.Table {

	adminSetting := table.NewDefaultTable(table.DefaultConfigWithDriver("mysql").SetPrimaryKey("id", db.Bigint))

	info := adminSetting.GetInfo().HideFilterArea()

	info.AddField("Id", "id", db.Bigint).
		FieldFilterable()
	info.AddField("Admin_id", "admin_id", db.Int)
	info.AddField("Group_id", "group_id", db.Int)
	info.AddField("Name", "name", db.Varchar)
	info.AddField("Type", "type", db.Tinyint)
	info.AddField("Field", "field", db.Varchar)
	info.AddField("Value", "value", db.Text)
	info.AddField("Data", "data", db.Text)
	info.AddField("Created_at", "created_at", db.Datetime)
	info.AddField("Updated_at", "updated_at", db.Datetime)
	info.SetTable("admin_setting").SetTitle("AdminSetting").SetDescription("AdminSetting")

	formList := adminSetting.GetForm()
	formList.AddField("Id", "id", db.Bigint, form.Default)
	formList.AddField("Admin_id", "admin_id", db.Int, form.Number)
	formList.AddField("Group_id", "group_id", db.Int, form.Number)
	formList.AddField("Name", "name", db.Varchar, form.Text)
	formList.AddField("Type", "type", db.Tinyint, form.Number)
	formList.AddField("Field", "field", db.Varchar, form.Text)
	formList.AddField("Value", "value", db.Text, form.RichText)
	formList.AddField("Data", "data", db.Text, form.RichText)
	formList.SetTable("admin_setting").SetTitle("AdminSetting").SetDescription("AdminSetting")

	return adminSetting
}
