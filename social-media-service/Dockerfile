# Build stage
FROM golang:1.20-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o social-service ./cmd/main.go

# Final stage
FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/social-service .

ENV PORT=8080
EXPOSE 8080

CMD ["./social-service"]
