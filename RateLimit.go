package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

var (
	requestCount = 0
	mu sync.Mutex
)

func rateLimiter(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		mu.Lock()
		requestCount++
		count := requestCount
		mu.Unlock()

		if count > 5 {
			http.Error(w, "Too many requests", http.StatusTooManyRequests)
			return
		}

		next(w, r)
	}
}

func resetCounter() {
	for {
		time.Sleep(1 * time.Second)

		mu.Lock()
		requestCount = 0
		mu.Unlock()
	}
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello! Request accepted.")
}

func main() {

	go resetCounter()

	http.HandleFunc("/", rateLimiter(helloHandler))

	fmt.Println("Server running at http://localhost:8080")

	http.ListenAndServe(":8080", nil)
}

//This version allows 5 requests per second

// Middleware is a function that sits between the client request and your actual handler.
// A goroutine is a lightweight way to run a function concurrently in Go.