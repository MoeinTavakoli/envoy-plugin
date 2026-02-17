# Go Update Header Example - Header Manipulation

A simple and clean example demonstrating header manipulation using Go, simulating what you would do with Envoy Go extensions.

## Overview

This project shows how to modify HTTP headers using Go. While this is a simulation using a standalone HTTP server, the same header manipulation logic applies to actual Envoy Go extensions.

## What This Header Update Does

### Request Headers Added:
- `x-go-update`: "processed-by-go" - Indicates the Go update header is active
- `x-original-user-agent`: Preserves the original user agent
- `x-processing-time`: Timestamp when request was processed
- `x-request-id`: Unique request identifier
- `x-special-route`: "true" (only when accessing `/special` path)

### Response Headers Added:
- `x-go-response`: "processed-by-go-update" - Indicates response processing
- `x-response-time`: Timestamp when response was sent

## Quick Start

### Prerequisites
- Docker and Docker Compose
- Go 1.21+ (for local development)

### Running with Docker Compose

1. **Build and start the Go update header:**
   ```bash
   docker compose -f docker-compose-host.yml up -d --build
   ```

2. **Test the update header with localhost access:**
   ```bash
   # Basic request
   curl -v http://localhost:8080/
   
   # Request with custom user agent
   curl -v -H "User-Agent: MyCustomAgent/1.0" http://localhost:8080/
   
   # Special route to see conditional logic
   curl -v http://localhost:8080/special
   ```

3. **Stop the service:**
   ```bash
   docker compose -f docker-compose-host.yml down
   ```

### Local Development

1. **Run directly:**
   ```bash
   go run simple_update.go
   ```

2. **Build and run:**
   ```bash
   go build -o update_header simple_update.go
   ./update_header
   ```

## Example Responses

### Request to `http://localhost:8080/`
```bash
curl -v -H "User-Agent: TestAgent/1.0" http://localhost:8080/
```

**Response Headers:**
```
HTTP/1.1 200 OK
x-go-update: processed-by-go
x-processing-time: 2026-02-16T11:30:00Z
x-request-id: go-1739701800000000000
x-original-user-agent: TestAgent/1.0
x-go-response: processed-by-go-update
x-response-time: 2026-02-16T11:30:00Z
content-type: text/plain
```

**Response Body:**
```
Hello from Go Update Header Simulation!

This simulates how an Envoy Go update header would modify headers.

Response Headers Added:
- x-go-update: processed-by-go
- x-go-response: processed-by-go-update
- x-processing-time: 2026-02-16T11:30:00Z
- x-request-id: go-1739701800000000000
- x-original-user-agent: TestAgent/1.0

Request Path: /
```

### Request to `http://localhost:8080/special`
```bash
curl -v http://localhost:8080/special
```

**Additional Response Header:**
```
x-special-route: true
```

## File Structure

```
.
├── simple_update.go    # Go update header simulation
├── go.mod             # Go module definition
├── Dockerfile         # Docker build file
├── docker-compose.yml # Docker Compose setup
└── README.md          # This documentation
```

## Understanding the Code

### Go Update Header (`simple_update.go`)

The update header demonstrates common header manipulation patterns:

```go
// Add custom headers
w.Header().Set("x-go-update", "processed-by-go")

// Preserve original headers
if userAgent := r.Header.Get("user-agent"); userAgent != "" {
    w.Header().Set("x-original-user-agent", userAgent)
}

// Conditional logic
if r.URL.Path == "/special" {
    w.Header().Set("x-special-route", "true")
}
```

## Development Tips

### Adding New Headers
```go
// Request headers (in real Envoy update header)
header.Set("x-custom-header", "custom-value")

// Response headers
w.Header().Set("x-response-custom", "response-value")
```

### Conditional Logic
```go
if path := r.URL.Path; path == "/admin" {
    w.Header().Set("x-admin-access", "true")
}

// Check for specific headers
if auth := r.Header.Get("authorization"); auth != "" {
    w.Header().Set("x-auth-processed", "true")
}
```

### Time-based Headers
```go
w.Header().Set("x-timestamp", time.Now().Format(time.RFC3339))
w.Header().Set("x-unix-time", fmt.Sprintf("%d", time.Now().Unix()))
```

## Real Envoy Go Extension

For actual Envoy Go extensions, you would:

1. Use the Envoy Go API:
```go
import "github.com/envoyproxy/envoy/contrib/golang/common/go/api"

type headerFilter struct {
    api.FilterConfig
}

func (f *headerFilter) DecodeHeaders(header api.RequestHeaderMap, endStream bool) api.StatusType {
    header.Set("x-go-update", "processed-by-go")
    return api.Continue
}
```

2. Build as shared library:
```bash
go build -buildmode=c-shared -o update_header.so main.go
```

3. Configure in Envoy YAML to load the update header.

## Clean Up

```bash
docker compose -f docker-compose-host.yml down
docker system prune -f
```

## Next Steps

- Add authentication logic
- Implement rate limiting
- Add request/response body transformation
- Integrate with external services
- Add metrics and logging
- Convert to real Envoy Go extension

## Resources

- [Envoy Go Extensions](https://www.envoyproxy.io/docs/envoy/latest/api-docs/extensions/filters/http/golang)
- [Go HTTP Server](https://golang.org/pkg/net/http/)
- [Docker Compose](https://docs.docker.com/compose/)
