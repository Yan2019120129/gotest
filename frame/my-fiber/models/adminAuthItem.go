package models

import (
	"database/sql"
)

type AdminAuthItem struct {
	Name        string         `json:"name"`        //名称
	Type        int            `json:"type"`        //类型
	Description string         `json:"description"` //描述
	RuleName    string         `json:"rule_name"`   //规则名称
	Data        sql.NullString `json:"data"`        //数据
	CreatedAt   int            `json:"created_at"`  //创建时间
	UpdatedAt   int            `json:"updated_at"`  //更新时间
}

const (
	// AdminAuthItemTypeManage 管理员名称
	AdminAuthItemTypeManage = 1
	// AdminAuthItemTypeRoute 请求路由
	AdminAuthItemTypeRoute = 2
	// AdminAuthItemTypeRouteName 请求路由名称
	AdminAuthItemTypeRouteName = 3
)
