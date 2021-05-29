package assets

import (
	"embed"
)

//go:embed Talkie/dist/*.css
//go:embed Talkie/dist/*.js
//go:embed Talkie/dist/*.js.map
var Talkie embed.FS

//go:embed template/index.html
var IndexHTMLTemplate []byte
