# Build stage
FROM golang:1.21-alpine AS builder

WORKDIR /app

# Install build dependencies
RUN apk add --no-cache git

# Pre-copy/cache go.mod for better layer caching
COPY go.mod go.sum ./
RUN go mod download -x && go mod verify

# Copy the source code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -o /go/bin/app

# Final stage
FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy the binary from builder
COPY --from=builder /go/bin/app .
# Copy any config files if needed
COPY --from=builder /app/.env.example ./.env

# Expose the port the app runs on
EXPOSE 8080

# Command to run the executable
CMD ["./app"]