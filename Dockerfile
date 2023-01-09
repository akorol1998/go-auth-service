# syntax=docker/dockerfile:1

FROM golang:1.19

WORKDIR /usr/src/app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY . .

RUN ls -la
RUN go build -o /app ./cmd

CMD ["/app"]

