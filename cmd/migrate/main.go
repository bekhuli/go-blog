package main

import (
	"fmt"
	"github.com/bekhuli/go-blog/internal/common"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"log"
	"os"
)

const migrationsDir = "file://migrations"

func main() {
	if len(os.Args) < 2 {
		log.Fatal("missing command: up | down | drop")
	}
	command := os.Args[1]

	cfg := common.DBEnv

	dbURL := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=%s",
		cfg.User,
		cfg.Password,
		cfg.PublicHost,
		cfg.Port,
		cfg.Name,
		cfg.SSLMode,
	)

	m, err := migrate.New(migrationsDir, dbURL)
	if err != nil {
		log.Fatalf("failed to create migrate instance: %v", err)
	}

	switch command {
	case "up":
		if err := m.Up(); err != nil && err != migrate.ErrNoChange {
			log.Fatalf("migration up failed: %v", err)
		}
		log.Println("migration up complete")
	case "down":
		if err := m.Steps(-1); err != nil {
			log.Fatalf("migration down failed: %v", err)
		}
		log.Println("migration down one step")
	case "drop":
		if err := m.Drop(); err != nil {
			log.Fatalf("migration drop failed: %v", err)
		}
		log.Println("all tables dropped")
	default:
		log.Fatalf("unknown command: %s", command)
	}
}
