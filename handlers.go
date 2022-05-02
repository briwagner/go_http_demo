package main

import (
	"log"
	"net/http"
)

// HomeHandler handles the home route.
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		log.Printf("visitor to %s", r.URL.Path)

		http.Redirect(w, r, "/", http.StatusNotFound)
		return
	}

	switch r.Method {
	case http.MethodGet:
		log.Println("visitor to homepage")
		w.Write([]byte("This is the home page"))

	case http.MethodOptions:
		w.Header().Set("Allow", "GET, POST, OPTIONS")
		w.WriteHeader(http.StatusNoContent)

	default:
		w.Header().Set("Allow", "GET, POST, OPTIONS")
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}
}

// SayHandler handles the /say route.
func SayHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		w.Write([]byte("This is the say route. What did you say?"))
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
