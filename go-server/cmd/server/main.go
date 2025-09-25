package main

import (
	"github.com/yourusername/news-server/internal/api"
	"github.com/yourusername/news-server/internal/config"
	"github.com/yourusername/news-server/internal/repository"
)

func main() {
	appConfig := config.Load()
	repository.InitDB(appConfig.DB_FILE_PATH)
	defer func() {
		err := repository.Close()
		if err != nil {

		}
	}()
	api.InitFiber(appConfig.SERVER_PORT)
}
