package main

import (
	"news-server/internal/config"
	"news-server/internal/handlers"
	"news-server/internal/repository"
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
