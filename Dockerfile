# Stage 1: Build the Go binary
FROM golang:1.23-alpine AS builder

# Install necessary dependencies
RUN apk add --no-cache curl git

# Set the working directory
WORKDIR /app

# Copy go.mod and go.sum to install dependencies
COPY go.mod go.sum ./
RUN go mod tidy

# Copy the remaining code and build the binary
COPY . .
RUN GOOS=linux GOARCH=amd64 go build -o main .

# Stage 2: Lambda Runtime Stage
FROM public.ecr.aws/lambda/provided:al2 AS lambda

COPY --from=builder /app/main /var/task/
CMD ["main"]

# Stage 3: Local Development Image
FROM alpine:3.18 AS local

# Expose port 8080 for local access
EXPOSE 8080

# Copy the binary from the builder stage
COPY --from=builder /app/main /main

# Command to run the binary
CMD ["/main"]
