package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

// Simple HTTP server that acts like an Envoy update for demonstration
func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Add custom headers (simulating Envoy Go update behavior)
		w.Header().Set("x-go-update", "processed-by-go")
		w.Header().Set("x-processing-time", time.Now().Format(time.RFC3339))
		w.Header().Set("x-request-id", fmt.Sprintf("go-%d", time.Now().UnixNano()))
		
		// Preserve original user agent
		if userAgent := r.Header.Get("user-agent"); userAgent != "" {
			w.Header().Set("x-original-user-agent", userAgent)
		}
		
		// Special route logic
		if r.URL.Path == "/special" {
			w.Header().Set("x-special-route", "true")
		}
		
		// Response headers
		w.Header().Set("x-go-response", "processed-by-go-update")
		w.Header().Set("x-response-time", time.Now().Format(time.RFC3339))
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		
		// Response body
		response := `Hello from Go Update Simulation!

This simulates how an Envoy Go update would modify headers.

Response Headers Added:
- x-go-update: processed-by-go
- x-go-response: processed-by-go-update
- x-processing-time: ` + time.Now().Format(time.RFC3339) + `
- x-request-id: ` + fmt.Sprintf("go-%d", time.Now().UnixNano()) + `
- x-original-user-agent: ` + r.Header.Get("user-agent") + `

Request Path: ` + r.URL.Path + `
`
		
		if r.URL.Path == "/special" {
			response += `
Special route detected! x-special-route header was added.
`
		}
		
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, response)
	})
	
	log.Println("Go update simulation server running on port 8080")
	log.Println("Test with:")
	log.Println("  curl -v http://localhost:8080/")
	log.Println("  curl -v http://localhost:8080/special")
	
	// Listen on all interfaces, not just localhost
	if err := http.ListenAndServe("0.0.0.0:8080", nil); err != nil {
		log.Fatal("Server failed to start:", err)
	}
}
