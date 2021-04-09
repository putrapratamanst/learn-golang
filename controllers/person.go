package controllers

import (
	"fmt"
	"learn-gonic/structs"
	"net/http"

	"github.com/gin-gonic/gin"
)

//to get one data from database
func (idb *InDB) GetPerson(c *gin.Context) {
	var (
		person structs.Person
		result gin.H
	)

	id := c.Param("id")
		err := idb.DB.Where("id = ?", id).First(&person).Error

	if err != nil {
		result = gin.H{
			"result": err.Error(),
			"count":  0,
		}
	} else {
		result = gin.H{
			"result": person,
			"count":  1,
		}
	}

	fmt.Println(result)
	c.JSON(http.StatusOK, result)
}
