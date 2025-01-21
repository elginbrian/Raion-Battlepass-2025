FROM golang:1.23-alpine as builder

RUN apk add --no-cache git bash

WORKDIR /app

COPY go.mod go.sum ./ 
RUN go mod download

COPY vendor/ ./vendor/

COPY . .

COPY db/migrations /app/db/migrations

COPY public/ /app/public/

COPY cmd/index.html /app/cmd/index.html

RUN go build -mod=vendor -o /app/raion-battlepass ./cmd/main.go

FROM alpine:latest

WORKDIR /app

RUN apk add --no-cache bash

COPY --from=builder /app/raion-battlepass /app/raion-battlepass

EXPOSE 8080

CMD ["/app/raion-battlepass"]