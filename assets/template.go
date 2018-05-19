package assets

import (
	"time"

	"github.com/jessevdk/go-assets"
)

var _Template334fce7394182881cdad66f000b8c4ad9c7cb38f = "<html>\n<head>\n  <meta charset=\"utf-8\">\n  <title>Talkie.js - HTML/CSS/JavaScript Presentation Library</title>\n  <link rel=\"stylesheet\" href=\"./assets/Talkie/dist/talkie.css\">\n  <link rel=\"stylesheet\" href=\"./assets/Talkie/dist/talkie.theme-default.css\">\n  <link rel=\"stylesheet\" href=\"https://cdnjs.cloudflare.com/ajax/libs/highlight.js/9.12.0/styles/monokai-sublime.min.css\">\n  {{range .CustomCSSs}}\n  <link rel=\"stylesheet\" href=\"{{.}}\">\n  {{end}}\n</head>\n<body>\n\n<!-- put your slides -->\n{{range .Pages}}\n{{if .Option.Script}}\n<script\n{{else}}\n<tk-slide\n{{end}}\n          layout=\"{{.Option.Layout}}\" type=\"text/x-markdown\"\n{{with .Option.Backface}} \n          backface=\"{{.}}\"\n{{end}}\n{{with .Option.BackfaceFilter}} \n          backface-filter=\"{{.}}\"\n{{end}}\n{{if .Option.Invert}}\n          invert\n{{end}}\n{{with .Option.Style}}\n          style=\"{{.}}\"\n{{end}}\n>\n{{.Body}}\n{{if .Option.Script}}\n</script>\n{{else}}\n</tk-slide>\n{{end}}\n{{end}}\n\n<tk-pager></tk-pager>\n<tk-progress></tk-progress>\n\n<script src=\"//cdnjs.cloudflare.com/ajax/libs/highlight.js/9.12.0/highlight.min.js\"></script>\n<script src=\"//cdnjs.cloudflare.com/ajax/libs/highlight.js/9.12.0/languages/dockerfile.min.js\"></script>\n<script src=\"//cdnjs.cloudflare.com/ajax/libs/highlight.js/9.12.0/languages/go.min.js\"></script>\n<script src=\"//cdnjs.cloudflare.com/ajax/libs/highlight.js/9.12.0/languages/typescript.min.js\"></script>\n<script src=\"//cdnjs.cloudflare.com/ajax/libs/highlight.js/9.12.0/languages/vim.min.js\"></script>\n<script src=\"//cdnjs.cloudflare.com/ajax/libs/highlight.js/9.12.0/languages/yaml.min.js\"></script>\n<script src=\"./assets/Talkie/dist/webcomponents-loader.js\"></script>\n\n<script src=\"./assets/Talkie/dist/talkie.js\"></script>\n{{range .CustomScripts}}\n<script src=\"{{.}}\"></script>\n{{end}}\n<script>\nwindow.addEventListener('WebComponentsReady', function(e) {\n  document.body.className += ' webcomponents-ready';\n  talkie.run({\n    wide: {{.RootOption.Wide}},\n    control: {{.RootOption.Control}},\n    pointer: {{.RootOption.Pointer}},\n    progress: {{.RootOption.Progress}},\n    backface: {{.RootOption.Backface}},\n    notransition: {{.RootOption.NoTransition}},\n    linkshouldblank: {{.RootOption.LinkShouldBlank}},\n  });\n});\n</script>\n</body>\n</html>\n"

// Template returns go-assets FileSystem
var Template = assets.NewFileSystem(map[string][]string{"/": []string{"index.html"}}, map[string]*assets.File{
	"/": &assets.File{
		Path:     "/",
		FileMode: 0x800001ed,
		Mtime:    time.Unix(1526772502, 1526772502260756925),
		Data:     nil,
	}, "/index.html": &assets.File{
		Path:     "/index.html",
		FileMode: 0x1a4,
		Mtime:    time.Unix(1526772502, 1526772502260844534),
		Data:     []byte(_Template334fce7394182881cdad66f000b8c4ad9c7cb38f),
	}}, "")
