package main

import (
	"github.com/bekhuli/go-blog/pkg/db"
	"log"
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
