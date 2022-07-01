.PHONY: build test lint cover clean
all: build
build:
	go get -d -t
	go mod tidy -v
	go build -v -o jirou
clean:
	go clean
	rm -f jirou
