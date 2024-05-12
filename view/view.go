package view

import (
	_ "embed"
	"html/template"
)

//go:embed index.go.html
var indexPage string

//go:embed article.go.html
var articlePage string

type Templates struct {
	Index   *template.Template
	Article *template.Template
}

func NewTemplates() *Templates {
	index := template.Must(template.New("index").Parse(indexPage))
	article := template.Must(template.New("article").Parse(articlePage))
	return &Templates{
		Index:   index,
		Article: article,
	}
}
