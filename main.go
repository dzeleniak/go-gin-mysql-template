package main

import (
	"log"

	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Could not get environment variables")
	}
}

func main() {
	// Connect database
	// Start Server
}
