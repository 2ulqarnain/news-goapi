package scrapers

import (
	"testing"
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
