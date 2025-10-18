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

	app.Get("/", func(c *fiber.Ctx) error { return c.SendString("News Server Running...") })
	registerNewsRoutes(app)
	registerScrapeRoutes(app)

	handle404Errors(app)

	if err := app.Listen(":" + port); err != nil {
		log.Fatalf("%v", err)
	}
}

func registerScrapeRoutes(app *fiber.App) {
	scrape := app.Group("/crawl")
	scrape.Get("/", InitScrape)
}

func registerNewsRoutes(app *fiber.App) {
	news := app.Group("/news")
	news.Get("/", GetAllNews)
	news.Get("/:slug", GetNewsBySlug)
}

func handle404Errors(app *fiber.App) {
	app.Use(func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"ok":      false,
			"message": "Not Found",
		})
	})
}
