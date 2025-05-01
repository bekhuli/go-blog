package db

import (
	"database/sql"
	"fmt"
	"github.com/bekhuli/go-blog/internal/common/config"
	"log"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func Connect() {
	databaseURL := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=%s",
		config.DBEnv.User,
		config.DBEnv.Password,
		config.DBEnv.PublicHost,
		config.DBEnv.Port,
		config.DBEnv.Name,
		config.DBEnv.SSLMode,
	)

	var err error
	DB, err = sql.Open("postgres", databaseURL)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	if err = DB.Ping(); err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}

	log.Println("Database connected successfully")

}
