package main

import (
	"log"

	"github.com/bekhuli/go-blog/pkg/db"
)

func main() {
	db.Connect()

	defer func() {
		log.Println("Closing database connection")
		if err := db.DB.Close(); err != nil {
			log.Println("Failed to close database:", err)
		} else {
			log.Println("Database disconnected successfully")
		}
	}()
}
