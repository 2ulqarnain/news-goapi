package handlers

import (
	"fmt"
	"strconv"

	"news-server/internal/model"
	"news-server/internal/repository"

	"github.com/gofiber/fiber/v2"
)

func GetAllNews(ctx *fiber.Ctx) error {
	limitParam := ctx.Query("limit", "0")
	limit, err := strconv.Atoi(limitParam)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(model.Response{
			Message: "Invalid limit parameter!",
			Data:    nil,
		})
	}
	news, err := repository.GetAllNews(limit)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(model.Response{
			Ok:      false,
			Message: fmt.Sprintf("%v", err),
			Data:    nil,
		})
	}
	return ctx.JSON(model.Response{
		Ok:   true,
		Data: news,
	})
}

func GetNewsBySlug(ctx *fiber.Ctx) error {
	slug := ctx.Params("slug")
	return ctx.SendString(slug)
}
