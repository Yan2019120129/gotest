package app

import (
	"errors"
	_ "github.com/GoAdminGroup/go-admin/adapter/gin" // web framework adapter
	"github.com/GoAdminGroup/go-admin/engine"
	_ "github.com/GoAdminGroup/go-admin/modules/db/drivers/mysql" // sql driver
	"github.com/GoAdminGroup/go-admin/template"
	"github.com/GoAdminGroup/go-admin/template/chartjs"
	_ "github.com/GoAdminGroup/themes/sword" // ui theme
	"github.com/gin-gonic/gin"
	"io"
	"log"
	admin "my-admin/app/admin/router"
	web "my-admin/app/web/router"
	"my-admin/configs"
	"my-admin/tables"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

// InitServer 初始化服务
func InitServer() {
	template.AddComp(chartjs.NewChart())
	gin.DefaultWriter = io.Discard
	r := gin.Default()
	admin.InitRouter(r)

	adminConfig := configs.GetGoAdmin()
	eng := engine.Default()
	if err := eng.AddConfigFromJSON(adminConfig.ConfigPath).
		AddGenerators(tables.Generators).
		Use(r); err != nil {
		panic(err)
	}
	web.InitRouter(eng)

	ginConfig := configs.GetGin()
	if err := r.Run(ginConfig.Port); err != nil && errors.Is(err, http.ErrServerClosed) {
		log.Printf("listen: %s\n", err)
	}

	// 检测系统信号关闭
	// syscall.SIGINT 中断信号，通常在用户按下 Ctrl+C 时发送。
	// syscall.SIGTERM 终止信号，通常用于请求程序终止。
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	eng.MysqlConnection().Close()
}
