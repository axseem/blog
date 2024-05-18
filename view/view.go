package view

import (
	"embed"
	_ "embed"
	"html/template"

	"github.com/yuin/goldmark"
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
		goldmark.WithExtensions(extension.GFM),
	)

	return &View{
		Template: template.Must(template.ParseFS(templateFiles, "*")),
		Markdown: md,
	}
}
