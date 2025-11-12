# ---------- Этап system ----------
FROM golang:1.23-alpine AS system
RUN apk add --no-cache git build-base
RUN rm -rf /var/cache/apk/*

# ---------- Этап build ----------
FROM system AS build
WORKDIR /app
COPY go.mod ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 go build -ldflags="-s -w" -o main .

# ---------- Этап run ----------
FROM alpine:latest
RUN adduser -D -u 1000 appuser
WORKDIR /app
COPY --from=build /app/main .
RUN chown appuser:appuser /app/main && chmod +x /app/main
USER appuser
ENV APP_NAME="Go MultiStage App"
ENV APP_PORT=8080
EXPOSE 8080
CMD ["./main"]
