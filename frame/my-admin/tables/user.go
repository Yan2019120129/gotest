package tables

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template/types/form"
)

func GetUserTable(ctx *context.Context) table.Table {

	user := table.NewDefaultTable(table.DefaultConfigWithDriver("mysql").SetPrimaryKey("id", db.Bigint))

	info := user.GetInfo().HideFilterArea()

	info.AddField("Id", "id", db.Bigint).
		FieldFilterable()
	info.AddField("Created_at", "created_at", db.Datetime)
	info.AddField("Updated_at", "updated_at", db.Datetime)
	info.AddField("Deleted_at", "deleted_at", db.Datetime)
	info.AddField("Admin_id", "admin_id", db.Int)
	info.AddField("Parent_id", "parent_id", db.Int)
	info.AddField("Country_id", "country_id", db.Int)
	info.AddField("Username", "username", db.Varchar)
	info.AddField("Nickname", "nickname", db.Varchar)
	info.AddField("Email", "email", db.Varchar)
	info.AddField("Telephone", "telephone", db.Varchar)
	info.AddField("Avatar", "avatar", db.Varchar)
	info.AddField("Score", "score", db.Tinyint)
	info.AddField("Sex", "sex", db.Tinyint)
	info.AddField("Birthday", "birthday", db.Datetime)
	info.AddField("Password", "password", db.Varchar)
	info.AddField("Security_key", "security_key", db.Varchar)
	info.AddField("Money", "money", db.Decimal)
	info.AddField("Type", "type", db.Tinyint)
	info.AddField("Status", "status", db.Smallint)
	info.AddField("Data", "data", db.Text)
	info.AddField("Desc", "desc", db.Text)

	info.SetTable("user").SetTitle("User").SetDescription("User")

	formList := user.GetForm()
	formList.AddField("Id", "id", db.Bigint, form.Default)
	formList.AddField("Created_at", "created_at", db.Datetime, form.Datetime)
	formList.AddField("Updated_at", "updated_at", db.Datetime, form.Datetime)
	formList.AddField("Deleted_at", "deleted_at", db.Datetime, form.Datetime)
	formList.AddField("Admin_id", "admin_id", db.Int, form.Number)
	formList.AddField("Parent_id", "parent_id", db.Int, form.Number)
	formList.AddField("Country_id", "country_id", db.Int, form.Number)
	formList.AddField("Username", "username", db.Varchar, form.Text)
	formList.AddField("Nickname", "nickname", db.Varchar, form.Text)
	formList.AddField("Email", "email", db.Varchar, form.Email)
	formList.AddField("Telephone", "telephone", db.Varchar, form.Text)
	formList.AddField("Avatar", "avatar", db.Varchar, form.Text)
	formList.AddField("Score", "score", db.Tinyint, form.Number)
	formList.AddField("Sex", "sex", db.Tinyint, form.Number)
	formList.AddField("Birthday", "birthday", db.Datetime, form.Datetime)
	formList.AddField("Password", "password", db.Varchar, form.Password)
	formList.AddField("Security_key", "security_key", db.Varchar, form.Text)
	formList.AddField("Money", "money", db.Decimal, form.Currency)
	formList.AddField("Type", "type", db.Tinyint, form.Number)
	formList.AddField("Status", "status", db.Smallint, form.Number)
	formList.AddField("Data", "data", db.Text, form.RichText)
	formList.AddField("Desc", "desc", db.Text, form.RichText)

	formList.SetTable("user").SetTitle("User").SetDescription("User")

	return user
}
