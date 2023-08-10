package main

import (
	"log"
	"os"
	"task/config"
	"task/routes"

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

	// Routes Initialize
	r := routes.Routes()

	// static file
	os.Mkdir("./uploads", os.ModePerm)
	r.Static("/uploads", "./uploads")

	// run routes
	if err := r.Run(":4444"); err != nil {
		log.Fatal(err)
	}
}
