# syntax=docker/dockerfile:1
FROM golang:1.21-alpine

WORKDIR /app

COPY go.mod ./
COPY main.go ./

RUN go build -o main .

EXPOSE 8081

CMD ["./main"]
