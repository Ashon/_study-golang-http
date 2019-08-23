FROM golang:alpine AS builder

WORKDIR /app

RUN apk update && apk add --no-cache git

COPY . .
RUN go get -d -v
RUN GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o /app/main

ENTRYPOINT ["/app/main"]
