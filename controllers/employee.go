package controllers

import (
	"happyhr/db"
	"happyhr/models"

	"happyhr/router"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func GetEmployeeALL(c *gin.Context) {
	result := db.DB.Find(&models.Employee{})
	if result.Error != nil {
		log.Error(result.Error)
		c.JSON(500, gin.H{
			"msg":  "select db err",
			"code": "code",
		})
	}
	if result.RowsAffected != 0 {
		c.JSON(200, gin.H{
			"data": result,
		})
	} else {
		c.JSON(200, gin.H{
			"msg":  "not employee",
			"data": nil,
		})
	}
}

func init() {
	router.Route.GET("/employees", GetEmployeeALL)
}
