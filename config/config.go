package config

import (
	"learn-gonic/structs"

	"github.com/jinzhu/gorm"
)

func DBinit() *gorm.DB {
	db, err := gorm.Open("mysql", "root:baterai@tcp(127.0.0.1:3306)/golang?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(structs.Person{})

	return db
}
