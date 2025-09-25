package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DbFilePath string
	ServerPort string
}

func Load() Config {

	if err := godotenv.Load(); err != nil {
		log.Fatalf("Couldn't Load .env file: %v", err)
	}

	PORT := os.Getenv("SERVER_PORT")

	if PORT == "" {
		log.Fatal("Missing SERVER_PORT .env variable!")
	}

	DbFilePath := os.Getenv("DB_FILE_PATH")

	if DbFilePath == "" {
		log.Fatal("Missing DB_FILE_PATH .env variable")
	}

	return Config{
		DbFilePath: DbFilePath,
		ServerPort: PORT,
	}
}
