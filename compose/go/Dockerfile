# Use the official Go image
FROM golang:1.21-alpine as builder

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files to the working directory
COPY go.mod .
COPY go.sum .

# Download dependencies
RUN go mod download

# Copy the entire project to the working directory
COPY . .

# Build the Go application
RUN go build -o user-api ./cmd/user-api

# Start a new stage for a smaller image
FROM alpine:latest

# Set the working directory inside the container
WORKDIR /app

# Copy the binary from the builder stage to the current stage
COPY --from=builder /app/user-api .
COPY --from=builder /app/docs ./docs

# Copy the .env file
COPY --from=builder /app/.env .

# Expose the port that the application will run on
EXPOSE 8080

# Command to run the executable
CMD ["./user-api"]
