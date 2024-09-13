# syntax=docker/dockerfile:1

FROM golang:1.21.5

WORKDIR /app

COPY ./ ./
RUN go mod download


RUN go build -o /risks-api

EXPOSE 8080

CMD ["/risks-api"]