package database

import (
	"blomple/model"
	"strings"
)

func (s *Storage) ArticleList() ([]model.Article, error) {
	const query = "SELECT id, title, created_at FROM article"
	rows, err := s.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer deferErr(rows.Close, &err)

	var articles []model.Article
	for rows.Next() {
		var a model.Article
		_ = rows.Scan(
			&a.ID,
			&a.Title,
			&a.CreatedAt,
		)
		articles = append(articles, a)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return articles, nil
}

func (s *Storage) ArticleCreate(article *model.Article) error {
	const query = "INSERT INTO article (id, title, content) VALUES (?, ?, ?)"
	_, err := s.db.Exec(query,
		titleToID(article.Title),
		article.Title,
		article.Content,
	)

	return err
}

func titleToID(title string) string {
	return strings.ToLower(
		strings.ReplaceAll(strings.TrimSpace(title), " ", "_"),
	)
}

func (s *Storage) ArticleFind(id string) (model.Article, error) {
	const qurey = "SELECT * FROM article WHERE id = ?"
	row := s.db.QueryRow(qurey, id)

	var a model.Article
	err := row.Scan(
		&a.ID,
		&a.Title,
		&a.Content,
		&a.CreatedAt,
	)

	return a, err
}
