package main

import (
	"fmt"
	"net/http"
)

func login(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	password := r.FormValue("password")

	fmt.Println("Username:", username)
	fmt.Println("Password:", password)

	fmt.Fprintln(w, "Login received")
}

func main() {
	
	http.HandleFunc("/login", login)

	fmt.Println("Server running on port 8080")
	http.ListenAndServe(":8080", nil)
}