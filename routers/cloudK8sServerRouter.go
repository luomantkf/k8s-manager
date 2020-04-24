package routers

import (
	"github.com/gin-gonic/gin"
	"private-cloud-k8s/view/result"
)

func CloudK8sServerRouter(router *gin.RouterGroup, base string) {
	r := router.Group(base)
	r.GET("/test/:id", Test)
}

// @Summary 测试
// @Produce  json
// @Param   id     path    int     true		"测试swagger参数"
// @Success 200 {object}  result.Result
// @Router /api/v1/k8s/test/{id} [get]
func Test(c *gin.Context) {
	c.JSON(200, result.OkResultWithMsg("测试成功"+c.Param("id")))
}
