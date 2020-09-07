package injection

import (
	"fmt"
	"io"
)

func Greet(writer io.Writer, name string) error {
	_, err := fmt.Fprintf(writer, "Hello, %s", name)
	if err != nil {
		return err
	}

	return nil
}
