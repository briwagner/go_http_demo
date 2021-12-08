package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

var PORT string

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", HomeHandler)

	// Set env var with colon ":80".
	appPort := fmt.Sprintf(":%s", os.Getenv("PORT"))
	err := http.ListenAndServe(appPort, mux)
	log.Fatal(err)
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		log.Printf("visitor to %s", r.URL.Path)

		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	switch r.Method {
	case http.MethodGet:
		log.Println("visitor to homepage")
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
