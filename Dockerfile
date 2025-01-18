# Step 1: Build the Go application    
FROM golang:1.23-alpine as builder

RUN apk add --no-cache git bash

WORKDIR /app

# Copy Go modules and dependencies
COPY go.mod go.sum ./ 
RUN go mod download

# Copy vendor folder if you are using vendoring
COPY vendor/ ./vendor/

# Copy the source code
COPY . .

# Copy the migrations directory
COPY db/migrations /app/db/migrations

# Ensure the public directory exists and copy it
COPY public/ /app/public/

# Copy the cmd directory into the container
COPY cmd/index.html /app/cmd/index.html

# Build the Go application
RUN go build -mod=vendor -o /app/fiber-starter ./cmd/main.go

# Step 2: Create the final image
FROM alpine:latest

WORKDIR /app

# Install bash
RUN apk add --no-cache bash

# Copy the Go executable from the builder stage
COPY --from=builder /app/fiber-starter /app/fiber-starter

# Copy the wait-for-it.sh script into the final image
COPY wait-for-it.sh /wait-for-it.sh
RUN chmod +x /wait-for-it.sh

# Expose the port
EXPOSE 8080

# Set the command to run the Go application with wait-for-it
CMD ["/wait-for-it.sh", "db:5432", "--", "/app/fiber-starter"]
