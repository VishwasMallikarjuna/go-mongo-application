# Use the official Golang image as the build environment
FROM golang:1.20-alpine AS builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the Go app
RUN go build -o main ./src/main.go

# Use a minimal image as the runtime environment
FROM alpine:latest

# Set the Current Working Directory inside the container
WORKDIR /root/

# Copy the pre-built binary file from the builder stage
COPY --from=builder /app/main .

# Command to run the executable
CMD ["./main"]
