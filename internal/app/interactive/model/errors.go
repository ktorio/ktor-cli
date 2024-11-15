package model

import (
	"slices"
	"strings"
)

type ErrorKind int

const (
	TerminalHeightError ErrorKind = iota
	UnableFetchPluginsError
	ProjectDirNotEmptyError
	DirNotExistError
	ProjectDirTooLongError
)

func (mdl *State) SetError(k ErrorKind, s string) {
	mdl.errorMap[k] = s
}

func (mdl *State) RemoveErrors(ks ...ErrorKind) {
	for _, k := range ks {
		delete(mdl.errorMap, k)
	}
}

func (mdl *State) FormatErrors() string {
	var errs []string
	for _, e := range mdl.errorMap {
		errs = append(errs, e)
	}

	slices.Sort(errs)

	return strings.Join(errs, "; ")
}
