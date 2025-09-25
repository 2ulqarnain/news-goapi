package handlers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/yourusername/news-server/internal/model"
	"github.com/yourusername/news-server/internal/repository"
)

func GetAllNews(ctx *fiber.Ctx) error {
	news, err := repository.GetAllNews()
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(model.Response{
			Status:  fiber.StatusInternalServerError,
			Message: fmt.Sprintf("%v", err),
			Data:    nil,
		})
	}
	return ctx.JSON(model.Response{
		Status: fiber.StatusOK,
		Data:   news,
	})
}

func GetNewsBySlug(ctx *fiber.Ctx) error {
	slug := ctx.Params("slug")
	return ctx.SendString(slug)
}
