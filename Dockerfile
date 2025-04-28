# 第一階段：Builder
FROM golang:1.24.2 AS builder

WORKDIR /source

COPY go.mod go.sum ./
RUN go mod download

COPY . .

ENV CGO_ENABLED=0 GOOS=linux GOARCH=arm64

# 編譯出來的 binary 叫 cryptoo-pricing
RUN go build -o cryptoo-pricing ./cmd

# 第二階段：Production Image
FROM debian:bookworm-slim

RUN apt-get update && apt-get install -y ca-certificates && rm -rf /var/lib/apt/lists/*

WORKDIR /app

# 複製 binary
COPY --from=builder /source/cryptoo-pricing .

# 複製 .env
COPY --from=builder /source/config/.env ./config/.env

# 預設啟動指令
CMD ["./cryptoo-pricing"]