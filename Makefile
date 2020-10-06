build: gen
	go build

deps: go-tools

go-tools:
	go get -u github.com/jessevdk/go-assets-builder

gen:
	go generate
