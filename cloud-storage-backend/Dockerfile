# Stage 1: Build the Golang app
FROM golang:1.21-alpine AS build

# Set the working directory
WORKDIR /app

# Copy the rest of the application code
COPY . .

# Build the Golang app
RUN go build -o main .

# Stage 2: Serve the app using the binary
FROM alpine:latest

# Set the working directory
WORKDIR /root/

# Copy the built binary from the previous stage
COPY --from=build /app/main .

# Expose port 5000
EXPOSE 5000

# Start the Golang app
CMD ["./main"]