package tests

import (
	"fmt"
	"news-server/scrapers"
	"testing"
)

func TestScrapeProPakistani_Pk(t *testing.T) {
	fmt.Println("TestScrapeProPakistani_Pk | Scraping https://propakistani.pk")
	news, err := scrapers.ScrapeProPakistaniPk()
	if err != nil {
		t.Error(err)
		return
	} else if len(news) == 0 {
		t.Log("No news found!")
		return
	}
	t.Logf("%d News Found!", len(news))
	for _, newsItem := range news {
		fmt.Printf("News: %v\n", newsItem)
	}
}
