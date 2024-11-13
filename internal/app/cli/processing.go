package cli

import (
	"errors"
	"slices"
	"strings"
)

type Command string

const (
	NewCommand     Command = "new"
	VersionCommand Command = "version"
	HelpCommand    Command = "help"
	OpenAPI        Command = "openapi"
)

var allCommandsSpec = map[Command]commandSpec{
	OpenAPI:        {args: map[string]Arg{"spec.yml": {required: true}}, description: "generate Ktor project by given OpenAPI specification"},
	NewCommand:     {args: map[string]Arg{"project-name": {required: false}}, description: "generate new Ktor project. If the project name is omitted, an interactive mode will be invoked."},
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
	OutDir       = "outDir"
)

var allFlagsSpec = map[Flag]flagSpec{
	Version: {aliases: []string{"-V", "--version"}, description: "print version"},
	Help:    {aliases: []string{"-h", "--help"}, description: "show the help"},
	Verbose: {aliases: []string{"-v", "--verbose"}, description: "enable verbose mode"},
}

var commandFlagSpec = map[Command]map[Flag]flagSpec{
	OpenAPI: {
		OutDir: {aliases: []string{"-o", "--output"}, description: "output directory", hasArg: true},
	},
}

type flagSpec struct {
	aliases     []string
	description string
	hasArg      bool
}

type Input struct {
	Command        Command
	CommandArgs    []string
	CommandOptions map[Flag]string
	Verbose        bool
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

		found, fl := searchFlag(f, allFlagsSpec)

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

	commandOpts := make(map[Flag]string)
	i := 0
	argsIndex := i
	for i < len(args.CommandArgs) {
		arg := args.CommandArgs[i]
		if !strings.HasPrefix(arg, "-") {
			argsIndex = i
			break
		}

		spec, ok := commandFlagSpec[Command(args.Command)]

		if !ok {
			continue
		}

		if ok, f := searchFlag(arg, spec); ok {
			if spec[f].hasArg {
				if i+1 >= len(args.CommandArgs) {
					return nil, &Error{Err: FlagError{Flag: arg}, Kind: NoArgumentForFlag}
				}

				commandOpts[f] = args.CommandArgs[i+1]
				i++
			} else {
				commandOpts[f] = ""
			}
		}

		i++
	}

	if spec := allCommandsSpec[Command(args.Command)]; requiredArgsCount(spec.args) > 0 && requiredArgsCount(spec.args) != len(args.CommandArgs[argsIndex:]) || len(args.CommandArgs[argsIndex:]) > len(spec.args) {
		return nil, &Error{
			Err:  CommandError{Command: Command(args.Command)},
			Kind: WrongNumberOfArgumentsError,
		}
	}

	return &Input{Command: Command(args.Command), CommandArgs: args.CommandArgs[argsIndex:], CommandOptions: commandOpts, Verbose: flags[Verbose]}, nil
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

func searchFlag(f string, flagMap map[Flag]flagSpec) (bool, Flag) {
	for name, spec := range flagMap {
		if slices.Contains(spec.aliases, f) {
			return true, name
		}
	}

	return false, ""
}
