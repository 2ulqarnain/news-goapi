package main

import (
	"github.com/yourusername/news-server/internal/config"
	"github.com/yourusername/news-server/internal/handlers"
	"github.com/yourusername/news-server/internal/repository"
)

func main() {
	appConfig := config.Load()
	repository.InitDB(appConfig.DbFilePath)
	defer func() {
		err := repository.Close()
		if err != nil {

		}
	}()
	handlers.InitFiber(appConfig.ServerPort)
}
