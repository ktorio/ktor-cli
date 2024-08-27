package progress

import (
	"io"
)

func clearCurrentLine(w io.Writer) error {
	_, err := io.WriteString(w, "\033[2K\r")
	return err
}
