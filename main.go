package main

import (
	"log"
	"net/http"
	"os"
)

var GOPORT string

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", index)

	// Set env var with colon ":80".
	err := http.ListenAndServe(os.Getenv("GOPORT"), mux)
	log.Fatal(err)
}

func index(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	switch r.Method {
	case http.MethodGet:
		w.Write([]byte("This is the home page"))

	case http.MethodPost:
		// Handle the POST request...

	case http.MethodOptions:
		w.Header().Set("Allow", "GET, POST, OPTIONS")
		w.WriteHeader(http.StatusNoContent)

	default:
		w.Header().Set("Allow", "GET, POST, OPTIONS")
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}
}
