package template

import (
	"embed"
)

//go:embed dmloader.js game.js game.json project.config.json project.private.config.json
//go:embed all:weapp-adapter
var Files embed.FS
