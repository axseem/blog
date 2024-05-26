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
	Summary string
	Date    time.Time
	Content template.HTML
}

func (a Article) ID() string {
	return strings.ToLower(
		strings.ReplaceAll(strings.TrimSpace(a.Title), " ", "_"),
	)
}

func ExtractFromFS(files fs.FS) ([]Article, error) {
	md := goldmark.New(goldmark.WithExtensions(extension.GFM, meta.Meta))

	var articles []Article
	err := fs.WalkDir(files, ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() || filepath.Ext(path) != ".md" {
			return nil
		}

		articleRaw, err := fs.ReadFile(files, path)
		if err != nil {
			return err
		}
		article, err := markdownToArticle(md, &articleRaw)
		if err != nil {
			return err
		}
		articles = append(articles, article)

		return nil
	})
	if err != nil {
		return nil, err
	}

	return articles, nil
}

func markdownToArticle(md goldmark.Markdown, file *[]byte) (Article, error) {
	buf, metaData, err := parseMarkdown(md, file)
	if err != nil {
		return Article{}, err
	}
	keys := [...]string{"title", "date", "summary"}
	if err := validateMetadata(metaData, keys[:]); err != nil {
		return Article{}, err
	}

	title, _ := metaData["title"].(string)
	summary, _ := metaData["summary"].(string)
	dateRaw, _ := metaData["date"].(string)
	date, err := time.Parse("02.01.2006", dateRaw)
	if err != nil {
		return Article{}, err
	}

	return Article{
		Title:   title,
		Summary: summary,
		Date:    date,
		Content: template.HTML(buf.String()),
	}, nil
}

func parseMarkdown(md goldmark.Markdown, file *[]byte) (bytes.Buffer, map[string]interface{}, error) {
	var buf bytes.Buffer
	context := parser.NewContext()
	if err := md.Convert(*file, &buf, parser.WithContext(context)); err != nil {
		return bytes.Buffer{}, nil, err
	}
	metaData := meta.Get(context)

	return buf, metaData, nil
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
