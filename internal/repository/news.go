package repository

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/yourusername/news-server/internal/model"
)

func GetAllNews(limit int) ([]model.News, error) {
	var (
		dbQuery string
		rows    *sql.Rows
		err     error
	)
	if limit > 0 {
		dbQuery = "SELECT * FROM news LIMIT ?"
		rows, err = db.Query(dbQuery, limit)
	} else {
		dbQuery = "SELECT * FROM news"
		rows, err = db.Query(dbQuery)
	}
	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {

		}
	}(rows)

	var news []model.News
	for rows.Next() {
		var n model.News
		err := rows.Scan(&n.Slug, &n.Title, &n.PublishedOn, &n.NewsUrl, &n.ImageUrl, &n.Content)
		if err != nil {
			log.Printf("Couldn't load data from database!, Error: %v", err)
			return nil, err
		}
		news = append(news, n)
	}
	return news, rows.Err()
}

func GetNewsBySlug(slug string) ([]model.News, error) {
	rows, err := db.Query("SELECT * from news;")
	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {

		}
	}(rows)

	var news []model.News
	for rows.Next() {
		var n model.News
		err := rows.Scan(&n.Slug, &n.Title, &n.PublishedOn, &n.NewsUrl, &n.ImageUrl, &n.Content)
		if err != nil {
			log.Printf("Couldn't load data from database!, Error: %v", err)
			return nil, err
		}
		news = append(news, n)
	}
	return news, rows.Err()
}

//func AddMultipleNews(newsList []model.News) {
//	tx, err := db.Begin()
//	if err != nil {
//		fmt.Println("Failed to begin transaction:", err)
//		return
//	}
//
//	stmt, err := tx.Prepare("INSERT INTO news (slug, title, published_on, news_url, image_url, content) VALUES (?,?,?,?,?,?)")
//	if err != nil {
//		fmt.Println("Failed to prepare statement:", err)
//		return
//	}
//	defer func(stmt *sql.Stmt) {
//		err := stmt.Close()
//		if err != nil {
//
//		}
//	}(stmt)
//
//	for _, news := range newsList {
//		_, err := stmt.Exec(news.Slug, news.Title, news.PublishedOn, news.NewsUrl, news.ImageUrl, news.Content)
//		if err != nil {
//			fmt.Printf("Failed to insert slug %s: %v\n", news.Slug, err)
//			// Just skip this one, continue with others
//			continue
//		}
//	}
//
//	if err := tx.Commit(); err != nil {
//		fmt.Println("Failed to commit transaction:", err)
//		return
//	}
//
//	fmt.Println("Inserted all valid news successfully!")
//}

func AddSingleNews(news model.News) {
	tx, err := db.Begin()
	if err != nil {
		fmt.Println("Failed to begin transaction:", err)
		return
	}

	stmt, err := tx.Prepare("INSERT INTO news (slug, title, published_on, news_url, image_url, content) VALUES (?,?,?,?,?,?)")
	if err != nil {
		fmt.Println("Failed to prepare statement:", err)
		return
	}
	defer func(stmt *sql.Stmt) {
		err := stmt.Close()
		if err != nil {

		}
	}(stmt)

	_, err = stmt.Exec(news.Slug, news.Title, news.PublishedOn, news.NewsUrl, news.ImageUrl, news.Content)
	if err != nil {
		fmt.Printf("Failed to insert slug %s: %v\n", news.Slug, err)
	}

	if err := tx.Commit(); err != nil {
		fmt.Println("Failed to commit transaction:", err)
		return
	}
}
