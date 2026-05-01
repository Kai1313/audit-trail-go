# --- Stage 1: Build Stage ---
FROM golang:1.21-alpine AS builder

# Install build dependencies
RUN apk add --no-cache git

# Set working directory
WORKDIR /app

# Copy go.mod and go.sum first to leverage Docker cache
COPY go.mod go.sum ./
RUN go mod download

# Copy the vendor folder and source code
COPY . .

# Build the application
# CGO_ENABLED=0 ensures a static binary for Alpine
RUN CGO_ENABLED=0 GOOS=linux go build -o audit-service main.go

# --- Stage 2: Final Run Stage ---
FROM alpine:latest

# Install ca-certificates in case you need to call external APIs (HTTPS)
RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy the binary from the builder stage
COPY --from=builder /app/audit-service .

# Copy the .env file (Optional: usually handled via environment variables in Docker)
COPY --from=builder /app/.env .

# Copy swagger docs so they can be served
COPY --from=builder /app/docs ./docs

# Expose the application port
EXPOSE 8080

# Command to run the executable
CMD ["./audit-service"]