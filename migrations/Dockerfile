FROM golang:1.19-alpine AS builder

RUN go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

WORKDIR /usr/src/app
COPY . .

FROM alpine:latest

COPY --from=builder /usr/src/app .
COPY --from=builder /go/bin/migrate .
CMD ./migrate -path . -database $DB_URL -verbose up
