package cli

import (
	"errors"
	"slices"
	"strings"
)

type Command string

const (
	NewCommand        Command = "new"
	VersionCommand    Command = "version"
	HelpCommand       Command = "help"
	CompletionCommand Command = "completions"
	AddCommand        Command = "add"
)

var AllCommandsSpec = map[Command]commandSpec{
	NewCommand:        {args: map[string]Arg{"project-name": {required: false}}, Description: "generate new Ktor project. If the project name is omitted run an interactive mode."},
	AddCommand:        {args: map[string]Arg{"...module": {required: true}}, Description: "add Ktor modules to a project"},
	VersionCommand:    {args: map[string]Arg{}, Description: "print version"},
	HelpCommand:       {args: map[string]Arg{}, Description: "show the help"},
	CompletionCommand: {args: map[string]Arg{"shell": {required: true}}, Description: "auto completions for different shells"},
}

type Arg struct {
	required bool
}

type commandSpec struct {
	args        map[string]Arg
	Description string
}

type Flag string

const (
	Version Flag = "version"
	Help         = "help"
	Verbose      = "verbose"
)

var AllFlagsSpec = map[Flag]flagSpec{
	Version: {Aliases: []string{"-V", "--version"}, Description: "print version"},
	Help:    {Aliases: []string{"-h", "--help"}, Description: "show the help"},
	Verbose: {Aliases: []string{"-v", "--verbose"}, Description: "enable verbose mode"},
}

type flagSpec struct {
	Aliases     []string
	Description string
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
		if slices.Contains(AllFlagsSpec[Version].Aliases, f) {
			version = true
			break
		}

		if slices.Contains(AllFlagsSpec[Help].Aliases, f) {
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

	if _, ok := AllCommandsSpec[Command(args.Command)]; !ok {
		return nil, &Error{Err: CommandError{Command: Command(args.Command)}, Kind: CommandNotFoundError}
	}

	if spec := AllCommandsSpec[Command(args.Command)]; !hasVararg(spec) && (requiredArgsCount(spec.args) > 0 && requiredArgsCount(spec.args) != len(args.CommandArgs) || len(args.CommandArgs) > len(spec.args)) {
		return nil, &Error{
			Err:  CommandError{Command: Command(args.Command)},
			Kind: WrongNumberOfArgumentsError,
		}
	}

	return &Input{Command: Command(args.Command), CommandArgs: args.CommandArgs, Verbose: flags[Verbose]}, nil
}

func hasVararg(spec commandSpec) bool {
	for k := range spec.args {
		if strings.HasPrefix(k, "...") {
			return true
		}
	}
	return false
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
	for name, spec := range AllFlagsSpec {
		if slices.Contains(spec.Aliases, f) {
			return true, name
		}
	}

	return false, ""
}
