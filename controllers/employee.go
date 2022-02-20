package controllers

import (
	"happyhr/db"
	"happyhr/models"

	"happyhr/metadata"
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

func Query(c *gin.Context) {
	var query metadata.QueryEmployee
	c.Bind(&query)
	log.Debugf("bind query: %v", query)
	if query.ID != 0 {
		employee := models.Employee{}
		result := db.DB.First(&employee, query.ID)
		if result.Error != nil {
			log.Error(result.Error)
		}
		if result.RowsAffected != 0 {
			c.JSON(200, metadata.QueryResult{
				Msg:  "ok",
				Code: 200,
				Data: []models.Employee{employee},
			})
		} else {
			c.JSON(404, metadata.QueryResult{
				Msg:  "employe not found",
				Code: 404,
				Data: []models.Employee{}})
		}
	} else if query.Name != "" {
		var employees []models.Employee
		result := db.DB.Where("name <> ?", query.Name).Find(&employees)
		//result := db.DB.Where(&models.Employee{Name: query.Name}).Find(&employees)
		if result.Error != nil {
			log.Error(result.Error)
		}
		log.Debug(result.RowsAffected)
		if result.RowsAffected != 0 {
			c.JSON(200, metadata.QueryResult{
				Msg:  "ok",
				Code: 200,
				Data: employees})
		} else {
			c.JSON(404, metadata.QueryResult{
				Msg:  "employe not found",
				Code: 404,
				Data: []models.Employee{}})
		}
	} else if query.SerialNumber != 0 {
		employee := models.Employee{}
		result := db.DB.Where("serial_number = ?", query.SerialNumber).First(&employee)
		if result.Error != nil {
			log.Error(result.Error)
			if result.RowsAffected != 0 {
				c.JSON(200, metadata.QueryResult{
					Msg:  "ok",
					Code: 200,
					Data: []models.Employee{employee},
				})
			} else {
				c.JSON(404, metadata.QueryResult{
					Msg:  "employe not found",
					Code: 404,
					Data: []models.Employee{}})
			}

		} else {
			c.JSON(400, metadata.QueryResult{
				Msg:  "wrong request parameter",
				Code: 400,
				Data: []models.Employee{}})
		}
	}
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
	router.Route.POST("/employees/query", Query)
	router.Route.POST("/employees", Create)
}
