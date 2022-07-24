FROM golang:1.17-alpine

WORKDIR /app

ENV CONFIG=docker

COPY .. /app
COPY ../.env /app

RUN apk add --update nodejs npm
RUN go get github.com/githubnemo/CompileDaemon
RUN go mod download


ENTRYPOINT CompileDaemon --build="go build -o main stores/main.go" --command=./main