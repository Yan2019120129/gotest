package tables

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template/types/form"
)

func GetInviteTable(ctx *context.Context) table.Table {

	invite := table.NewDefaultTable(table.DefaultConfigWithDriver("mysql").SetPrimaryKey("id", db.Bigint))

	info := invite.GetInfo().HideFilterArea()

	info.AddField("Id", "id", db.Bigint).
		FieldFilterable()
	info.AddField("Admin_id", "admin_id", db.Int)
	info.AddField("User_id", "user_id", db.Int)
	info.AddField("Code", "code", db.Varchar)
	info.AddField("Status", "status", db.Tinyint)
	info.AddField("Data", "data", db.Text)
	info.AddField("Created_at", "created_at", db.Datetime)
	info.AddField("Updated_at", "updated_at", db.Datetime)

	info.SetTable("invite").SetTitle("Invite").SetDescription("Invite")

	formList := invite.GetForm()
	formList.AddField("Id", "id", db.Bigint, form.Default)
	formList.AddField("Admin_id", "admin_id", db.Int, form.Number)
	formList.AddField("User_id", "user_id", db.Int, form.Number)
	formList.AddField("Code", "code", db.Varchar, form.Text)
	formList.AddField("Status", "status", db.Tinyint, form.Number)
	formList.AddField("Data", "data", db.Text, form.RichText)

	formList.SetTable("invite").SetTitle("Invite").SetDescription("Invite")

	return invite
}
