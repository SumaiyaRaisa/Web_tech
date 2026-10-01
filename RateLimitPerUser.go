// fixed window - per user
package main

import (
	"fmt"
	"net"
	"net/http"
	"sync"
	"time"
)

type User struct {
	count int
}

var (
	users = make(map[string]*User)
	mu    sync.Mutex
)

// middleware
func rateLimiter(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		ip, _, err := net.SplitHostPort(r.RemoteAddr)
		if err != nil {
			http.Error(w, "Invalid request", http.StatusBadRequest)
			return
		}

		mu.Lock()

		if users[ip] == nil {
			users[ip] = &User{}
		}

		users[ip].count++
		count := users[ip].count

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
		time.Sleep(5 * time.Second)

		mu.Lock()

		for _, user := range users {
			user.count = 0
		}

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