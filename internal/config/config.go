package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type DBConfig struct {
	PublicHost string
	Port       string
	User       string
	Password   string
	Name       string
}

var Env = initConfig()

func initConfig() DBConfig {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using default database configuration")
	}

	return DBConfig{
		PublicHost: getEnv("PUBLIC_HOST", "localhost"),
		Port:       getEnv("DB_PORT", "5432"),
		User:       getEnv("DB_USER", "root"),
		Password:   getEnv("DB_PASSWORD", "1234"),
		Name:       getEnv("DB_NAME", "blog"),
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return fallback
}
