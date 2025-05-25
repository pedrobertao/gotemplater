package generator

import (
	"embed"
)

//go:embed templates/**/structure.json
var Templates embed.FS
