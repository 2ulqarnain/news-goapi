package main

import (
	"database/sql"
	"fmt"
	"log"
	"news-server/db"
	"news-server/internal/config"
	"news-server/internal/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	cfg := config.Load()
	db := database.Init(cfg.DbFilePath)
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(db)
	app := fiber.New(fiber.Config{
		AppName:      "News Server v1.0.0",
		ServerHeader: "News-Server",
	})
	routes.Register(app, db)
	if err := app.Listen(":" + cfg.ServerPort); err != nil {
		log.Fatal(err)
	}
}
