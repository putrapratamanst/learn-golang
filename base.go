package main

import (
	"learn-gonic/controllers"

	"learn-gonic/config"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db := config.DBinit()
	inDB := &controllers.InDB{DB: db}

	router := gin.Default()

	router.GET("person/:id", inDB.GetPerson)

	router.Run(":3000")

}
//gaskeun