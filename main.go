package main

import (
	"github.com/raboley/go-by-tests/injection/injection"
	"log"
	"os"
)

func main() {
	writer := os.Stdout
	if err := injection.Greet(writer, "Russell"); err != nil {
		log.Fatalf("Got an error trying to make a greeting! error: %q", err)
	}
}
