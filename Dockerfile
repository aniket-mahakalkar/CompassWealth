# Use the official Golang image to build the binary
FROM golang:1.25-alpine AS builder

# Install necessary build tools
RUN apk add --no-cache git

# Set the working directory
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the source code
COPY . .

# Build the application
# We build the binary from cmd/server/main.go
RUN CGO_ENABLED=0 GOOS=linux go build -o main ./cmd/server/main.go

# Use a minimal alpine image for the final stage
FROM alpine:latest

# Install CA certificates for HTTPS requests if needed
RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy the binary from the builder stage
COPY --from=builder /app/main .

# Copy the database file if you want to include the initial one
# NOTE: Cloud Run filesystem is ephemeral. Data in this file won't persist across restarts.
COPY --from=builder /app/compass_wealth.db .

# Expose the port
EXPOSE 8080

# Command to run the binary
CMD ["./main"]
