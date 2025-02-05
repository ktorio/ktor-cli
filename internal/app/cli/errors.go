package cli

import "fmt"

type ErrorKind int

type Error struct {
	Kind ErrorKind
	Err  error
}

func (e Error) Error() string {
	return fmt.Sprintf("cli args error: %v", e.Kind)
}

const (
	NoCommandError ErrorKind = iota
	CommandNotFoundError
	WrongNumberOfArgumentsError
	UnrecognizedFlagsError
	UnrecognizedCommandFlagsError
	NoArgumentForFlag
)

type UnrecognizedFlags []string

func (f UnrecognizedFlags) Error() string {
	return fmt.Sprintf("unrecognized flags: %sv", []string(f))
}

type UnrecognizedCommandFlags struct {
	Command string
	Flags   []string
}

func (e UnrecognizedCommandFlags) Error() string {
	return fmt.Sprintf("command %s unrecognized flags: %v", e.Command, e.Flags)
}

type CommandError struct {
	Command Command
}

func (e CommandError) Error() string {
	return fmt.Sprintf("command '%s' error", e.Command)
}

type FlagError struct {
	Flag string
}

func (e FlagError) Error() string {
	return fmt.Sprintf("flag '%s' error", e.Flag)
}
