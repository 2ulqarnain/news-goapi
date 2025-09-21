package main

import (
	"github.com/yourusername/news-server/internal/api"
	"github.com/yourusername/news-server/internal/config"
	"github.com/yourusername/news-server/internal/repository"
)

func main() {
	config := config.Load()
	repository.InitDB(config.DB_FILE_PATH)
	defer repository.Close()
	api.InitFiber(config.SERVER_PORT)
}
