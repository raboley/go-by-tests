package injection

import (
	"bytes"
	"log"
	"os"
	"testing"
)

func ExampleGreet() {
	name := "Russell"
	writer := os.Stdout
	if err := Greet(writer, name); err != nil {
		log.Fatalf("error thrown greeting: %s with writer: %#v", name, writer)
	}
	// Output: Hello, Russell
}

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	if err := Greet(&buffer, "Chris"); err != nil {
		t.Errorf("want no error, got %q", err)
	}

	got := buffer.String()
	want := "Hello, Chris"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
