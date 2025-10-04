package scrapers

import (
	"fmt"
	"log"
	"strings"
	"testing"

	"github.com/gocolly/colly"
	"github.com/yourusername/news-server/internal/model"
)

func TestScrapeNews(t *testing.T) {
	url := "https://www.radio.gov.pk/25-09-2025/pakistan-committed-to-build-constructive-ties-with-bangladesh-pm"
	news := ScrapeRadioPkNewsByURL(url)
	if news == "" {
		t.Error("Could not scrape!")
	} else {
		t.Log("Scraped Successfully: ", news[:100])
	}
}

func TestScrapeMultipleNews(t *testing.T) {
	failed := true
	urls := []string{
		"https://www.radio.gov.pk/25-09-2025/pm-shehbaz-to-meet-president-trump-in-washington-today",
		"https://www.radio.gov.pk/25-09-2025/pakistan-bahrain-agree-to-enhance-engagement-across-diverse-sectors",
		"https://www.radio.gov.pk/25-09-2025/iranian-president-welcomes-defence-pact-bw-saudi-arabia-pakistan",
	}

	newsList := ScrapeMultipleNews(urls)

	for _, news := range newsList {
		if news != "" {
			failed = false
			break
		}
	}

	if failed {
		t.Error("Test failed!")
	} else {
		t.Log("Scraped Successfully: ", newsList)
	}
}

func TestRadioPkHomePageScraping(t *testing.T) {
	var newsList []model.News

	// Create collector
	c := colly.NewCollector(
		colly.AllowedDomains("radio.gov.pk"),
	)

	//seen := make(map[string]bool)
	// Find all links that have a child with class containing 'title'
	//c.OnHTML("a:has([class*='title']), a:has(.multipanelimage):has([class*='title']), div.col-lg-6.col-md-6 a:has([class*='title'])", func(e *colly.HTMLElement) {
	//	newsLink := e.Attr("href")
	//	if newsLink == "" {
	//		return
	//	}
	//
	//	// Extract published_on and slug from URL
	//	parts := strings.Split(strings.Trim(newsLink, "/"), "/")
	//	if len(parts) < 2 {
	//		return
	//	}
	//	publishedOn := parts[len(parts)-2]
	//	slug := parts[len(parts)-1]
	//
	//	if seen[slug] {
	//		return
	//	}
	//	seen[slug] = true
	//
	//	// Extract headline text
	//	headline := e.ChildText("[class*='title']")
	//
	//	// Extract image (if exists)
	//	imageURL := ""
	//	if src := e.ChildAttr("img", "src"); src != "" {
	//		if strings.HasPrefix(src, "http") {
	//			imageURL = src
	//		} else {
	//			imageURL = "https:" + src
	//		}
	//	}
	//
	//	newsItem := model.News{
	//		Slug:        slug,
	//		Title:       headline,
	//		PublishedOn: utils.ConvertDateToISO(publishedOn),
	//		NewsUrl:     siteURL + newsLink,
	//		ImageUrl:    &imageURL,
	//	}
	//	newsList = append(newsList, newsItem)
	//	newsURLs = append(newsURLs, siteURL+newsLink)
	//})

	c.OnHTML("div.col-lg-6.col-md-6", func(e *colly.HTMLElement) {
		// Get image src and normalize it
		imgSrc := e.ChildAttr("img", "src")
		if strings.HasPrefix(imgSrc, "//") {
			imgSrc = "https:" + imgSrc
		}

		// Get panel title text (handles different placements)
		title := e.ChildText(".paneltitle")

		// Only print if we got something
		if imgSrc != "" && title != "" {
			fmt.Println("Image:", imgSrc)
			fmt.Println("Title:", title)
			fmt.Println("-----")
		}
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
		t.Errorf("Visiting failed: %s", err)
	}

	t.Logf("\nTotal News Found: %d\n", len(newsList))
}
