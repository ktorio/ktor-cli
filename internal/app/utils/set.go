package utils

import (
	"errors"
	"fmt"
)

type StringSet map[string]struct{}

func NewStringSet() StringSet {
	return make(StringSet)
}

func (s *StringSet) Add(v string) {
	(*s)[v] = struct{}{}
}

func (s *StringSet) Entries() (entries []string) {
	for k := range *s {
		entries = append(entries, k)
	}

	return
}

func (s *StringSet) Single() (v string, err error) {
	if len(*s) > 1 {
		err = errors.New(fmt.Sprintf("expected at most one entry in the set, got %d", len(*s)))
		return
	}
	for k := range *s {
		v = k
		return
	}

	return
}
