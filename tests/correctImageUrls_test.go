package tests

import (
	"fmt"
	"strings"
	"testing"

	"github.com/gocolly/colly"
)

func TestCorrectImageUrlsScraped(t *testing.T) {
	c := colly.NewCollector(
		colly.AllowedDomains("radio.gov.pk"),
	)

	var imageLinks []string

	c.OnHTML("div.col-md-6.col-lg-6:has(img[src*='newsimage'])", func(el *colly.HTMLElement) {
		title := el.ChildText("[class*='title']")
		imageSrc := el.ChildAttr("img[src*='newsimage']", "src")
		if imageSrc == "" || !strings.Contains(imageSrc, "newsimage") {
			t.Errorf("Image src is empty for news: %s\n", title)
		}
		imageLinks = append(imageLinks, imageSrc)
	})
	c.OnScraped(func(r *colly.Response) {
		fmt.Println("Scrap Finished!")
	})
	err := c.Visit("https://radio.gov.pk")
	if err != nil {
		t.Error(err)
	}
}
