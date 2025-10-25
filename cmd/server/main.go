package main

import (
	"news-server/db"
	"news-server/internal/config"
	"news-server/internal/repository"
)

func main() {
	cfg := config.Load()
	db := database.Init(cfg.DbFilePath)
	repo := repository.NewNewsRepository(db)
}
