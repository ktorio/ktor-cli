package utils

import (
	"io"
	"os"
)

func IsDirEmpty(dir string) bool {
	f, err := os.Open(dir)

	if err != nil {
		return false
	}

	defer f.Close()

	names, err := f.Readdirnames(1)
	return err == io.EOF && len(names) == 0
}
