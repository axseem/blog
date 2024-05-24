package static

import (
	"embed"
)

//go:embed assets blog
var static embed.FS

func Static() embed.FS {
	return static
}
