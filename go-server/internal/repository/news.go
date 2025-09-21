package repository

import (
	"log"

	"github.com/yourusername/news-server/internal/model"
)

func GetAllNews() ([]model.News, error) {
	rows, err := db.Query("SELECT * from news;")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var news []model.News
	for rows.Next() {
		var n model.News
		err := rows.Scan(&n.Slug, &n.Headline, &n.PublishedOn, &n.NewsUrl, &n.ImageUrl)
		if err != nil {
			log.Printf("Couldn't load data from database!, Error: %v", err)
			return nil, err
		}
		news = append(news, n)
	}
	return news, rows.Err()
}
