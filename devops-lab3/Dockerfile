# ---------- Этап 1: Сборка ----------
FROM golang:1.23-alpine AS builder

RUN apk add --no-cache git
WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 go build -ldflags="-s -w" -o main .

# ---------- Этап 2: Запуск ----------
FROM alpine:latest

RUN adduser -D -u 1000 appuser
WORKDIR /app

COPY --from=builder /app/main .
RUN chown appuser:appuser /app/main && chmod +x /app/main

USER appuser

ENV APP_NAME="Go MultiStage App"
ENV APP_PORT=8080

EXPOSE 8080
CMD ["./main"]
