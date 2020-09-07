package main

import (
	"fmt"
	"github.com/raboley/go-by-tests/injection/injection"
	"log"
	"net/http"
	"os"
)

func main() {
	// io.writer can be used for stdout to terminal
	writer := os.Stdout
	if err := injection.Greet(writer, "Russell"); err != nil {
		log.Fatalf("Got an error trying to make a greeting! error: %q", err)
	}

	// io.writer can also be used for http requests
	port := "5000"
	if err := http.ListenAndServe(fmt.Sprintf(":%s", port), http.HandlerFunc(injection.MyGreeterHandler)); err != nil {
		log.Fatalf("Error trying to serve over port %s", port)
	}
}
