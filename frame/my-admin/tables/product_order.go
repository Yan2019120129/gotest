package tables

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template/types/form"
)

func GetProductOrderTable(ctx *context.Context) table.Table {

	productOrder := table.NewDefaultTable(table.DefaultConfigWithDriver("mysql").SetPrimaryKey("id", db.Bigint))

	info := productOrder.GetInfo().HideFilterArea()

	info.AddField("Id", "id", db.Bigint).
		FieldFilterable()
	info.AddField("Created_at", "created_at", db.Datetime)
	info.AddField("Updated_at", "updated_at", db.Datetime)
	info.AddField("Deleted_at", "deleted_at", db.Datetime)
	info.AddField("Admin_id", "admin_id", db.Int)
	info.AddField("User_id", "user_id", db.Int)
	info.AddField("Product_id", "product_id", db.Int)
	info.AddField("Order_sn", "order_sn", db.Varchar)
	info.AddField("Money", "money", db.Decimal)
	info.AddField("Fee", "fee", db.Decimal)
	info.AddField("Type", "type", db.Int)
	info.AddField("Status", "status", db.Tinyint)
	info.AddField("Data", "data", db.Text)
	info.AddField("Expired_at", "expired_at", db.Datetime)

	info.SetTable("product_order").SetTitle("ProductOrder").SetDescription("ProductOrder")

	formList := productOrder.GetForm()
	formList.AddField("Id", "id", db.Bigint, form.Default)
	formList.AddField("Created_at", "created_at", db.Datetime, form.Datetime)
	formList.AddField("Updated_at", "updated_at", db.Datetime, form.Datetime)
	formList.AddField("Deleted_at", "deleted_at", db.Datetime, form.Datetime)
	formList.AddField("Admin_id", "admin_id", db.Int, form.Number)
	formList.AddField("User_id", "user_id", db.Int, form.Number)
	formList.AddField("Product_id", "product_id", db.Int, form.Number)
	formList.AddField("Order_sn", "order_sn", db.Varchar, form.Text)
	formList.AddField("Money", "money", db.Decimal, form.Currency)
	formList.AddField("Fee", "fee", db.Decimal, form.Text)
	formList.AddField("Type", "type", db.Int, form.Number)
	formList.AddField("Status", "status", db.Tinyint, form.Number)
	formList.AddField("Data", "data", db.Text, form.RichText)
	formList.AddField("Expired_at", "expired_at", db.Datetime, form.Datetime)

	formList.SetTable("product_order").SetTitle("ProductOrder").SetDescription("ProductOrder")

	return productOrder
}
