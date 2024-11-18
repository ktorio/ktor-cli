package model

type ErrorKind int

const (
	TerminalHeightError ErrorKind = iota
	UnableFetchPluginsError
	ProjectDirNotEmptyError
	DirNotExistError
	ProjectDirTooLongError
	ProjectNameEmptyError
	ProjectNameAllowedCharsError
)

func (mdl *State) SetError(k ErrorKind, s string) {
	mdl.errorMap[k] = s
}

func (mdl *State) RemoveErrors(ks ...ErrorKind) {
	for _, k := range ks {
		delete(mdl.errorMap, k)
	}
}

func (mdl *State) GetErrors(ks ...ErrorKind) []string {
	var errs []string
	for _, k := range ks {
		if e, ok := mdl.errorMap[k]; ok {
			errs = append(errs, e)
		}
	}
	return errs
}
