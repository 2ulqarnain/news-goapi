package scrapers

import (
	"fmt"
	"log"
	"strings"
	"sync"
	"time"

	"github.com/gocolly/colly"
	"github.com/yourusername/news-server/internal/model"
)

var siteURL = "https://radio.gov.pk"

func ScrapeRadioPakistan() ([]model.News, error) {
	var newsList []model.News
	var newsURLs []string

	// Create collector
	c := colly.NewCollector(
		colly.AllowedDomains("radio.gov.pk"),
	)

	// Find all links that have a child with class containing 'title'
	c.OnHTML("a:has([class*='title'])", func(e *colly.HTMLElement) {
		newsLink := e.Attr("href")
		if newsLink == "" {
			return
		}

		// Extract headline text
		headline := e.ChildText("[class*='title']")

		// Extract image (if exists)
		imageURL := ""
		if src := e.ChildAttr("img", "src"); src != "" {
			if strings.HasPrefix(src, "http") {
				imageURL = src
			} else {
				imageURL = "https:" + src
			}
		}

		// Extract published_on and slug from URL
		parts := strings.Split(strings.Trim(newsLink, "/"), "/")
		if len(parts) < 2 {
			return
		}
		publishedOn := parts[len(parts)-2]
		slug := parts[len(parts)-1]

		newsItem := model.News{
			Slug:        slug,
			Headline:    headline,
			PublishedOn: publishedOn,
			NewsUrl:     siteURL + newsLink,
			ImageUrl:    &imageURL,
		}
		newsList = append(newsList, newsItem)
		newsURLs = append(newsURLs, siteURL+newsLink)
		W
	})

	// Handle errors
	c.OnError(func(_ *colly.Response, err error) {
		log.Println("Request failed:", err)
		return
	})

	// After scraping
	c.OnScraped(func(_ *colly.Response) {
		fmt.Println("Site Scraped Successfully!")
	})

	// Start
	fmt.Println("Visiting", siteURL)
	err := c.Visit(siteURL)
	if err != nil {
		return nil, err
	}

	allNewsBodies := ScrapeMultipleNews(newsURLs)
	for i := range newsList {
		newsList[i].Body = allNewsBodies[i]
	}

	return newsList, nil
}

func ScrapeRadioPkNewsByURL(url string) string {
	var news string
	c := colly.NewCollector(
		colly.UserAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/118 Safari/537.36"),
	)
	err := c.Limit(&colly.LimitRule{
		DomainGlob:  "*radio.gov.pk*",
		RandomDelay: 3 * time.Second,
	})
	if err != nil {
		return ""
	}
	c.OnHTML(".newsdetailcontent", func(e *colly.HTMLElement) {
		news = e.Text
	})
	c.OnError(func(_ *colly.Response, err error) {
		log.Println("Request failed:", err)
		return
	})

	if err := c.Visit(url); err != nil {
		log.Println(err)
	}
	return news
}

func ScrapeMultipleNews(urls []string) []string {
	newsList := make([]string, len(urls))
	var wg sync.WaitGroup
	for i, url := range urls {
		wg.Add(1)
		go func(i int, url string) {
			defer wg.Done()
			newsList[i] = ScrapeRadioPkNewsByURL(url)
		}(i, url)
	}
	wg.Wait()
	return newsList
}
