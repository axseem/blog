package static

import (
	"embed"
)

// //go:embed assets
// var assets embed.FS

// func Assets() embed.FS {
// 	return assets
// }

//go:embed *.md
var articles embed.FS

func Articles() embed.FS {
	return articles
}
