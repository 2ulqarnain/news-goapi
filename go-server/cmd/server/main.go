package main

import (
	"github.com/yourusername/news-server/internal/api"
	"github.com/yourusername/news-server/internal/config"
	"github.com/yourusername/news-server/internal/repository"
)

func main() {
	appConfig := config.Load()
	repository.InitDB(appConfig.DB_FILE_PATH)
	defer repository.Close()
	api.InitFiber(appConfig.SERVER_PORT)
}
