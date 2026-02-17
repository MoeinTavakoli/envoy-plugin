FROM golang:1.21-alpine

WORKDIR /app

# Copy go mod files
COPY go.mod ./

# Copy source code
COPY simple_update_header.go .

# Build the application
RUN go build -o update_header simple_update_header.go

# Expose port
EXPOSE 8080

# Run the application
CMD ["./update_header"]
