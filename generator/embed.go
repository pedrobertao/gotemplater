package generator

import (
	"embed"
)

//go:embed templates/**/structure.yaml
var Templates embed.FS
