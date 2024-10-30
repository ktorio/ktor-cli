package model

import (
	"errors"
	"io"
	"os"
	"path/filepath"
	"strings"
)

func IsDirEmptyOrAbsent(dir string) bool {
	f, err := os.Open(dir)
	if err != nil {
		return true
	}
	defer f.Close()

	_, err = f.Readdirnames(1)
	if err == io.EOF {
		return true
	}
	return false
}

func HasNonExistentDirsInPath(p string) (bool, string) {
	volume := filepath.VolumeName(p)
	rest := strings.TrimPrefix(p, volume)
	segments := strings.Split(rest, string(os.PathSeparator))
	for i := 1; i < len(segments); i++ {
		if segments[i-1] == "" {
			continue
		}
		d := volume + strings.Join(segments[:i], string(os.PathSeparator))
		if _, err := os.Stat(d); errors.Is(err, os.ErrNotExist) {
			return true, d
		}
	}

	return false, ""
}
