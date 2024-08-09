# Use the official Golang image to build the Go application
FROM golang:1.22.5 AS builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy the Go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the Go app
RUN go build -tags netgo -ldflags '-s -w' -o app

# Start a new stage from scratch
FROM debian:bullseye-slim

# Set the Current Working Directory inside the container
WORKDIR /root/

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /app/app .

# Expose port 8080 for the application
EXPOSE 5000

# Set environment variables (default values)
ENV DB_HOST=localhost \
    DB_USERNAME=user \
    DB_NAME=dbname \
    DB_PORT=5432 \
    DB_PASSWORD=password

# Command to run the executable
CMD ["./app"]
