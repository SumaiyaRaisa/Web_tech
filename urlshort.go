package main

import (
	"fmt"
	"net/http"
)

// urls stores short code -> original URL
var urls = map[string]string{}

// counter gives each new URL a unique number
var counter = 1

func shortenHandler(w http.ResponseWriter, r *http.Request) {
	longURL := r.URL.Query().Get("url")
	if longURL == "" {
		fmt.Fprintln(w, "Please provide a url, e.g. /shorten?url=https://example.com")
		return
	}

	code := fmt.Sprintf("%d", counter)
	counter++
	urls[code] = longURL

	fmt.Fprintf(w, "Short link: http://localhost:8080/%s\n", code)
}

func redirectHandler(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Path[1:] // remove leading "/"

	longURL, found := urls[code]
	if !found {
		http.Error(w, "Short URL not found", http.StatusNotFound)
		return
	}

	http.Redirect(w, r, longURL, http.StatusFound)
}

func main() {
	http.HandleFunc("/shorten", shortenHandler)
	http.HandleFunc("/", redirectHandler)

	fmt.Println("Server running at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}