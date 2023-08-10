package main

import (
	"log"
	"task/config"

	"github.com/joho/godotenv"
)

func main() {
	// Loading .env file
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	// Database instance
	db, err := config.ConnDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

}
