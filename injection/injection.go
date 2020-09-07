package injection

import (
	"fmt"
	"io"
	"net/http"
)

func Greet(writer io.Writer, name string) error {
	_, err := fmt.Fprintf(writer, "Hello, %s", name)
	if err != nil {
		return err
	}

	return nil
}

func MyGreeterHandler(w http.ResponseWriter, _ *http.Request) {
	_ = Greet(w, "world")
}
