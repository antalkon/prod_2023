# Dockerfile was generated from
# https://github.com/lodthe/dockerfiles/blob/main/go/Dockerfile

# Stage 1: Builder
FROM golang:1.23-alpine3.19 AS builder

# Install required software for building the application
RUN apk update && apk add --no-cache ca-certificates git gcc g++ libc-dev binutils

WORKDIR /opt

# Download and verify dependencies
COPY go.mod go.sum ./
RUN go mod download && go mod verify

# Copy application source
COPY . .

# Build the application, targeting cmd/network/main.go
RUN go build -o bin/application cmd/network/main.go

# Stage 2: Runner
FROM alpine:3.19 AS runner

# Install runtime dependencies
RUN apk update && apk add --no-cache ca-certificates libc6-compat openssh bash && rm -rf /var/cache/apk/*

WORKDIR /opt

# Copy built application from builder stage
COPY --from=builder /opt/bin/application ./

# Run the application
CMD ["./application"]