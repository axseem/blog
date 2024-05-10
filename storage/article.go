package storage

import (
	"blomple/article"
	"log"
)

func (s *Storage) ArticleList() ([]article.Article, error) {
	const query = "SELECT * FROM article"
	rows, err := s.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer func() {
		deferErr := rows.Close()
		if err != nil {
			if deferErr != nil {
				log.Printf("ERROR: %v\n", err)
			}
			return
		}
		err = deferErr
	}()

	var articles []article.Article
	for rows.Next() {
		var a article.Article
		rows.Scan(a)
		articles = append(articles, a)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return articles, nil
}
