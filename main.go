package main

import (
	"fmt"
	"github.com/raboley/go-by-tests/injection/injection"
	"github.com/raboley/go-by-tests/injection/mocking"
	"log"
	"net/http"
	"os"
)

func main() {
	var packageToRun string
	if len(os.Args) > 1 {
		packageToRun = "injection"
	}
	if packageToRun == "injection" {
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

	sleeper := &mocking.DefaultSleeper{}
	mocking.Countdown(os.Stdout, sleeper)

}
