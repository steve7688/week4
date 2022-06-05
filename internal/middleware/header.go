package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func HeaderHandler() gin.HandlerFunc {

	return func(c *gin.Context) {
		headers := c.Request.Header
		fmt.Println(headers) //打印请求头
	}

}
