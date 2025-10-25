package service

import (
	"news-server/internal/model"
	"news-server/internal/repository"
)

type ScrapeService struct {
	repo *repository.NewsRepository
}

func NewScrapeService(repo *repository.NewsRepository) *ScrapeService {
	return &ScrapeService{repo: repo}
}

func (s *ScrapeService) ScrapeNews(limit int) ([]model.News, error) {
	// TODO: change implementation to scrape news instead of getting from DB
	return s.repo.GetAllNews(limit)
}
