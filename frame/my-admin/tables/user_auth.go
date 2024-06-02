package tables

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template/types/form"
)

func GetUserAuthTable(ctx *context.Context) table.Table {

	userAuth := table.NewDefaultTable(table.DefaultConfigWithDriver("mysql").SetPrimaryKey("id", db.Bigint))

	info := userAuth.GetInfo().HideFilterArea()

	info.AddField("Id", "id", db.Bigint).
		FieldFilterable()
	info.AddField("Admin_id", "admin_id", db.Int)
	info.AddField("User_id", "user_id", db.Int)
	info.AddField("Real_name", "real_name", db.Varchar)
	info.AddField("Number", "number", db.Varchar)
	info.AddField("Photo1", "photo1", db.Varchar)
	info.AddField("Photo2", "photo2", db.Varchar)
	info.AddField("Photo3", "photo3", db.Varchar)
	info.AddField("Address", "address", db.Varchar)
	info.AddField("Type", "type", db.Tinyint)
	info.AddField("Status", "status", db.Tinyint)
	info.AddField("Data", "data", db.Varchar)
	info.AddField("Created_at", "created_at", db.Datetime)
	info.AddField("Updated_at", "updated_at", db.Datetime)

	info.SetTable("user_auth").SetTitle("UserAuth").SetDescription("UserAuth")

	formList := userAuth.GetForm()
	formList.AddField("Id", "id", db.Bigint, form.Default)
	formList.AddField("Admin_id", "admin_id", db.Int, form.Number)
	formList.AddField("User_id", "user_id", db.Int, form.Number)
	formList.AddField("Real_name", "real_name", db.Varchar, form.Text)
	formList.AddField("Number", "number", db.Varchar, form.Text)
	formList.AddField("Photo1", "photo1", db.Varchar, form.Text)
	formList.AddField("Photo2", "photo2", db.Varchar, form.Text)
	formList.AddField("Photo3", "photo3", db.Varchar, form.Text)
	formList.AddField("Address", "address", db.Varchar, form.Text)
	formList.AddField("Type", "type", db.Tinyint, form.Number)
	formList.AddField("Status", "status", db.Tinyint, form.Number)
	formList.AddField("Data", "data", db.Varchar, form.Text)

	formList.SetTable("user_auth").SetTitle("UserAuth").SetDescription("UserAuth")

	return userAuth
}
