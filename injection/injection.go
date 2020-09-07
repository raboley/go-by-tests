package injection

import (
	"fmt"
	"io"
	"net/http"
)

// Greet takes in a writer and outputs a greeting to that writer.
// A common example writer would be stdout to write the greeting to the terminal.
func Greet(writer io.Writer, name string) error {
	_, err := fmt.Fprintf(writer, "Hello, %s", name)
	if err != nil {
		return err
	}

	return nil
}

// MyGreeterHandler takes in an http.ResponseWriter to make an http request
// that will host the greeting in the desired format.
func MyGreeterHandler(w http.ResponseWriter, _ *http.Request) {
	_ = Greet(w, "world")
}
