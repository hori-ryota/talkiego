build: gen
	go build

deps: go-tools npm-tools-global npm-install

go-tools:
	go get -u github.com/jessevdk/go-assets-builder

npm-install:
	cd assets/Talkie && yarn install

npm-tools-global: 
	npm install -g yarn

gen:
	go generate
