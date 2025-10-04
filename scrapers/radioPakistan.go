package scrapers

import (
	"fmt"
	"log"
	"strings"
	"sync"
	"time"

	"github.com/gocolly/colly"
	"github.com/yourusername/news-server/internal/model"
	"github.com/yourusername/news-server/internal/repository"
	"github.com/yourusername/news-server/internal/utils"
)

var siteURL = "https://radio.gov.pk"

func ScrapeRadioPakistan() ([]model.News, error) {
	var newsList []model.News
	var newsURLs []string

	// Create collector
	c := colly.NewCollector(
		colly.AllowedDomains("radio.gov.pk"),
		colly.UserAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/118 Safari/537.36"),
	)

	seen := make(map[string]bool)
	//Find all links that have a child with class containing 'title'
	c.OnHTML("div.col-md-6.col-lg-6:has(img[src*='newsimage'])", func(e *colly.HTMLElement) {
		aTagChildren := e.DOM.ChildrenFiltered("a")
		if aTagChildren.Length() == 0 {
			return
		}
		newsLink := e.ChildAttr("a", "href")
		if newsLink == "" {
			return
		}

		// Extract published_on and slug from URL
		parts := strings.Split(strings.Trim(newsLink, "/"), "/")
		if len(parts) < 2 {
			return
		}
		publishedOn := parts[len(parts)-2]
		slug := parts[len(parts)-1]

		if seen[slug] {
			return
		}
		seen[slug] = true

		// Extract headline text
		headline := strings.Trim(e.Text, " \n")

		// Extract image (if exists)
		imageURL := ""
		if src := e.ChildAttr("img[src*='newsimage']", "src"); src != "" {
			if strings.HasPrefix(src, "http") {
				imageURL = src
			} else {
				imageURL = "https:" + src
			}
		}

		newsItem := model.News{
			Slug:        slug,
			Title:       headline,
			PublishedOn: utils.ConvertDateToISO(publishedOn),
			NewsUrl:     siteURL + newsLink,
			ImageUrl:    &imageURL,
		}
		newsList = append(newsList, newsItem)
		newsURLs = append(newsURLs, siteURL+newsLink)
	})

	// Handle errors
	c.OnError(func(_ *colly.Response, err error) {
		log.Println("Request failed:", err)
		return
	})

	// After scraping
	c.OnScraped(func(_ *colly.Response) {
		fmt.Println("Site Scraped Successfully!")
		fmt.Printf("\nTotal News Found: %d\n", len(newsList))
	})

	// Start
	fmt.Println("Visiting", siteURL)
	err := c.Visit(siteURL)
	if err != nil {
		return nil, err
	}

	allNewsBodies := ScrapeMultipleNews(newsURLs)
	for i := range newsList {
		newsList[i].Content = strings.Trim(allNewsBodies[i], " \n")
		repository.AddSingleNews(newsList[i])
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
