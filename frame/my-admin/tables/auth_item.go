package tables

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template/types/form"
)

func GetAuthItemTable(ctx *context.Context) table.Table {

	authItem := table.NewDefaultTable(table.DefaultConfigWithDriver("mysql").SetPrimaryKey("id", db.Bigint))

	info := authItem.GetInfo().HideFilterArea()

	info.AddField("Id", "id", db.Bigint).
		FieldFilterable()
	info.AddField("Name", "name", db.Varchar)
	info.AddField("Type", "type", db.Tinyint)
	info.AddField("Desc", "desc", db.Varchar)
	info.AddField("Rule", "rule", db.Varchar)
	info.AddField("Data", "data", db.Varchar)
	info.AddField("Created_at", "created_at", db.Datetime)
	info.AddField("Updated_at", "updated_at", db.Datetime)
	info.SetTable("auth_item").SetTitle("AuthItem").SetDescription("AuthItem")

	formList := authItem.GetForm()
	formList.AddField("Id", "id", db.Bigint, form.Default)
	formList.AddField("Name", "name", db.Varchar, form.Text)
	formList.AddField("Type", "type", db.Tinyint, form.Number)
	formList.AddField("Desc", "desc", db.Varchar, form.Text)
	formList.AddField("Rule", "rule", db.Varchar, form.Text)
	formList.AddField("Data", "data", db.Varchar, form.Text)

	formList.SetTable("auth_item").SetTitle("AuthItem").SetDescription("AuthItem")

	return authItem
}
