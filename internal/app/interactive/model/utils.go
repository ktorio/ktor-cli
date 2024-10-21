package model

import (
	"errors"
	"io"
	"os"
	"path/filepath"
	"strings"
	"unicode"
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

func FindNonAlphaNumeric(start int, str string) int {
	for i, r := range str {
		if i >= start+1 && !unicode.IsLetter(r) && !unicode.IsNumber(r) {
			return i
		}
	}

	return len(str)
}

func FindNonAlphaNumericFromEnd(start int, str string) int {
	runes := []rune(str)

	specialSeq := true
	for i := start - 1; i >= 0; i-- {
		isSpecial := !unicode.IsLetter(runes[i]) && !unicode.IsNumber(runes[i])
		if isSpecial && specialSeq {
			continue
		}

		if !isSpecial {
			specialSeq = false
		}

		if isSpecial {
			return i
		}
	}

	return 0
}
