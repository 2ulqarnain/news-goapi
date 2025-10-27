package service

import (
	"news-server/internal/repository"
	"news-server/scrapers"
)

type ScrapeService interface {
	ScrapeRadioGovPk() error
	ScrapePropakistaniPk() error
}

type scrapeService struct {
	repo repository.NewsRepository
}

func (s *scrapeService) ScrapeRadioGovPk() error {
	_, err := scrapers.ScrapeRadioPakistan()
	if err != nil {
		return err
	}
	return nil
}

func (s *scrapeService) ScrapePropakistaniPk() error {
	_, err := scrapers.ScrapeProPakistaniPk()
	if err != nil {
		return err
	}
	return nil
}
