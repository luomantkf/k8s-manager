package routers

import (
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"private-cloud-k8s/conf"
	_ "private-cloud-k8s/docs"
	"private-cloud-k8s/middleware"
	"private-cloud-k8s/view/result"
)

var Router *gin.Engine

func NewRouter() *gin.Engine {
	Router = gin.Default()

	//设置跨域
	Router.Use(middleware.Cors())

	Router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	Router.NoRoute(func(c *gin.Context) {
		c.JSON(200, result.FailResultWithMsg("请求无法进行响应！！！"))
	})

	//配置各种路由信息
	r := Router.Group(conf.Conf.ServerConfig.ServerContextPath)
	CloudK8sServerRouter(r, "/k8s")
	return Router
}
