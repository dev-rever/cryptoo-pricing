# 第一階段：Builder
FROM golang:1.24.2 AS builder

WORKDIR /source

ENV GOOS=linux
ENV GOARCH=amd64
ENV CGO_ENABLED=0

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o cryptoo-pricing ./cmd

FROM debian:bookworm-slim

RUN apt-get update && apt-get install -y ca-certificates && rm -rf /var/lib/apt/lists/*

WORKDIR /app

COPY --from=builder /source/cryptoo-pricing .
COPY --from=builder /source/config/.env ./config/.env
COPY --from=builder /source/docs/ ./docs
COPY --from=builder /source/templates ./templates

COPY README.md ./README.md
COPY README.zh.md ./README.zh.md

CMD ["./cryptoo-pricing"]