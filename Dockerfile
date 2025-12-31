# Stage 1: Build the Go app
FROM golang:1.25-alpine AS builder

# Set working directory
WORKDIR /app

# Copy go.mod and go.sum first for caching dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the source code
COPY . .

# Build the binary (with CGO disabled for alpine compatibility)
RUN CGO_ENABLED=0 GOOS=linux go build -o main ./cmd/api/main.go

# Stage 2: Create a small runtime image
FROM alpine:3.20

# Install ca-certificates for HTTPS (if needed)
#RUN apk --no-cache add ca-certificates

# Set working directory
WORKDIR /app

# Copy the binary from builder stage
COPY --from=builder /app/main .

# Expose the port (match with your .env PORT, e.g., 8080)
EXPOSE 8080

# Run the app
CMD ["./main"]