package tables

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template/types"
	"github.com/GoAdminGroup/go-admin/template/types/form"
)

func GetUsersTable(ctx *context.Context) table.Table {
	//table.DefaultConfigWithDriver("mysql").SetPrimaryKey("id", db.Int)
	//users := table.DefaultConfigWithDriver("mysql").SetPrimaryKey("id", db.Int)
	users := table.NewDefaultTable(table.Config{
		Driver: "mysql",
		//DriverMode: "",
		Connection: table.DefaultConnectionName, // 默认链接名
		CanAdd:     true,                        // 是否可以新增
		Editable:   true,                        // 是否可以编辑
		Deletable:  true,                        // 是否可以删除
		Exportable: true,                        // 是否可以导出为excel
		PrimaryKey: table.PrimaryKey{ // 自定义主键，默认为id
			Type: db.Int,                      // 字段类型
			Name: table.DefaultPrimaryKeyName, // 字段名
		},
		//SourceURL:      "",
		//GetDataFun:     nil,
		//OnlyInfo:       false,
		//OnlyNewForm:    false,
		//OnlyUpdateForm: false,
		//OnlyDetail:     false,
	})

	info := users.GetInfo().HideFilterArea()

	// FieldFilterable 设置主键id为可排序
	info.AddField("Id", "id", db.Int).FieldFilterable()
	info.AddField("Name", "name", db.Varchar)
	info.AddField("Gender", "gender", db.Tinyint).FieldDisplay(func(v types.FieldModel) interface{} {
		if v.Value == "1" {
			return "men"
		}
		if v.Value == "2" {
			return "women"
		}
		return "unknown"
	})
	info.AddField("City", "city", db.Varchar)
	info.AddField("Ip", "ip", db.Varchar)
	info.AddField("Phone", "phone", db.Varchar)
	info.AddField("Created_at", "created_at", db.Timestamp)
	info.AddField("Updated_at", "updated_at", db.Timestamp)

	info.SetTable("users").SetTitle("Users").SetDescription("Users")

	formList := users.GetForm()
	formList.AddField("Id", "id", db.Int, form.Default)
	formList.AddField("Name", "name", db.Varchar, form.Text)
	formList.AddField("Gender", "gender", db.Tinyint, form.Number)
	formList.AddField("City", "city", db.Varchar, form.Text)
	formList.AddField("Ip", "ip", db.Varchar, form.Ip)
	formList.AddField("Phone", "phone", db.Varchar, form.Text)
	formList.AddField("Created_at", "created_at", db.Timestamp, form.Datetime)
	formList.AddField("Updated_at", "updated_at", db.Timestamp, form.Datetime)

	formList.SetTable("users").SetTitle("Users").SetDescription("Users")

	return users
}
