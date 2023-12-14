package gin

import (
	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
	"gotest/my_frame/app/admin/router"
	"gotest/my_frame/config"
)

var engin *gin.Engine

func Init() {
	cfg := config.GetGin()
	// 配置gin
	engin = gin.Default()

	// 初始化后台路由
	router.InitRouter(engin)

	err := endless.ListenAndServe(cfg.Port, engin)
	if err != nil {
		panic(err)
		return
	}
}

// GetEngin 获取gin实例
func GetEngin() *gin.Engine {
	return engin
}
