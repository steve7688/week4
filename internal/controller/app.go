package controller

import (
	"Week4/internal/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AppInfo(c *gin.Context) {

	c.IndentedJSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "ok",
		"data": &model.UserModel{
			ID:      "1",
			Name:    "Steve",
			Sex:     0,
			Address: "深圳",
		},
	})

}
