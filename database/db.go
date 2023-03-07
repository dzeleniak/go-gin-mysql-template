package database

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func Init() {
	log.Printf("Connecting database\n")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")

	// "USER:PASSWORD@tcp(HOST)/DBNAME"
	url := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local", user, password, host, dbName)

	database, err := gorm.Open(mysql.Open(url))
	if err != nil {
		log.Fatalf("Could not connect to database\n")
	}
	log.Printf("Database connected successfully\n")
	db = database
}

func GetDatabase() *gorm.DB {
	return db
}
