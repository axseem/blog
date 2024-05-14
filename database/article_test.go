package database_test

import (
	"blog/database"
	"blog/model"
	"blog/test"
	"reflect"
	"testing"
)

// depends on ArticleCreate method
func TestArticleList(t *testing.T) {
	testCases := []struct {
		desc string
		prep func(s *database.Storage) error
		cond func(articles []model.Article) bool
	}{
		{
			desc: "empty table",
			prep: func(s *database.Storage) error { return nil },
			cond: func(articles []model.Article) bool {
				return len(articles) == 0
			},
		},
		{
			desc: "one article",
			prep: func(s *database.Storage) error {
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
			prep: func(s *database.Storage) error {
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
			assert := test.NewAssert(t)
			db := database.MustOpen(":memory:")
			assert.Nil(database.Migrate(db))
			defer assert.NilDefer(db.Close)
			s := database.NewStorage(db)
			assert.Nil(tC.prep(s))

			articles, err := s.ArticleList()
			assert.Nil(err)
			assert.True(tC.cond(articles), tC.desc)
		})
	}
}

func TestArticleCreate(t *testing.T) {
	assert := test.NewAssert(t)
	db := database.MustOpen(":memory:")
	assert.Nil(database.Migrate(db))
	defer assert.NilDefer(db.Close)
	s := database.NewStorage(db)

	assert.Nil(s.ArticleCreate(&model.Article{
		Title:   "A",
		Content: "Aaa",
	}))
	want := &model.Article{
		ID:      "a",
		Title:   "A",
		Content: "Aaa",
	}

	var got model.Article

	const query = "SELECT * FROM article WHERE id = ?"
	assert.Nil(db.QueryRow(query, want.ID).Scan(
		&got.ID,
		&got.Title,
		&got.Content,
		&got.CreatedAt,
	))

	isValid := got.ID == want.ID && got.Title == want.Title && got.Content == want.Content
	assert.True(isValid, "article create test")
}

func TestArticleFind(t *testing.T) {
	assert := test.NewAssert(t)
	db := database.MustOpen(":memory:")
	assert.Nil(database.Migrate(db))
	defer assert.NilDefer(db.Close)
	s := database.NewStorage(db)

	want := &model.Article{
		ID:      "a",
		Title:   "A",
		Content: "Aaa",
	}

	const query = "INSERT INTO article (id, title, content) VALUES (?, ?, ?)"
	_, err := db.Exec(query,
		want.ID,
		want.Title,
		want.Content,
	)
	assert.Nil(err)

	got, err := s.ArticleFind(want.ID)
	assert.Nil(err)

	isValid := got.ID == want.ID && got.Title == want.Title && got.Content == want.Content
	assert.True(isValid, "article find test")
}
