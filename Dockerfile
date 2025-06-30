# Start from the official Golang image
FROM golang:1.24 AS builder

WORKDIR /task-manager

# Copy go mod files and download deps
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the app
COPY . .

# Generate Swagger docs
RUN go install github.com/swaggo/swag/cmd/swag@latest && swag init

# Build the Go binary
ENV CGO_ENABLED=0
RUN go build -o main .

# Production image
FROM debian:bullseye-slim

WORKDIR /task-manager
COPY --from=builder /task-manager/main .
COPY --from=builder /task-manager/docs ./docs

# Expose the port
EXPOSE 8080

# Run the app
CMD ["./main"]
