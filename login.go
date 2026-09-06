package main

import (
	"encoding/json"
	"log"
	"net/http"
)

// healthHandler responds to GET /health with a simple status message.
func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"status": "ok",
	})
}

// Credentials represents the JSON body sent by the client to log in.
type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// loginHandler responds to POST /login with success or failure JSON.
func loginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var creds Credentials
	if err := json.NewDecoder(r.Body).Decode(&creds); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	// Hardcoded check for now — replace with database lookup later.
	if creds.Username == "admin" && creds.Password == "1234" {
		json.NewEncoder(w).Encode(map[string]string{
			"message": "login successful",
		})
	} else {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(map[string]string{
			"message": "invalid username or password",
		})
	}
}

func main() {
	http.HandleFunc("/health", healthHandler)
	http.HandleFunc("/login", loginHandler)

	log.Println("Server running at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}