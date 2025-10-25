package service

import (
	"news-server/internal/model"
	"news-server/internal/repository"
)

type NewsService struct {
	repo *repository.NewsRepository
}

func NewNewsService(repo *repository.NewsRepository) *NewsService {
	return &NewsService{repo: repo}
}

func (s *NewsService) GetAllNews(limit int) ([]model.News, error) {
	return s.repo.GetAllNews(limit)
}

func (s *NewsService) GetNewsBySlug(slug string) {
	return
}
