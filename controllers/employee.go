package controllers

import (
	"happyhr/db"
	"happyhr/models"

	"happyhr/router"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func GetALL(c *gin.Context) {
	var employess []models.Employee
	result := db.DB.Find(&employess)
	if err := result.Error; err != nil {
		log.Error(err)
		c.JSON(500, gin.H{
			"msg":  "select db err",
			"code": "code",
		})
	}
	if result.RowsAffected != 0 {
		log.Debugf("result number: %v", result.RowsAffected)
		c.JSON(200, gin.H{
			"data": employess,
		})
	} else {
		c.JSON(200, gin.H{
			"msg":  "not employee",
			"data": nil,
		})
	}
}

func GetByID(c *gin.Context) {
	id := c.Param("id")
	log.Debugf("select id is %v", id)
	employee := models.Employee{}
	result := db.DB.First(&employee, id)
	if result.Error != nil {
		log.Error(result.Error)
		c.JSON(500, gin.H{
			"msg":  "select db err",
			"code": "code",
		})
	}
	if result.RowsAffected != 0 {
		c.JSON(200, gin.H{
			"data": employee,
		})
	} else {
		c.JSON(200, gin.H{
			"msg":  "employee not found",
			"data": nil,
		})
	}
}

func GetBySerialNumber(c *gin.Context) {

}

func GetByName(c *gin.Context) {

}

func Create(c *gin.Context) {
	var employee models.Employee

	if c.ShouldBindJSON(&employee) != nil {
		log.Error("Parameter error")
	}
	log.Debug(employee)
	if err := db.DB.Create(&employee).Error; err != nil {
		log.Errorf("Create employee error: %v", err)
	}
}

func Update(c *gin.Context) {

}

func Delete(c *gin.Context) {

}

func init() {
	router.Route.GET("/employees", GetALL)
	router.Route.GET("/employees/:id", GetByID)
	router.Route.POST("/employees", Create)
}
