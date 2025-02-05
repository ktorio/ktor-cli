package cli

import "fmt"

type ShellError struct {
	Shell string
}

func (se ShellError) Error() string {
	return fmt.Sprintf("unrecognized shell %s", se.Shell)
}
