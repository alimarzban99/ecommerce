# =============================================================================
# Build Stage
# =============================================================================
FROM golang:1.25-alpine AS builder

# Set build arguments
ARG VERSION=unknown
ARG BUILD_TIME=unknown
ARG COMMIT_SHA=unknown

# Install build dependencies
RUN apk add --no-cache git ca-certificates tzdata

# Set working directory
WORKDIR /build

# Copy dependency files first for better layer caching
COPY go.mod go.sum ./

# Download dependencies (cached layer if go.mod/go.sum unchanged)
RUN go mod download && \
    go mod verify

# Copy source code
COPY . .

# Build the application
# - CGO_ENABLED=0: Disable CGO for static binary
# - GOOS=linux: Target Linux OS
# - GOARCH=amd64: Target AMD64 architecture
# - -ldflags: Inject build metadata
# - -trimpath: Remove file system paths from the resulting executable
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build \
    -ldflags="-w -s -X 'main.Version=${VERSION}' -X 'main.BuildTime=${BUILD_TIME}' -X 'main.CommitSHA=${COMMIT_SHA}'" \
    -trimpath \
    -o /build/bin/api \
    ./cmd/api/main.go

# =============================================================================
# Runtime Stage
# =============================================================================
FROM alpine:3.20

# Install runtime dependencies
RUN apk --no-cache add \
    ca-certificates \
    tzdata \
    wget \
    && \
    # Create non-root user for security
    addgroup -g 1000 appuser && \
    adduser -D -u 1000 -G appuser appuser && \
    # Create app directory
    mkdir -p /app && \
    chown -R appuser:appuser /app

# Set working directory
WORKDIR /app

# Copy binary from builder stage
COPY --from=builder --chown=appuser:appuser /build/bin/api /app/api

# Copy keys directory if it exists (for JWT keys)
# Note: In production, keys should be mounted as volumes or use secrets management
COPY --chown=appuser:appuser keys/ /app/keys/

# Switch to non-root user
USER appuser

# Expose application port
EXPOSE 8080

# Health check
HEALTHCHECK --interval=30s --timeout=3s --start-period=5s --retries=3 \
    CMD wget --no-verbose --tries=1 --spider http://localhost:8080/health || exit 1

# Set default environment variables
ENV APP_ENVIRONMENT=production \
    APP_LOG_LEVEL=info

# Run the application
ENTRYPOINT ["/app/api"]