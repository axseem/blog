package static

import (
	"embed"
)

//go:embed *
var static embed.FS

func Static() embed.FS {
	return static
}

//go:embed article
var articles embed.FS

func Articles() embed.FS {
	return articles
}
