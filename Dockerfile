FROM golang:1.19-alpine

RUN apk update && apk upgrade
RUN apk add --no-cache bash make
RUN go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

WORKDIR /usr/src/app

COPY go.mod go.sum ./

RUN go mod download && go mod verify

COPY . .

CMD ["make"]
