package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"private-cloud-k8s/conf"
	"private-cloud-k8s/middleware"
	"private-cloud-k8s/pkg/logging"
	"private-cloud-k8s/routers"
	"private-cloud-k8s/view/result"
	"strconv"
)

func Load() {
	gin.SetMode(conf.Conf.RunMode)

	router := routers.NewRouter()
	//配置跨域
	router.Use(middleware.Cors())

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, result.OkResultWithMsg("成功"))
	})

	//配置http服务器服务器
	serverConf := conf.Conf.ServerConfig
	server := http.Server{
		Addr:           fmt.Sprintf(":%d", serverConf.HTTPPort),
		Handler:        router,
		ReadTimeout:    serverConf.ReadTimeout,
		WriteTimeout:   serverConf.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	logging.Info("启动HTTP服务器成功，监听端口：【{}】", strconv.Itoa(serverConf.HTTPPort))
	server.ListenAndServe()
}
