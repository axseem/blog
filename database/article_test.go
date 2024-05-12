package database_test

import (
	"blomple/database"
	"blomple/model"
	"blomple/test"
	"reflect"
	"testing"
)

// depends on ArticleCreate method
func TestArticleList(t *testing.T) {
	testCases := []struct {
		desc string
		seed func(s *database.Storage) error
		cond func(articles []model.Article) bool
	}{
		{
			desc: "empty table",
			seed: func(s *database.Storage) error { return nil },
			cond: func(articles []model.Article) bool {
				return len(articles) == 0
			},
		},
		{
			desc: "one article",
			seed: func(s *database.Storage) error {
				return s.ArticleCreate(&model.Article{Title: "A", Content: "Aaa"})
			},
			cond: func(articles []model.Article) bool {
				expect := []model.Article{{
					ID:        "a",
					Title:     "A",
					Content:   "",
					CreatedAt: articles[0].CreatedAt,
				}}
				return reflect.DeepEqual(articles, expect)
			},
		},
		{
			desc: "several articles",
			seed: func(s *database.Storage) error {
				articles := []model.Article{
					{Title: "A", Content: "Aaa"},
					{Title: "B", Content: "Bbb"},
					{Title: "C", Content: "Ccc"},
				}
				for _, v := range articles {
					if err := s.ArticleCreate(&v); err != nil {
						return err
					}
				}
				return nil
			},
			cond: func(articles []model.Article) bool {
				expect := []model.Article{
					{
						ID:        "a",
						Title:     "A",
						Content:   "",
						CreatedAt: articles[0].CreatedAt,
					},
					{
						ID:        "b",
						Title:     "B",
						Content:   "",
						CreatedAt: articles[1].CreatedAt,
					},
					{
						ID:        "c",
						Title:     "C",
						Content:   "",
						CreatedAt: articles[2].CreatedAt,
					},
				}
				return reflect.DeepEqual(articles, expect)
			},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			assert := test.New(t)

			db := database.MustOpen(":memory:")
			assert.Nil(database.Migrate(db))
			defer assert.NilDefer(db.Close)
			s := database.NewStorage(db)
			assert.Nil(tC.seed(s))

			articles, err := s.ArticleList()
			assert.Nil(err)

			assert.True(tC.cond(articles), tC.desc)
		})
	}
}
