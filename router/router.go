package router

import (
	"happyhr/controllers"

	"github.com/gin-gonic/gin"
)

var route *gin.Engine = gin.Default()

func Create_gin() *gin.Engine {

	route.LoadHTMLGlob("views/*")

	route.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	route.GET("/", controllers.Index)
	return route
}
