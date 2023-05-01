build:
	go build  -o app ./cmd/app/main.go

run: build
	./app

.DEFAULT_GOAL := run

