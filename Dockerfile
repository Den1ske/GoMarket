
FROM golang:latest

WORKDIR /app

COPY ./src /app

RUN go mod download

RUN go get github.com/githubnemo/CompileDaemon

ENTRYPOINT CompileDaemon --build="make build" --command=./server