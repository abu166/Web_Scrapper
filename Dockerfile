# Build stage
FROM golang:1.24 AS builder
WORKDIR /app
COPY . .
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o web_scrapper cmd/main.go

FROM chromedp/headless-shell:latest
WORKDIR /app
COPY --from=builder /app/web_scrapper /app/web_scrapper
COPY .env /app/.env
RUN chmod +x /app/web_scrapper
ENTRYPOINT ["/app/web_scrapper"]