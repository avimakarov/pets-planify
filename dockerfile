FROM golang:alpine AS builder

WORKDIR /build
COPY . .

RUN go build -o bot ./cmd/bot
RUN go build -o server ./cmd/server

FROM alpine

WORKDIR /app

COPY --from=builder /build/bot /app/bot
COPY --from=builder /build/server /app/server
