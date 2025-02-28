# Stage 1: Build Stage
FROM golang:1.23-alpine AS builder

# Create and set the working directory
WORKDIR /app

# Copy the application code
COPY app/ ./

# Get Go dependencies & Build the application (No DEBUG and No C Code in GO)
RUN go mod tidy && \
    CGO_ENABLED=0 go build -ldflags="-s -w" -o myapp ./main.go

# Stage 2: Minimal Runtime Image
FROM scratch

# Set working directory
WORKDIR /app

# Copy built app from the builder stage
COPY --from=builder /app/myapp .

# Expose the app port
EXPOSE 8080

# Run the Go application
CMD ["/app/myapp"]