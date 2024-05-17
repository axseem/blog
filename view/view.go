package view

import (
	"embed"
	_ "embed"
	"html/template"

	"github.com/gomarkdown/markdown/html"
	"github.com/gomarkdown/markdown/parser"
)

//go:embed *.html
var templateFiles embed.FS

type View struct {
	Template *template.Template
	Parser   *parser.Parser
	Renderer *html.Renderer
}

func New() *View {
	extensions := parser.CommonExtensions | parser.AutoHeadingIDs | parser.NoEmptyLineBeforeBlock
	opts := html.RendererOptions{Flags: html.CommonFlags | html.HrefTargetBlank}

	return &View{
		Template: template.Must(template.ParseFS(templateFiles, "*")),
		Parser:   parser.NewWithExtensions(extensions),
		Renderer: html.NewRenderer(opts),
	}
}
