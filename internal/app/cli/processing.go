package cli

import (
	"errors"
	"fmt"
	"github.com/ktorio/ktor-cli/internal/app/i18n"
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
	OpenAPI:        {args: map[string]Arg{"spec.yml": {required: true}}, description: i18n.Get(i18n.OpenApiCommandDescr)},
	NewCommand:     {args: map[string]Arg{"project-name": {required: false}}, description: i18n.Get(i18n.NewCommandDescr)},
	VersionCommand: {args: map[string]Arg{}, description: i18n.Get(i18n.VersionCommandDescr)},
	HelpCommand:    {args: map[string]Arg{}, description: i18n.Get(i18n.HelpCommandDescr)},
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
	Version: {aliases: []string{"-V", "--version"}, description: i18n.Get(i18n.VersionCommandDescr)},
	Help:    {aliases: []string{"-h", "--help"}, description: i18n.Get(i18n.HelpCommandDescr)},
	Verbose: {aliases: []string{"-v", "--verbose"}, description: i18n.Get(i18n.VerboseOptionDescr)},
}

var commandFlagSpec = map[Command]map[Flag]flagSpec{
	OpenAPI: {
		OutDir: {aliases: []string{"-o", "--output"}, description: i18n.Get(i18n.OutputDirOptionDescr), hasArg: true},
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
	for i, f := range args.Flags {
		if slices.Contains(allFlagsSpec[Version].aliases, f) {
			version = true
			break
		}

		if slices.Contains(allFlagsSpec[Help].aliases, f) {
			help = true
			break
		}

		fl, _, found := searchFlag(f, allFlagsSpec, i)

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

		if f, argPos, ok := searchFlag(arg, spec, i); ok {
			if spec[f].hasArg {
				if argPos == i {
					parts := strings.Split(arg, "=")
					if len(parts) < 2 || parts[1] == "" {
						return nil, &Error{Err: FlagError{Flag: parts[0]}, Kind: NoArgumentForFlag}
					}
					commandOpts[f] = parts[1]
				} else {
					if argPos >= len(args.CommandArgs) {
						return nil, &Error{Err: FlagError{Flag: arg}, Kind: NoArgumentForFlag}
					}

					commandOpts[f] = args.CommandArgs[argPos]
					i++
				}
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

func searchFlag(f string, flagMap map[Flag]flagSpec, argIndex int) (Flag, int, bool) {
	for name, spec := range flagMap {
		for _, al := range spec.aliases {
			if al == f {
				return name, argIndex + 1, true
			}

			if strings.HasPrefix(f, fmt.Sprintf("%s=", al)) {
				return name, argIndex, true
			}
		}
	}

	return "", 0, false
}
