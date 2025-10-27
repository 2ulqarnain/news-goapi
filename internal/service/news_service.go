package service

import (
	"context"
	"news-server/internal/model"
	"news-server/internal/repository"
)

type NewsService interface {
	GetAllNews(ctx context.Context) ([]model.News, error)
	AddSingleNews(ctx context.Context, news model.News) error
	AddMultipleNews(ctx context.Context, news []model.News) error
}

type newsService struct {
	repo repository.NewsRepository
}

func NewNewsService(repo repository.NewsRepository) NewsService {
	return &newsService{repo: repo}
}

func (s *newsService) AddSingleNews(ctx context.Context, news model.News) error {
	return s.repo.AddSingleNews(ctx, news)
}

func (s *newsService) AddMultipleNews(ctx context.Context, news []model.News) error {
	return s.repo.AddMultipleNews(ctx, news)
}

func (s *newsService) GetAllNews(ctx context.Context) ([]model.News, error) {
	return s.repo.GetAllNews(ctx)
}
