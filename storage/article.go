package storage

import "blomple/article"

func (s *Storage) ArticleList() ([]article.Article, error) {
	const query = "SELECT * FROM article"
	rows, err := s.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer deferErr(rows.Close, &err)

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
