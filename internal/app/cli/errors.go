package cli

type ErrorType int

func (e ErrorType) Error() string {
	return "CLI args error"
}

const (
	NotEnoughArgsError ErrorType = iota
	NoCommandError
	CommandNotFoundError
	WrongNumberOfArgumentsError
)
