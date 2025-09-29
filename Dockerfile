# 1. Build stage
FROM golang:1.25-alpine AS builder

# Install git for dependency fetching
RUN apk add --no-cache git

# Set working directory
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o smartsearch ./cmd/server/main.go

FROM alpine:latest

RUN addgroup -S appgroup && adduser -S appuser -G appgroup

WORKDIR /app

COPY --from=builder /app/smartsearch .

RUN chown -R appuser:appgroup /app

USER appuser

COPY ./configs ./configs

RUN mkdir -p /app/.appdata

EXPOSE 8080

CMD ["./smartsearch"]
