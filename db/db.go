package db

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"joiner-management/models"
	"log"
)

const DB_USERNAME = "root"
const DB_PASSWORD = "admin"
const DB_NAME = "db_joiner"
const DB_HOST = "127.0.0.1"
const DB_PORT = "3306"

func InitDb() *gorm.DB {
	dbUrl := DB_USERNAME + ":" + DB_PASSWORD + "@tcp" + "(" + DB_HOST + ":" + DB_PORT + ")/" + DB_NAME + "?" + "parseTime=true&loc=Local"

	db, err := gorm.Open(mysql.Open(dbUrl), &gorm.Config{})

	if err != nil {
		log.Fatalf("Error connecting to database : error=%v\n\n", err)
	}

	db.AutoMigrate(&models.Joiner{})

	return db
}
