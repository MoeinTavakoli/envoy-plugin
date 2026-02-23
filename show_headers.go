package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

// Simple HTTP server
func main() {
	http.HandleFunc("/authz", func(w http.ResponseWriter, r *http.Request) {
		// Log Authorization header, current time, and route
		log.Printf("Authorization header received: %s", r.Header.Get("Authorization"))
		log.Printf("Request path: %s", r.URL.Path)
		log.Printf("Timestamp: %s", time.Now().Format(time.RFC3339))

		// Respond to Envoy with a simple message
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("OK"))
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Log Authorization header, current time, and route
		log.Printf("Authorization header received: %s", r.Header.Get("Authorization"))
		log.Printf("Request path: %s", r.URL.Path)
		log.Printf("Timestamp: %s", time.Now().Format(time.RFC3339))

		// Response headers
		w.Header().Set("x-go-update", "processed-by-go")
		w.Header().Set("content-type", "text/plain; charset=utf-8")
		w.Header().Set("Authorization", r.Header.Get("Authorization"))

		// Response body
		response := `Hello from Go Update Simulation!

Request Path: ` + r.URL.Path + `
Authorization received by upstream: ` + r.Header.Get("Authorization") + `
`

		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, response)
	})

	log.Println("Go service running on :8080")
	log.Println("AuthZ endpoint:  http://127.0.0.1:8080/authz")
	log.Println("Demo endpoint:   http://127.0.0.1:8080/")
	if err := http.ListenAndServe("0.0.0.0:8080", nil); err != nil {
		log.Fatal(err)
	}
}

