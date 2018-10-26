build: gen
	go build

deps: go-tools npm-install

go-tools:
	go get -u github.com/jessevdk/go-assets-builder

npm-install:
	cd assets/Talkie && npm install

gen:
	go generate
