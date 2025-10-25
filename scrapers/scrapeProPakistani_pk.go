package scrapers

import (
	"fmt"
	"news-server/internal/model"
	"strings"

	"github.com/gocolly/colly"
)

func ScrapeProPakistaniPk() ([]model.News, error) {
	var newsList []model.News
	url := "https://propakistani.pk"
	c := colly.NewCollector(
		colly.AllowedDomains("propakistani.pk"),
		colly.UserAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/118 Safari/537.36"),
	)

	c.OnHTML(".g1-mosaic-item, .tnews-inner", func(el *colly.HTMLElement) {
		link := el.ChildAttr("a", "href")
		title := el.ChildText("h3, h5")
		image := el.ChildAttr("img", "src")
		linkParts := strings.Split(link, "/")
		slug := linkParts[6]
		date := fmt.Sprintf("%s-%s-%s", linkParts[3], linkParts[4], linkParts[5])
		newsList = append(newsList, model.News{
			NewsUrl:     link,
			Title:       title,
			ImageUrl:    &image,
			Slug:        slug,
			PublishedOn: date,
			Source:      "propakistani.pk",
			Content:     "Dummy",
		})
	})

	err := c.Visit(url)
	if err != nil {
		return nil, err
	}

	return newsList, nil
}
