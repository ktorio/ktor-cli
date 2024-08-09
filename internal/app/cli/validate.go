package cli

// A map of a command to a number of args
var availCommands = map[string]int{
	"new": 1,
}

// ValidateArgs returns ErrorType or error
func ValidateArgs(args *Args) error {
	if args.Command == nil {
		return NoCommandError
	}

	if _, ok := availCommands[*args.Command]; !ok {
		return CommandNotFoundError
	}

	if count := availCommands[*args.Command]; count != len(args.CommandArgs) {
		return WrongNumberOfArgumentsError
	}

	return nil
}
