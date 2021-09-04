package router

import (
	"github.com/gin-gonic/gin"
)

var Route *gin.Engine = gin.Default()

func ping() {

	Route.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
}

func init() {
	ping()
}
