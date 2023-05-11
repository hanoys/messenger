migrate-up:
	migrate -path internal/repository/migrations -database ${DB_URL} -verbose up

migrate-down:
	migrate -path internal/repository/migrations -database ${DB_URL} -verbose down

build:
	go build -o app.exe ./cmd/app/main.go

run: build
	./app

.DEFAULT_GOAL := run

