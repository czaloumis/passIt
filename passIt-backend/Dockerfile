# syntax=docker/dockerfile:1

# -------- Stage 1: Build --------
FROM golang:1.24 AS builder

WORKDIR /app

# Copy go.mod and go.sum first
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the source code
COPY . .

# Build the Go app
ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64
RUN go build -ldflags="-extldflags=-static" -o main ./cmd/api/main.go

# -------- Stage 2: Run --------
FROM gcr.io/distroless/static:nonroot

WORKDIR /app
COPY --from=builder /app/main .

# Run the binary
CMD ["./main"]
