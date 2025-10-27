package handlers

import (
	"fmt"
	"news-server/internal/model"
	"news-server/scrapers"

	"github.com/gofiber/fiber/v2"
)

func ScrapeAllSites(ctx *fiber.Ctx) error {
	newsList, err := scrapers.ScrapeRadioPakistan()

	if err != nil {
		fmt.Println(err)
	}
	if len(newsList) == 0 {
		return ctx.JSON(model.Response{
			Ok:      false,
			Message: "No News Scraped, Likely DOM has changed!",
		})
	}

	return ctx.JSON(model.Response{
		Ok:      true,
		Message: "Crawled Successfully!",
	})

}
