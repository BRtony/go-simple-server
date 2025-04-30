# Build stage
FROM golang:1.21-alpine AS builder

WORKDIR /app

# Install git and build dependencies
RUN apk add --no-cache git

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy source code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -o main .

# Final stage
FROM alpine:latest

WORKDIR /app

# Install CA certificates and timezone data
RUN apk add --no-cache ca-certificates tzdata

# Copy the binary from builder
COPY --from=builder /app/main .
COPY --from=builder /app/static ./static

# Create necessary directories
RUN mkdir -p static/images

# Expose port
EXPOSE 8080

# Configure startup command
ENV HOST=0.0.0.0
ENV PORT=8080

# Run the application
CMD ["./main"] 