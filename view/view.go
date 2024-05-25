package view

import (
	"bytes"
	"embed"
	_ "embed"
	"html/template"

	"github.com/axseem/website/article"
)

//go:embed *.html
var templateFiles embed.FS

type View struct {
	Template *template.Template
}

func New() *View {
	return &View{
		Template: template.Must(template.ParseFS(templateFiles, "*")),
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
