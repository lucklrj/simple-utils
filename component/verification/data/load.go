package data

import (
	"embed"
	_ "embed"
)

//go:embed *
var StaticFile embed.FS
