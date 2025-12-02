FROM golang:alpine AS builder

WORKDIR /build
COPY . .

RUN go build -o server ./cmd/server

FROM alpine

WORKDIR /app

COPY --from=builder /build/server /app/server
