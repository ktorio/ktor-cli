package cli

import (
	"strings"
)

type Args struct {
	Flags       []string
	Command     string
	CommandArgs []string
}

type Options struct {
	Verbose bool
}

func ParseArgs(args []string) *Args {
	i := 1
	var flags []string
	for ; i < len(args); i++ {
		arg := args[i]
		if !strings.HasPrefix(arg, "-") {
			break
		}

		flags = append(flags, arg)
	}

	command := ""
	if i < len(args) {
		command = args[i]
	}
	var commandArgs []string
	if i+1 < len(args) {
		commandArgs = args[i+1:]
	}

	return &Args{
		Flags:       flags,
		Command:     command,
		CommandArgs: commandArgs,
	}
}
