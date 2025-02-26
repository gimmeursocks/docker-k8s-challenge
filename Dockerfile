# Stage 1: Build Stage
FROM golang:1.23 AS builder

# Create and set the working directory
WORKDIR /app

# Copy the application code
COPY app/ ./

# Get Go dependencies & Build the Go application
RUN go mod tidy && go build -o myapp ./main.go

# Stage 2: Smaller Runtime Image
FROM debian:bookworm-slim

# Set working directory
WORKDIR /app

# Copy built app from the builder stage
COPY --from=builder /app/myapp .

# Expose the app port
EXPOSE 8080

# Run the Go application
CMD ["/app/myapp"]