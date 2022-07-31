.PHONY: build build-run

build:
	go build -o ./bin/gh-gomerge

build-run: build
	./bin/gh-gomerge
