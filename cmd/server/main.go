package main

import (
	"fmt"
	"github.com/bekhuli/go-blog/internal/common"
	"github.com/bekhuli/go-blog/internal/routes"
	"log"
	"net/http"

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

	router := routes.InitRouter(db.DB)

	addr := fmt.Sprintf("%s:%s", common.ServerEnv.Host, common.ServerEnv.Port)
	log.Println("Server running on port:", common.ServerEnv.Port)
	log.Fatal(http.ListenAndServe(addr, router))
}
