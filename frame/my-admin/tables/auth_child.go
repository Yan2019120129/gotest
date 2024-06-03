package tables

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template/types/form"
)

func GetAuthChildTable(ctx *context.Context) table.Table {

	authChild := table.NewDefaultTable(table.DefaultConfigWithDriver("mysql").SetPrimaryKey("id", db.Bigint))

	info := authChild.GetInfo().HideFilterArea()

	info.AddField("Id", "id", db.Bigint).
		FieldFilterable()
	info.AddField("Parent", "parent", db.Varchar)
	info.AddField("Child", "child", db.Varchar)
	info.AddField("Type", "type", db.Tinyint)
	info.AddField("Created_at", "created_at", db.Datetime)
	info.AddField("Updated_at", "updated_at", db.Datetime)
	info.SetTable("auth_child").SetTitle("AuthChild").SetDescription("AuthChild")

	formList := authChild.GetForm()
	formList.AddField("Id", "id", db.Bigint, form.Default)
	formList.AddField("Parent", "parent", db.Varchar, form.Text)
	formList.AddField("Child", "child", db.Varchar, form.Text)
	formList.AddField("Type", "type", db.Tinyint, form.Number)

	formList.SetTable("auth_child").SetTitle("AuthChild").SetDescription("AuthChild")

	return authChild
}
