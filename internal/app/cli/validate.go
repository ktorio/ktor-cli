package cli

// ValidateArgs returns ErrorType or error
func ValidateArgs(args *Args) error {
	if args.Command == nil {
		return NoCommandError
	}

	if _, ok := allCommands[Command(*args.Command)]; !ok {
		return CommandNotFoundError
	}

	if count := allCommands[Command(*args.Command)]; count != len(args.CommandArgs) {
		return WrongNumberOfArgumentsError
	}

	return nil
}
