package routes

import (
	"database/sql"
	"news-server/internal/handlers"
	"news-server/internal/repository"
	"news-server/internal/service"

	"github.com/gofiber/fiber/v2"
)

func Register(app *fiber.App, db *sql.DB) {
	newsRepo := repository.NewNewsRepository(db)
	newsService := service.NewNewsService(newsRepo)
	newsHandler := handlers.NewNewsHandler(newsService)

	newsRoutes := app.Group("/news")
	newsRoutes.Get("/", newsHandler.GetAllNews)
	newsRoutes.Get("/:slug", newsHandler.GetNewsWithContentBySlug)
}
