# Build stage
FROM golang:1.24.6 AS builder

WORKDIR /app

# Download dependencies first to leverage Docker layer caching
COPY go.mod go.sum ./
RUN go mod download

# Copy source
COPY . .

# Build statically linked binary
RUN CGO_ENABLED=0 GOOS=linux go build -o server ./cmd/server

# Runtime stage
FROM debian:bookworm-slim

WORKDIR /app

# Install certs, ping/traceroute tools, and bash for the scripts
RUN apt-get update && apt-get install -y --no-install-recommends \
    ca-certificates \
    iputils-ping \
    traceroute \
    bash \
    && rm -rf /var/lib/apt/lists/*

# Non-root user
RUN useradd -u 10001 appuser

# App files
COPY --from=builder /app/server /app/server
COPY config ./config
COPY scripts ./scripts
COPY start.sh ./start.sh

# Ensure scripts are executable and owned by runtime user
RUN chmod +x /app/scripts/ping_test.sh /app/start.sh && chown -R appuser:appuser /app

EXPOSE 8080
ENV GIN_MODE=release

USER appuser
ENTRYPOINT ["/app/start.sh"]
