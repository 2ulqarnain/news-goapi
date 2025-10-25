package main

import (
	"database/sql"
	"log"
	"news-server/internal/config"
	"news-server/internal/handlers"
	"news-server/internal/repository"
)

func main() {
	appConfig := config.Load()
	db := repository.InitDB(appConfig.DbFilePath)
	handlers.InitFiber(appConfig.ServerPort)
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Printf("%v", err)
		}
	}(db)
	//repo := repository.NewNewsRepository(db)
	//newsService := service.NewNewsService(repo)
}
