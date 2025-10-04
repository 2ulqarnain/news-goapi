package handlers

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func InitFiber(port string) {
	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
	}))

	app.Use(logger.New())

	registerNewsRoutes(app)
	registerScrapeRoutes(app)

	if err := app.Listen(":" + port); err != nil {
		log.Fatalf("%v", err)
	}
}

func registerScrapeRoutes(app *fiber.App) {
	scrape := app.Group("/scrape")
	scrape.Get("/init", InitScrape)
}

func registerNewsRoutes(app *fiber.App) {
	news := app.Group("/news")
	news.Get("/", GetAllNews)
	news.Get("/:slug", GetNewsBySlug)
}
