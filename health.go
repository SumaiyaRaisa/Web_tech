package main

import (
	"encoding/json"
	"log"
	"net/http"
)

// healthHandler responds to GET /health with a simple status message.
func healthHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string {
		"status": "ok",
	})
}

func main() {
	http.HandleFunc("/health", healthHandler)

	log.Println("Server running at http://localhost:8080/health")
	log.Fatal(http.ListenAndServe(":8080", nil))
}