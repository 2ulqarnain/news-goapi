package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DB_FILE_PATH string
	SERVER_PORT  string
}

func Load() Config {

	if err := godotenv.Load(); err != nil {
		log.Fatalf("Couldn't Load .env file: %v", err)
	}

	PORT := os.Getenv("SERVER_PORT")

	if PORT == "" {
		log.Fatal("Missing SERVER_PORT .env variable!")
	}

	DB_FILE_PATH := os.Getenv("DB_FILE_PATH")

	if DB_FILE_PATH == "" {
		log.Fatal("Missing DB_FILE_PATH .env variable")
	}

	return Config{
		DB_FILE_PATH: DB_FILE_PATH,
		SERVER_PORT:  PORT,
	}
}
