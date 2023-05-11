BIN = app.exe

migrate-up:
	migrate -path internal/repository/migrations -database ${DB_URL} -verbose up

migrate-down:
	migrate -path internal/repository/migrations -database ${DB_URL} -verbose down

build: migrate-up
	go build -o ${BIN} ./cmd/app/main.go

run: build
	./${BIN}

.DEFAULT_GOAL := run

