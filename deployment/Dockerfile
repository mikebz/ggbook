# Build stage
FROM golang:1.24-alpine AS builder

WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code
COPY . .

# Build the Go binary
RUN CGO_ENABLED=0 GOOS=linux go build -o ggbook

# Final image
FROM alpine:3.22

WORKDIR /app

# Copy the binary from the builder
COPY --from=builder /app/ggbook .

# Copy static files (optional, if you want to serve html)
COPY --from=builder /app/html ./html

EXPOSE 8080

ENTRYPOINT ["./ggbook"]