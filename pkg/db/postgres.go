package db

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/bekhuli/go-blog/internal/common"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func Connect() {
	databaseURL := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=%s",
		common.DBEnv.User,
		common.DBEnv.Password,
		common.DBEnv.PublicHost,
		common.DBEnv.Port,
		common.DBEnv.Name,
		common.DBEnv.SSLMode,
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
