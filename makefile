BINARY_NAME=blog-app
.DEFAULT_GOAL := run

build:
	go build -o ./dist/${BINARY_NAME} main.go

run: build
	./dist/${BINARY_NAME}