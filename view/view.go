package view

import (
	_ "embed"
	"html/template"
)

//go:embed index.go.html
var indexPage string

type Templates struct {
	Index *template.Template
}

func NewTemplates() *Templates {
	index := template.Must(template.New("index").Parse(indexPage))
	return &Templates{
		Index: index,
	}
}
