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
	NewCommand        Command = "new"
	VersionCommand    Command = "version"
	HelpCommand       Command = "help"
	CompletionCommand Command = "completions"
	AddCommand        Command = "add"
	OpenAPI           Command = "openapi"
	DevCommand        Command = "dev"
)

var AllCommandsSpec = map[Command]commandSpec{
	DevCommand:        {args: map[string]Arg{}, Description: i18n.Get(i18n.DevCommandDescr)},
	OpenAPI:           {args: map[string]Arg{"spec.yml": {required: true}}, Description: i18n.Get(i18n.OpenApiCommandDescr)},
	NewCommand:        {args: map[string]Arg{"project-name": {required: false}}, Description: i18n.Get(i18n.NewCommandDescr)},
	AddCommand:        {args: map[string]Arg{"...module": {required: true}}, Description: i18n.Get(i18n.AddCommandDescr)},
	VersionCommand:    {args: map[string]Arg{}, Description: i18n.Get(i18n.VersionCommandDescr)},
	HelpCommand:       {args: map[string]Arg{}, Description: i18n.Get(i18n.HelpCommandDescr)},
	CompletionCommand: {args: map[string]Arg{"shell": {required: true}}, Description: i18n.Get(i18n.CompletionCommandDescr)},
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
	Version    Flag = "version"
	Help            = "help"
	Verbose         = "verbose"
	OutDir          = "outDir"
	ProjectDir      = "projectDir"
)

var AllFlagsSpec = map[Flag]flagSpec{
	Version: {Aliases: []string{"-V", "--version"}, Description: i18n.Get(i18n.VersionCommandDescr)},
	Help:    {Aliases: []string{"-h", "--help"}, Description: i18n.Get(i18n.HelpCommandDescr)},
	Verbose: {Aliases: []string{"-v", "--verbose"}, Description: i18n.Get(i18n.VerboseOptionDescr)},
}
var commandFlagSpec = map[Command]map[Flag]flagSpec{
	OpenAPI: {
		OutDir: {Aliases: []string{"-o", "--output"}, Description: i18n.Get(i18n.OutputDirOptionDescr), hasArg: true},
	},
	AddCommand: {
		ProjectDir: {Aliases: []string{"-p", "--project"}, Description: i18n.Get(i18n.ProjectDirOptionDescr), hasArg: true},
	},
	DevCommand: {
		ProjectDir: {Aliases: []string{"-p", "--project"}, Description: i18n.Get(i18n.ProjectDirOptionDescr), hasArg: true},
	},
}

type flagSpec struct {
	Aliases     []string
	Description string
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
		if slices.Contains(AllFlagsSpec[Version].Aliases, f) {
			version = true
			break
		}

		if slices.Contains(AllFlagsSpec[Help].Aliases, f) {
			help = true
			break
		}

		fl, _, found := searchFlag(f, AllFlagsSpec, i)

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

	commandOpts := make(map[Flag]string)
	i := 0
	argsIndex := i
	unrecognized = unrecognized[:0]
	for i < len(args.CommandArgs) {
		arg := args.CommandArgs[i]
		if !strings.HasPrefix(arg, "-") {
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
		} else {
			unrecognized = append(unrecognized, arg)
		}

		i++
		argsIndex = i
	}

	if len(unrecognized) > 0 {
		return nil, &Error{Err: UnrecognizedCommandFlags{Command: args.Command, Flags: unrecognized}, Kind: UnrecognizedCommandFlagsError}
	}

	spec := AllCommandsSpec[Command(args.Command)]
	actualArgs := args.CommandArgs[argsIndex:]

	if !hasVararg(spec) && (requiredArgsCount(spec.args) > 0 && requiredArgsCount(spec.args) != len(actualArgs) || len(actualArgs) > len(spec.args)) {
		return nil, &Error{
			Err:  CommandError{Command: Command(args.Command)},
			Kind: WrongNumberOfArgumentsError,
		}
	}

	return &Input{Command: Command(args.Command), CommandArgs: args.CommandArgs[argsIndex:], CommandOptions: commandOpts, Verbose: flags[Verbose]}, nil
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

func searchFlag(f string, flagMap map[Flag]flagSpec, argIndex int) (Flag, int, bool) {
	for name, spec := range flagMap {
		for _, al := range spec.Aliases {
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
