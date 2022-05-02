package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/newrelic/go-agent/v3/newrelic"
)

var PORT string

func main() {
	app, err := newrelic.NewApplication(
		newrelic.ConfigAppName("Render demo"),
		newrelic.ConfigLicense(os.Getenv("NEWRELIC")),
		newrelic.ConfigDistributedTracerEnabled(true),
	)

	if err != nil {
		panic(err)
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/say", SayHandler)
	mux.HandleFunc(newrelic.WrapHandleFunc(app, "/", HomeHandler))

	// Set env var with colon ":80".
	appPort := fmt.Sprintf(":%s", os.Getenv("PORT"))
	err = http.ListenAndServe(appPort, mux)
	log.Fatal(err)
}
