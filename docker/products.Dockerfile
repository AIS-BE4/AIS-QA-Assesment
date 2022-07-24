FROM golang:1.17-alpine

WORKDIR /app

ENV CONFIG=docker

COPY .. /app
COPY ../.env /app

RUN apk add --update nodejs npm
RUN go get github.com/githubnemo/CompileDaemon
RUN go mod download

CMD npx sequelize db:create && npm i && npx sequelize db:migrate

ENTRYPOINT CompileDaemon --build="go build -o main products/main.go" --command=./main