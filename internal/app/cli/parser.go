package cli

import (
	"strings"
)

type Args struct {
	Options     Options
	Command     *string
	CommandArgs []string
}

type Options struct {
	Verbose *bool
}

func ParseArgs(args []string) (*Args, error) {
	if len(args) < 2 {
		return nil, NotEnoughArgsError
	}

	verbose := false
	i := 1
	for ; i < len(args); i++ {
		arg := args[i]
		if !strings.HasPrefix(arg, "-") {
			break
		}

		if arg == "-v" {
			verbose = true
		}
	}

	var command *string
	if i < len(args) {
		command = &args[i]
	}
	var commandArgs []string
	if i < len(args) {
		commandArgs = args[i+1:]
	}

	return &Args{
		Options:     Options{Verbose: &verbose},
		Command:     command,
		CommandArgs: commandArgs,
	}, nil
}
