package main

import (
	"log"

	"github.com/dzeleniak/go-gin-mysql-template/database"
	"github.com/dzeleniak/go-gin-mysql-template/server"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Could not get environment variables")
	}
}

func main() {
	database.Init() // Connect database
	server.Init()   // Start server
}
