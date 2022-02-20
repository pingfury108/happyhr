package controllers

import (
	"fmt"
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
	if err := c.Bind(&query); err != nil {
		log.Error(err)
		c.JSON(400, metadata.QueryResult{
			Msg:  fmt.Sprint(err),
			Code: 400,
			Data: []models.Employee{},
		})
	}
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

	if err := c.ShouldBindJSON(&employee); err != nil {
		log.Error(err)
		c.JSON(400, metadata.CreateResult{
			Msg:  fmt.Sprint(err),
			Code: 400,
		})
		return
	}
	log.Debug(employee)
	if err := db.DB.Create(&employee).Error; err != nil {
		log.Errorf("Create employee error: %v", err)
		c.JSON(500, metadata.CreateResult{
			Msg:  fmt.Sprint(err),
			Code: 500,
		})
		return
	}
	c.JSON(200, metadata.CreateResult{
		Msg:  "ok",
		Code: 200,
	})
}

func Update(c *gin.Context) {
	//var from_employee metadata.UpdateEmployee
	var from_employee models.Employee
	var employee models.Employee

	if err := c.ShouldBindJSON(&from_employee); err != nil {
		log.Error(err)
		c.JSON(400, metadata.UpdateResult{
			Msg:  fmt.Sprint(err),
			Code: 400,
		})
		return
	}
	log.Debug(from_employee)
	if from_employee.ID == 0 {
		c.JSON(400, metadata.UpdateResult{
			Msg:  "employee id can not be empty",
			Code: 400})
		return
	}
	result := db.DB.First(&employee, from_employee.ID)
	if result.Error != nil {
		log.Error(result.Error)
		return
	}
	db.DB.Model(&employee).Updates(&from_employee)
	c.JSON(200, metadata.UpdateResult{
		Msg:  "ok",
		Code: 200,
	})
}

func Delete(c *gin.Context) {
	var parameter metadata.DeleteEmployee
	var employee models.Employee

	if err := c.Bind(&parameter); err != nil {
		log.Error(err)
		c.JSON(400, metadata.Result{
			Msg:  fmt.Sprint(err),
			Code: 400,
		})
	}
	if parameter.ID == 0 {
		c.JSON(400, metadata.Result{
			Msg:  "employee id can not be empty",
			Code: 400})
		return
	}
	result := db.DB.First(&employee, parameter.ID)
	if result.Error != nil {
		log.Error(result.Error)
		c.JSON(500, metadata.Result{
			Msg:  fmt.Sprint(result.Error),
			Code: 400,
		})
		return
	}
	db.DB.Delete(&employee)
	c.JSON(200, metadata.Result{
		Msg:  "ok",
		Code: 200,
	})
}

func init() {
	router.Route.GET("/employees", GetALL)
	router.Route.POST("/employees/query", Query)
	router.Route.POST("/employees", Create)
	router.Route.PUT("/employees", Update)
	router.Route.DELETE("/employees", Delete)
}
