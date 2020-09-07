package main

import (
	"fmt"
	"github.com/raboley/go-by-tests/injection/injection"
	"github.com/raboley/go-by-tests/injection/mocking"
	"log"
	"net/http"
	"os"
	"time"
)

var packageToRun string

func main() {
	packageToRun = getPackageToRun()
	if packageToRun == "injection" {
		mainInjection()
	}
	if packageToRun == "mocking" {
		mainMocking()
	}

	// In progress code here:

}

func getPackageToRun() (packageName string) {
	if len(os.Args) > 1 {
		packageName = "injection"
	}
	return packageToRun
}

func mainInjection() {
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

func mainMocking() {
	sleeper := mocking.NewConfigurableSleeper(5*time.Second, time.Sleep)
	mocking.Countdown(os.Stdout, sleeper)
}
