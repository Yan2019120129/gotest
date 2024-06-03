package tables

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template/types/form"
)

func GetAuthAssignmentTable(ctx *context.Context) table.Table {

	authAssignment := table.NewDefaultTable(table.DefaultConfigWithDriver("mysql").SetPrimaryKey("id", db.Bigint))

	info := authAssignment.GetInfo().HideFilterArea()

	info.AddField("Id", "id", db.Bigint).
		FieldFilterable()
	info.AddField("Admin_id", "admin_id", db.Int)
	info.AddField("Name", "name", db.Varchar)
	info.AddField("Created_at", "created_at", db.Datetime)
	info.AddField("Updated_at", "updated_at", db.Datetime)
	info.AddField("Deleted_at", "deleted_at", db.Datetime)
	info.SetTable("auth_assignment").SetTitle("AuthAssignment").SetDescription("AuthAssignment")

	formList := authAssignment.GetForm()
	formList.AddField("Id", "id", db.Bigint, form.Default)
	formList.AddField("Admin_id", "admin_id", db.Int, form.Number)
	formList.AddField("Name", "name", db.Varchar, form.Text)

	formList.SetTable("auth_assignment").SetTitle("AuthAssignment").SetDescription("AuthAssignment")

	return authAssignment
}
