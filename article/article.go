package article

import (
	"bytes"
	"errors"
	"html/template"
	"io/fs"
	"path/filepath"
	"strings"
	"time"

	"github.com/yuin/goldmark"
	meta "github.com/yuin/goldmark-meta"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/parser"
)

type Article struct {
	Title   string
	Date    time.Time
	Summary string
	Content template.HTML
}

func (a Article) ID() string {
	return strings.ToLower(
		strings.ReplaceAll(strings.TrimSpace(a.Title), " ", "_"),
	)
}

func ExtractFromFS(files fs.FS) ([]Article, error) {
	md := goldmark.New(
		goldmark.WithExtensions(extension.GFM, meta.Meta),
	)

	var articles []Article
	err := fs.WalkDir(files, ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			return nil
		}
		if filepath.Ext(path) != ".md" {
			return nil
		}

		articleRaw, err := fs.ReadFile(files, path)
		if err != nil {
			return err
		}

		var buf bytes.Buffer
		context := parser.NewContext()
		if err := md.Convert(articleRaw, &buf, parser.WithContext(context)); err != nil {
			return err
		}
		metaData := meta.Get(context)
		keys := [...]string{"title", "date", "summary"}
		if err := validateMetadata(metaData, keys[:]); err != nil {
			return err
		}

		title, _ := metaData["title"].(string)
		dateRaw, _ := metaData["date"].(string)
		date, err := time.Parse("02.01.2006", dateRaw)
		if err != nil {
			return err
		}
		summary, _ := metaData["summary"].(string)

		articles = append(articles, Article{
			Title:   title,
			Date:    date,
			Summary: summary,
			Content: template.HTML(buf.String()),
		})

		return nil
	})
	if err != nil {
		return nil, err
	}

	return articles, nil
}

func validateMetadata(m map[string]interface{}, keys []string) error {
	if len(m) != len(keys) {
		return errors.New("metadata tags length is not equals keys length")
	}
	for _, k := range keys {
		_, ok := m[k]
		if !ok {
			return errors.New("no key in metadata tags: " + k)
		}
	}
	return nil
}
