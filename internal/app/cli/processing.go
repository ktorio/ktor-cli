package cli

import (
	"errors"
	"slices"
)

type Command string

const (
	NewCommand     Command = "new"
	VersionCommand Command = "version"
	HelpCommand    Command = "help"
)

var allCommandsSpec = map[Command]commandSpec{
	NewCommand:     {args: map[string]Arg{"project-name": {required: false}}, description: "generate new Ktor project. If the project name is omitted run an interactive mode."},
	VersionCommand: {args: map[string]Arg{}, description: "print version"},
	HelpCommand:    {args: map[string]Arg{}, description: "show the help"},
}

type Arg struct {
	required bool
}

type commandSpec struct {
	args        map[string]Arg
	description string
}

type Flag string

const (
	Version Flag = "version"
	Help         = "help"
	Verbose      = "verbose"
)

var allFlagsSpec = map[Flag]flagSpec{
	Version: {aliases: []string{"-V", "--version"}, description: "print version"},
	Help:    {aliases: []string{"-h", "--help"}, description: "show the help"},
	Verbose: {aliases: []string{"-v", "--verbose"}, description: "enable verbose mode"},
}

type flagSpec struct {
	aliases     []string
	description string
}

type Input struct {
	Command     Command
	CommandArgs []string
	Verbose     bool
}

func ProcessArgs(args *Args) (*Input, error) {
	version := false
	help := false
	var unrecognized []string
	var flags = make(map[Flag]bool)
	for _, f := range args.Flags {
		if slices.Contains(allFlagsSpec[Version].aliases, f) {
			version = true
			break
		}

		if slices.Contains(allFlagsSpec[Help].aliases, f) {
			help = true
			break
		}

		found, fl := searchFlag(f)

		if !found && !slices.Contains(unrecognized, f) {
			unrecognized = append(unrecognized, f)
		}

		if found {
			flags[fl] = true
		}
	}

	if len(unrecognized) > 0 {
		return nil, &Error{Err: UnrecognizedFlags(unrecognized), Kind: UnrecognizedFlagsError}
	}

	if version {
		return &Input{Command: VersionCommand}, nil
	}

	if help {
		return &Input{Command: HelpCommand}, nil
	}

	if args.Command == "" {
		return nil, &Error{Err: errors.New("command expected"), Kind: NoCommandError}
	}

	if _, ok := allCommandsSpec[Command(args.Command)]; !ok {
		return nil, &Error{Err: CommandError{Command: Command(args.Command)}, Kind: CommandNotFoundError}
	}

	if spec := allCommandsSpec[Command(args.Command)]; requiredArgsCount(spec.args) > 0 && requiredArgsCount(spec.args) != len(args.CommandArgs) || len(args.CommandArgs) > len(spec.args) {
		return nil, &Error{
			Err:  CommandError{Command: Command(args.Command)},
			Kind: WrongNumberOfArgumentsError,
		}
	}

	return &Input{Command: Command(args.Command), CommandArgs: args.CommandArgs, Verbose: flags[Verbose]}, nil
}

func requiredArgsCount(args map[string]Arg) int {
	count := 0
	for _, arg := range args {
		if arg.required {
			count++
		}
	}

	return count
}

func searchFlag(f string) (bool, Flag) {
	for name, spec := range allFlagsSpec {
		if slices.Contains(spec.aliases, f) {
			return true, name
		}
	}

	return false, ""
}
