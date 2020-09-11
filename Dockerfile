FROM golang:1.15.1-alpine

RUN apk update && apk upgrade && \
    apk add --no-cache bash git openssh

WORKDIR /app

COPY ./app .
COPY ./app/go.mod ./app/go.sum ./
RUN go mod download

RUN go build -o main .

#RUN export PATH=$PATH:$GOPATH/bin

#RUN reflex -r '\.go$' -s -- sh -c "go build -o main ."

EXPOSE 8080

CMD ["./main"]