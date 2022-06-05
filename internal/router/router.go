package router

import (
	"Week4/internal/controller"
	"Week4/internal/middleware"
	"github.com/gin-gonic/gin"
)

func InitRouter() (c *gin.Engine) {

	r := gin.Default()

	r.Use(middleware.HeaderHandler()) //对请求头进行验签
	//r.Use(middleware.UUID())          //统计uuid(用户数)

	//app相关接口
	appGroup := r.Group("/app")
	{
		appGroup.GET("/info", controller.AppInfo)
	}

	return r
}
