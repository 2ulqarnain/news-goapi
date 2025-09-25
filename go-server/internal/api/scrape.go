package api

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/yourusername/news-server/scrapers"
)

func InitScrape(ctx *fiber.Ctx) error {
	newsList, err := scrapers.ScrapeRadioPakistan()

	if err != nil {
		fmt.Println(err)
	}

	return ctx.JSON(newsList)

}
