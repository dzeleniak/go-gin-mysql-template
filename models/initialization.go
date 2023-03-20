package models

import (
	"log"

	"github.com/dzeleniak/go-gin-mysql-template/database"
)

type model interface {}

var tables = []model{
	&ObjectModel{},
}

func Init() {
	for _, model := range tables {
		if database.GetDatabase().AutoMigrate(model) != nil {
			log.Fatalf("Could not complete database migration.\n")
		}
	}
	log.Printf("Database migration successful.\n")
}