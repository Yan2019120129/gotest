package tables

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template/types/form"
)

func GetUserLevelTable(ctx *context.Context) table.Table {

	userLevel := table.NewDefaultTable(table.DefaultConfigWithDriver("mysql").SetPrimaryKey("id", db.Bigint))

	info := userLevel.GetInfo().HideFilterArea()

	info.AddField("Id", "id", db.Bigint).
		FieldFilterable()
	info.AddField("Admin_id", "admin_id", db.Int)
	info.AddField("User_id", "user_id", db.Int)
	info.AddField("Name", "name", db.Varchar)
	info.AddField("Icon", "icon", db.Varchar)
	info.AddField("Symbol", "symbol", db.Tinyint)
	info.AddField("Money", "money", db.Decimal)
	info.AddField("Status", "status", db.Tinyint)
	info.AddField("Data", "data", db.Text)
	info.AddField("Expired_at", "expired_at", db.Datetime)
	info.AddField("Created_at", "created_at", db.Datetime)
	info.AddField("Updated_at", "updated_at", db.Datetime)
	info.SetTable("user_level").SetTitle("UserLevel").SetDescription("UserLevel")

	formList := userLevel.GetForm()
	formList.AddField("Id", "id", db.Bigint, form.Default)
	formList.AddField("Created_at", "created_at", db.Datetime, form.Datetime)
	formList.AddField("Updated_at", "updated_at", db.Datetime, form.Datetime)
	formList.AddField("Deleted_at", "deleted_at", db.Datetime, form.Datetime)
	formList.AddField("Admin_id", "admin_id", db.Int, form.Number)
	formList.AddField("User_id", "user_id", db.Int, form.Number)
	formList.AddField("Name", "name", db.Varchar, form.Text)
	formList.AddField("Icon", "icon", db.Varchar, form.Text)
	formList.AddField("Symbol", "symbol", db.Tinyint, form.Number)
	formList.AddField("Money", "money", db.Decimal, form.Currency)
	formList.AddField("Status", "status", db.Tinyint, form.Number)
	formList.AddField("Data", "data", db.Text, form.RichText)
	formList.AddField("Expired_at", "expired_at", db.Datetime, form.Datetime)

	formList.SetTable("user_level").SetTitle("UserLevel").SetDescription("UserLevel")

	return userLevel
}
