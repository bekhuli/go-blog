package common

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type DBConfig struct {
	PublicHost string
	Port       string
	User       string
	Password   string
	Name       string
	SSLMode    string
}

type JWTConfig struct {
	JWTSecret              string
	JWTExpirationInSeconds int64
}

type ServerConfig struct {
	Host string
	Port string
}

var DBEnv = initDBConfig()
var JWTEnv = initJWTConfig()
var ServerEnv = initServerConfig()

func initDBConfig() DBConfig {
	if err := godotenv.Load(); err != nil {
		log.Fatal("No .env file found")
	}

	return DBConfig{
		PublicHost: getEnv("PUBLIC_HOST"),
		Port:       getEnv("DB_PORT"),
		User:       getEnv("DB_USER"),
		Password:   getEnv("DB_PASSWORD"),
		Name:       getEnv("DB_NAME"),
		SSLMode:    getEnv("DB_SSL_MODE"),
	}
}

func initJWTConfig() JWTConfig {
	godotenv.Load()

	return JWTConfig{
		JWTSecret:              getEnv("JWT_SECRET"),
		JWTExpirationInSeconds: getEnvAsInt("JWT_EXP"),
	}
}

func initServerConfig() ServerConfig {
	godotenv.Load()

	return ServerConfig{
		Host: getEnv("SERVER_HOST"),
		Port: getEnv("SERVER_PORT"),
	}
}

func getEnv(key string) string {
	value, ok := os.LookupEnv(key)
	if !ok {
		log.Fatalf("Environment variable %s not set", key)
	}

	return value
}

func getEnvAsInt(key string) int64 {
	valueStr := getEnv(key)
	valueInt, err := strconv.ParseInt(valueStr, 10, 64)
	if err != nil {
		log.Fatalf("Invalid int value for environment variable %s: %v", key, err)
	}

	return valueInt
}
