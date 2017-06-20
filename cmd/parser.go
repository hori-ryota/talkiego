package cmd

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

type RootOption struct {
	Wide            bool
	Control         bool
	Pointer         bool
	Progress        bool
	Backface        bool
	NoTransition    bool
	LinkShouldBlank bool
}

type PageOption struct {
	Layout         string
	Backface       string
	BackfaceFilter string
	Invert         bool
	Style          string
	Script         bool
}

type Page struct {
	Option PageOption
	Body   string
}

type TalkieParam struct {
	RootOption RootOption
	Pages      []Page
}

var rootOptReg = regexp.MustCompile(`(?s)<!--\s*rootopt:(.*?)-->`)

func (p *TalkieParam) Parse(src string, divider string) error {
	// rootopt
	matcher := rootOptReg.FindAllStringSubmatch(src, -1)
	for _, m := range matcher {
		props := parseProps(m[1])
		for k, v := range props {
			if err := p.RootOption.Apply(k, v); err != nil {
				return errors.Wrap(err, "RootOption Apply failed")
			}
		}
	}

	divider = fmt.Sprintf("\n%s\n", divider)
	pageStrs := strings.Split(src, divider)

	pages := make([]Page, len(pageStrs))
	for i, s := range pageStrs {
		if err := pages[i].Parse(s); err != nil {
			return errors.Wrap(err, "Page.Parse failed")
		}
	}

	p.Pages = pages

	return nil
}

func (p *RootOption) Apply(key, value string) error {
	if value == "" {
		// default true
		value = "true"
	}
	v, err := strconv.ParseBool(value)
	if err != nil {
		return errors.Wrapf(err, "%s is not bool [key='%s', value='%s']", value, key, value)
	}

	switch strings.ToLower(key) {
	case "wide":
		p.Wide = v
	case "control":
		p.Control = v
	case "pointer":
		p.Pointer = v
	case "progress":
		p.Progress = v
	case "backface":
		p.Backface = v
	case "notransition":
		p.NoTransition = v
	case "linkshouldblank":
		p.LinkShouldBlank = v
	default:
		return errors.Errorf("%s is unknown key [key='%s', value='%s']", key, key, value)
	}

	return nil
}

func (p *Page) Parse(src string) error {
	if err := p.Option.Parse(src); err != nil {
		return errors.Wrap(err, "PageOption.Parse failed")
	}
	p.Body = src
	return nil
}

var pageOptReg = regexp.MustCompile(`(?s)<!--\s*([a-zA-Z]+):(.*?)-->`)

func (p *PageOption) Parse(src string) error {
	matcher := pageOptReg.FindAllStringSubmatch(src, -1)
	for _, m := range matcher {
		if m[1] == "rootopt" {
			continue
		}
		p.Layout = m[1]
		if len(m) == 2 {
			continue
		}
		props := parseProps(m[2])
		for k, v := range props {
			if err := p.Apply(k, v); err != nil {
				return errors.Wrap(err, "PageOption Apply failed")
			}
		}
	}

	return nil
}

func (p *PageOption) Apply(key, value string) error {
	switch strings.ToLower(key) {
	case "backface":
		p.Backface = value
	case "backfacefilter":
		p.BackfaceFilter = value
	case "invert":
		p.Invert = true
	case "style":
		p.Style = value
	case "script":
		p.Script = true
	default:
		return errors.Errorf("%s is unknown key [key='%s', value='%s']", key, key, value)
	}
	return nil
}

func parseProps(s string) map[string]string {
	dst := map[string]string{}
	fs := strings.Fields(s)

	for i := 0; i < len(fs); i++ {
		// split key:value or key=value
		sep := strings.IndexAny(fs[i], ":=")
		if sep < 0 {
			dst[fs[i]] = ""
			continue
		}

		k, v := fs[i][:sep], fs[i][sep+1:]

		if !strings.HasPrefix(v, "\"") {
			dst[k] = v
			continue
		}

		// quoted
		i++
		for i < len(fs) {
			v = fmt.Sprintf("%s %s", v, fs[i])
			if strings.HasSuffix(fs[i], "\"") && !strings.HasSuffix(fs[i], "\\\"") {
				break
			}
			i++
		}

		dst[k] = strings.Trim(v, "\"")
	}
	return dst
}
