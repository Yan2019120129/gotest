package gin

import (
	"github.com/gin-gonic/gin"
	"gotest/my_frame/api/admin"
	"gotest/my_frame/models"
	"net/http"
	"time"
)

func Init(cfg *models.GinConfig) {

	// 配置gin
	router := gin.Default()
	s := &http.Server{
		Addr:           cfg.Port,
		Handler:        router,
		ReadTimeout:    time.Duration(cfg.ReadTimeout) * time.Second,
		WriteTimeout:   time.Duration(cfg.WriteTimeout) * time.Second,
		MaxHeaderBytes: cfg.MaxHeaderBytes,
	}

	// 初始化后台路由
	admin.InitRouter(router)

	// 运行服务
	err := s.ListenAndServe()
	if err != nil {
		panic(err)
		return
	}
}
