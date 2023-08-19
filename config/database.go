package config

import (
	"backend/cake/entity"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB = nil

func init(){
	db,err := gorm.Open(sqlite.Open("cake_database.db"),&gorm.Config{})
	
	if err != nil{
		Sugar.Error(err)			
	}

	db.AutoMigrate(&entity.User{})

	DB = db

}
