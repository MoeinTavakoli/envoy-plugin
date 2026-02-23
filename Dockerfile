FROM golang:1.21-alpine

WORKDIR /app

# Copy go mod files
COPY go.mod ./

# Copy source code
COPY show_headers.go .

# Build the application
RUN go build -o show_header show_headers.go

# Expose port
EXPOSE 8080

# Run the application
CMD ["./show_header"]
