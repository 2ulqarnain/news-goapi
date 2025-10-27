package handlers

import (
	"fmt"
	"news-server/internal/model"
	"news-server/internal/service"

	"github.com/gofiber/fiber/v2"
)

type NewsHandler struct {
	svc service.NewsService
}

func NewNewsHandler(svc service.NewsService) *NewsHandler {
	return &NewsHandler{svc: svc}
}

func (h *NewsHandler) GetAllNews(c *fiber.Ctx) error {
	ctx := c.Context()
	news, err := h.svc.GetAllNews(ctx)
	if err != nil {
		fmt.Println(err)
		return c.Status(fiber.StatusInternalServerError).JSON(model.Response{
			Ok:      false,
			Message: fiber.ErrInternalServerError.Message,
			Data:    nil,
		})
	}
	return c.JSON(model.Response{
		Ok:   true,
		Data: news,
	})
}

func (h *NewsHandler) GetNewsWithContentBySlug(c *fiber.Ctx) error {
	slug := c.Params("slug")
	return c.SendString(slug)
}
