package web

import (
	"github.com/GoAdminGroup/go-admin/engine"
	"my-admin/app/web/other/pages"
)

// InitRouter 初始化路由
func InitRouter(eng *engine.Engine) {
	eng.HTML("GET", "/admin", pages.GetDashBoard)
	eng.HTMLFile("GET", "/admin/hello", "./app/web/app/template/hello.tmpl", map[string]interface{}{
		"msg": "Hello world",
	})
}
