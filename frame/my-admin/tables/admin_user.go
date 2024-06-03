package tables

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template/types/form"
)

func GetAdminUserTable(ctx *context.Context) table.Table {

	adminUser := table.NewDefaultTable(table.DefaultConfigWithDriver("mysql").SetPrimaryKey("id", db.Bigint))

	info := adminUser.GetInfo().HideFilterArea()

	info.AddField("Id", "id", db.Bigint).
		FieldFilterable()

	info.AddField("Parent_id", "parent_id", db.Int)
	info.AddField("Username", "username", db.Varchar)
	info.AddField("Nickname", "nickname", db.Varchar)
	info.AddField("Email", "email", db.Varchar)
	info.AddField("Avatar", "avatar", db.Varchar)
	info.AddField("Password", "password", db.Varchar)
	info.AddField("Security_key", "security_key", db.Varchar)
	info.AddField("Money", "money", db.Decimal)
	info.AddField("Status", "status", db.Smallint)
	info.AddField("Data", "data", db.Text)
	info.AddField("Domains", "domains", db.Varchar)
	info.AddField("Seat_link", "seat_link", db.Varchar)
	info.AddField("Online", "online", db.Varchar)
	info.AddField("Expired_at", "expired_at", db.Datetime)
	info.AddField("Created_at", "created_at", db.Datetime)
	info.AddField("Updated_at", "updated_at", db.Datetime)
	info.SetTable("admin_user").SetTitle("AdminUser").SetDescription("AdminUser")

	formList := adminUser.GetForm()
	formList.AddField("Id", "id", db.Bigint, form.Default)
	formList.AddField("Parent_id", "parent_id", db.Int, form.Number)
	formList.AddField("Username", "username", db.Varchar, form.Text)
	formList.AddField("Nickname", "nickname", db.Varchar, form.Text)
	formList.AddField("Email", "email", db.Varchar, form.Email)
	formList.AddField("Avatar", "avatar", db.Varchar, form.Text)
	formList.AddField("Password", "password", db.Varchar, form.Password)
	formList.AddField("Security_key", "security_key", db.Varchar, form.Text)
	formList.AddField("Money", "money", db.Decimal, form.Currency)
	formList.AddField("Status", "status", db.Smallint, form.Number)
	formList.AddField("Data", "data", db.Text, form.RichText)
	formList.AddField("Domains", "domains", db.Varchar, form.Text)
	formList.AddField("Seat_link", "seat_link", db.Varchar, form.Text)
	formList.AddField("Online", "online", db.Varchar, form.Text)
	formList.AddField("Expired_at", "expired_at", db.Datetime, form.Datetime)

	formList.SetTable("admin_user").SetTitle("AdminUser").SetDescription("AdminUser")

	return adminUser
}
