BINARY_NAME=blog-app
.DEFAULT_GOAL := run

build:
	go build -o ./dist/${BINARY_NAME} main.go

run: build
	./dist/${BINARY_NAME}

quick:
	git add . && git commit -m "No message"
	git push

push:
	git add . && git commit -m "${m}"
	git push
