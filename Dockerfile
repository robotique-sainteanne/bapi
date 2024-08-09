# Stage 1: Build the Go binary
FROM golang:1.20 AS builder

  # Set the Current Working Directory inside the container
WORKDIR /app

  # Copy the Go Modules manifests
COPY go.mod go.sum ./

  # Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

  # Copy the source code into the container
COPY . .

  # Build the Go binary
RUN go build -tags netgo -ldflags '-s -w' -o app main.go

  # Stage 2: Create a minimal Docker image
FROM alpine:latest

  # Install the necessary CA certificates
RUN apk --no-cache add ca-certificates

  # Set the Current Working Directory inside the container
WORKDIR /root/

  # Copy the Pre-built binary file from the previous stage
COPY --from=builder /app/app .

  # Command to run the executable
CMD ["./app"]
