package view

import (
	"blog/article"
	"bytes"
	"embed"
	_ "embed"
	"html/template"

	"github.com/yuin/goldmark"
	meta "github.com/yuin/goldmark-meta"
	"github.com/yuin/goldmark/extension"
)

//go:embed *.html
var templateFiles embed.FS

type View struct {
	Template *template.Template
	Markdown goldmark.Markdown
}

func New() *View {
	md := goldmark.New(
		goldmark.WithExtensions(extension.GFM, meta.Meta),
	)

	return &View{
		Template: template.Must(template.ParseFS(templateFiles, "*")),
		Markdown: md,
	}
}

func (v View) GenerateIndexPage(articles *[]article.Article) ([]byte, error) {
	var buf bytes.Buffer

	if err := v.Template.ExecuteTemplate(&buf, "index.html", articles); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func (v View) GenerateArticles(articles *[]article.Article) (map[string][]byte, error) {
	pages := make(map[string][]byte)

	for _, article := range *articles {
		var buf bytes.Buffer
		if err := v.Template.ExecuteTemplate(&buf, "article.html", article); err != nil {
			return nil, err
		}
		pages[article.ID()] = buf.Bytes()
	}

	return pages, nil
}
