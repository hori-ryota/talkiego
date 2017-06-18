package assets

import (
	"time"

	"github.com/jessevdk/go-assets"
)

var _Template334fce7394182881cdad66f000b8c4ad9c7cb38f = "<html>\n<head>\n  <meta charset=\"utf-8\">\n  <title>Talkie.js - HTML/CSS/JavaScript Presentation Library</title>\n  <link rel=\"stylesheet\" href=\"./talkie/talkie.min.css\">\n  <link rel=\"stylesheet\" href=\"./talkie/talkie-default.min.css\">\n  <link rel=\"stylesheet\" href=\"https://cdnjs.cloudflare.com/ajax/libs/highlight.js/8.9.1/styles/monokai_sublime.min.css\">\n</head>\n<body>\n\n<!-- put your slides -->\n{{range .Pages}}\n<template layout=\"{{.Option.Layout}}\" type=\"text/x-markdown\"\n{{with .Option.Backface}} \n          backface=\"{{.}}\"\n{{end}}\n{{with .Option.BackfaceFilter}} \n          backface-filter=\"{{.}}\"\n{{end}}\n{{if .Option.Invert}}\n          invert\n{{end}}\n{{with .Option.Style}}\n          style=\"{{.}}\"\n{{end}}\n>\n{{.Body}}\n</template>\n{{end}}\n\n<script src=\"//cdnjs.cloudflare.com/ajax/libs/highlight.js/8.9.1/highlight.min.js\"></script>\n<script src=\"./talkie/talkie.min.js\"></script>\n<script>\n  var talkie = Talkie({\n    {{with .RootOption.Wide}}wide: true,{{end}}\n    {{with .RootOption.Control}}control: true,{{end}}\n    {{with .RootOption.Pointer}}pointer: true,{{end}}\n    {{with .RootOption.Progress}}progress: true,{{end}}\n    {{with .RootOption.Backface}}backface: true,{{end}}\n    {{with .RootOption.NoTransition}}notransition: true,{{end}}\n    {{with .RootOption.LinkShouldBlank}}linkshouldblank: true,{{end}}\n  });\n  talkie.key('s').subscribe(talkie.next$);\n  talkie.key('n').subscribe(talkie.next$);\n  talkie.key('a').subscribe(talkie.prev$);\n  talkie.key('p').subscribe(talkie.prev$);\n  document.addEventListener('DOMContentLoaded', function() {\n    talkie.changed.subscribe(function(current) {\n      console.clear();\n      // print presenter notes\n      console.info(talkie.notes[current.getAttribute('data-page')]);\n    });\n  });\n</script>\n</body>\n</html>\n"

// Template returns go-assets FileSystem
var Template = assets.NewFileSystem(map[string][]string{"/": []string{"index.html"}}, map[string]*assets.File{
	"/": &assets.File{
		Path:     "/",
		FileMode: 0x800001ed,
		Mtime:    time.Unix(1497803067, 1497803067000000000),
		Data:     nil,
	}, "/index.html": &assets.File{
		Path:     "/index.html",
		FileMode: 0x1a4,
		Mtime:    time.Unix(1497800902, 1497800902000000000),
		Data:     []byte(_Template334fce7394182881cdad66f000b8c4ad9c7cb38f),
	}}, "")
