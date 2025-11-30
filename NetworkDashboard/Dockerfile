FROM golang:1.24.6 AS builder
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
# 不再 COPY package.json，也不再 build 前端
# static 文件直接随项目复制

RUN CGO_ENABLED=0 GOOS=linux go build -o server ./cmd/server

FROM debian:bookworm-slim
WORKDIR /app

RUN apt-get update && apt-get install -y --no-install-recommends \
    ca-certificates \
    iputils-ping \
    traceroute \
    bash \
    tzdata && \
    rm -rf /var/lib/apt/lists/*

ENV TZ=America/Los_Angeles
RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone


RUN useradd -u 10001 appuser

COPY --from=builder /app/server .
COPY --from=builder /app/static ./static
COPY config ./config
COPY scripts ./scripts
COPY start.sh ./start.sh

RUN chmod +x /app/scripts/ping_test.sh /app/start.sh \
    && chown -R appuser:appuser /app

EXPOSE 8080
ENV GIN_MODE=release

USER appuser
ENTRYPOINT ["/app/start.sh"]
