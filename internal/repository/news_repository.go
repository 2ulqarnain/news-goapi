package repository

import (
	"context"
	"database/sql"
	"news-server/internal/model"
)

type NewsRepository interface {
	AddSingleNews(news model.News) error
	AddMultipleNews(ctx context.Context, news []model.News) error
	GetAllNews(ctx context.Context) ([]model.News, error)
	SearchNewsBySlug(ctx context.Context, slug string) ([]model.News, error)
}

type newsRepository struct {
	db *sql.DB
}

func NewNewsRepository(db *sql.DB) NewsRepository {
	return &newsRepository{db: db}
}

func (r *newsRepository) AddSingleNews(news model.News) error {
	query := "INSERT INTO news (slug,title,published_on,news_url,image_url,content,source) VALUES (?,?,?,?,?,?,?)"
	_, err := r.db.Exec(query, news.Slug, news.Title, news.PublishedOn, news.NewsUrl, news.Content, news.Source)
	if err != nil {
		return err
	}
	return nil
}

func (r *newsRepository) AddMultipleNews(ctx context.Context, news []model.News) error {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	stmt, err := tx.PrepareContext(ctx, "INSERT INTO news (slug,title,published_on,news_url,image_url,content,source) VALUES (?,?,?,?,?,?,?)")
	if err != nil {
		return err
	}
	defer stmt.Close()
	for _, n := range news {
		if _, err := stmt.ExecContext(ctx, n.Slug, n.Title, n.PublishedOn, n.NewsUrl, n.ImageUrl, n.Content, n.Source); err != nil {
			return err
		}
	}
	if err := tx.Commit(); err != nil {
		return err
	}
	return nil
}

func (r *newsRepository) GetAllNews(ctx context.Context) ([]model.News, error) {
	query := "SELECT slug, title, published_on, image_url, news_url, source FROM news;"
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var newsList []model.News
	for rows.Next() {
		var n model.News
		if err := rows.Scan(&n.Slug, &n.Title, &n.PublishedOn, &n.ImageUrl, &n.NewsUrl, &n.Source); err != nil {
			return nil, err
		}
		newsList = append(newsList, n)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return newsList, nil
}

func (r *newsRepository) SearchNewsBySlug(ctx context.Context, slug string) ([]model.News, error) {
	query := "SELECT slug, title, source FROM news WHERE slug LIKE ?"
	rows, err := r.db.QueryContext(ctx, query, "%"+slug+"%")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var newsList []model.News
	for rows.Next() {
		var n model.News
		if err := rows.Scan(&n.Slug, &n.Title, &n.Source); err != nil {
			return nil, err
		}
		newsList = append(newsList, n)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return newsList, nil
}
