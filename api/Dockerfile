FROM golang:1.20.1-alpine3.17 AS builder

WORKDIR /pear-open-app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o api ./server.go

FROM alpine:3.17 AS pear-open-api

COPY --from=builder /pear-open-app/api /pear-open-api
EXPOSE 8000

CMD ["pear-open-api"]
