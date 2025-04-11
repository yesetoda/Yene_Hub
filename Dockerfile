# Build stage
FROM golang:1.24-alpine AS builder

# Install build dependencies
RUN apk add --no-cache git ca-certificates build-base postgresql-dev

WORKDIR /app

# Copy module files first
COPY go.mod go.sum ./

# Download dependencies with proper retry
RUN for i in $(seq 1 3); do \
      go mod download && break || sleep 5; \
    done

# Copy source code
COPY . .

# Build
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o /app/main .

# Final stage
FROM alpine:3.18

RUN apk add --no-cache ca-certificates postgresql-client && \
    adduser -D -u 1000 -g '' appuser && \
    mkdir -p /app/config && \
    chown -R appuser:appuser /app

WORKDIR /app
COPY --from=builder --chown=appuser:appuser /app/main .
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --chown=appuser:appuser .env.example .
COPY --chown=appuser:appuser config/*.yaml ./config/

USER appuser
EXPOSE 8080
CMD ["./main"]