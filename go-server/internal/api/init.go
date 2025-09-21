package api

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func InitFiber(port string) {
	app := fiber.New()
	registerRoutes(app)
	if err := app.Listen(":" + port); err != nil {
		log.Fatalf("%v", err)
	}
}

func registerRoutes(app *fiber.App) {
	news := app.Group("/news")
	news.Get("/", GetAllNews)
	news.Get("/:slug", GetNewsBySlug)
}
