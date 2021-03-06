talkiego
=========

talkiego is wrapper for [Talkie](https://github.com/ahomu/Talkie).

## Usage

```
Usage:
  talkiego serve [index.md] [flags]

Flags:
  -f, --backface          backface option
  -c, --control           control option
  -h, --help              help for serve
  -l, --linkShouldBlank   linkShouldBlank option
  -t, --noTransition      noTransition option
  -b, --pointer           pointer option
  -p, --port uint         listen port (default 3000)
  -r, --progress          progress option
  -w, --wide              wide option

Global Flags:
      --config string    config file (default is $HOME/.talkiego.yaml)
  -d, --divider string   divider for slide (default "---")
```

## Getting started

### Install

```
go get github.com/hori-ryota/talkiego
```

### Example

Following code is simple example.

```
<!-- rootopt: wide control -->

<!-- cover: -->
# Talkie.js

---

<!-- title: invert
            backface=https://farm8.staticflickr.com/7469/16209884502_211d01ac0d_z_d.jpg
            backfaceFilter="blur(3px) brightness(.9)" -->
# Backface (北極)

---

<!-- bullets: style="color:#ff0000;" -->

* foo
* bar
* baz
```

### Format

#### RootOpt

RootOpt can following formats.

```
<!-- rootopt: [option] [option] -->
<!-- rootopt: [option=boolean] [option=boolean] -->
<!-- rootopt: [option:boolean] [option:boolean] -->
```

If you input only propaty name, value  will be true.

Examples are following.

```
<!-- rootopt: wide control pointer progress backface notransition linkShouldBlank -->
<!-- rootopt: wide=true control=false -->
```

#### PageOpt

PageOpt can following formats.

```
<!-- [layout]: -->
<!-- [layout]: [option=value] [option=value] -->
<!-- [layout]: [option:value] [option:value] -->
```

Examples are following.

```
<!-- cover: -->
```

```
<!-- title: invert
            backface=https://farm8.staticflickr.com/7469/16209884502_211d01ac0d_z_d.jpg
            backfaceFilter="blur(3px) brightness(.9)" -->
```

## See Also

* [ahomu/Talkie](https://github.com/ahomu/Talkie)

## Inspired

* [yusukebe/revealgo](https://github.com/yusukebe/revealgo)
* [nakajmg/talkie-generator](https://github.com/nakajmg/talkie-generator)

## TODO

* [x] Allow RootOpt settings with command flags
* [ ] Allow custom thema
* [ ] Allow custom css
* [ ] Add watch option
