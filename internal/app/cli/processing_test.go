package cli

import (
	"errors"
	"reflect"
	"testing"
)

func TestProcessArgs(t *testing.T) {
	checkProcessError(t, []string{"ktor", "-f"}, &Error{Err: UnrecognizedFlags{"-f"}, Kind: UnrecognizedFlagsError})
	checkProcessError(t, []string{"ktor", "-a", "-b", "new", "proj"}, &Error{Err: UnrecognizedFlags{"-a", "-b"}, Kind: UnrecognizedFlagsError})
	checkProcessError(t, []string{"ktor"}, &Error{Err: errors.New("command expected"), Kind: NoCommandError})
	checkProcessError(t, []string{"ktor", "nonexistent"}, &Error{Err: CommandError{Command: "nonexistent"}, Kind: CommandNotFoundError})
	checkProcessError(t, []string{"ktor", "new", "a", "b"}, &Error{Err: CommandError{Command: "new"}, Kind: WrongNumberOfArgumentsError})
	checkProcessError(
		t,
		[]string{"ktor", "openapi", "-o", "file.yml"},
		&Error{Err: CommandError{Command: "openapi"}, Kind: WrongNumberOfArgumentsError},
	)
	checkProcessError(
		t,
		[]string{"ktor", "openapi", "file.yml", "-o", "dir"},
		&Error{Err: CommandError{Command: "openapi"}, Kind: WrongNumberOfArgumentsError},
	)
	checkProcessError(
		t,
		[]string{"ktor", "openapi", "-o"},
		&Error{Err: FlagError{Flag: "-o"}, Kind: NoArgumentForFlag},
	)
	checkProcessError(
		t,
		[]string{"ktor", "openapi", "--output"},
		&Error{Err: FlagError{Flag: "--output"}, Kind: NoArgumentForFlag},
	)
	checkProcessError(
		t,
		[]string{"ktor", "openapi", "--output="},
		&Error{Err: FlagError{Flag: "--output"}, Kind: NoArgumentForFlag},
	)
	checkProcessError(
		t,
		[]string{"ktor", "add", "-z", "client-core"},
		&Error{Err: UnrecognizedCommandFlags{Command: "add", Flags: []string{"-z"}}, Kind: UnrecognizedCommandFlagsError},
	)

	checkProcessing(t, []string{"ktor", "--version"}, &Input{Command: VersionCommand})
	checkProcessing(t, []string{"ktor", "-V"}, &Input{Command: VersionCommand})
	checkProcessing(t, []string{"ktor", "-h"}, &Input{Command: Help})
	checkProcessing(t, []string{"ktor", "--help"}, &Input{Command: HelpCommand})
	checkProcessing(t, []string{"ktor", "--version", "--help"}, &Input{Command: VersionCommand})
	checkProcessing(t, []string{"ktor", "--help", "--version"}, &Input{Command: HelpCommand})
	checkProcessing(t, []string{"ktor", "--version", "new"}, &Input{Command: VersionCommand})
	checkProcessing(t, []string{"ktor", "--help", "new", "some"}, &Input{Command: HelpCommand})
	checkProcessing(t, []string{"ktor", "add", "mod1", "mod2"}, &Input{Command: AddCommand, CommandArgs: []string{"mod1", "mod2"}, CommandOptions: map[Flag]string{}})
	checkProcessing(t, []string{"ktor", "new", "some"}, &Input{Command: NewCommand, CommandArgs: []string{"some"}, CommandOptions: map[Flag]string{}})
	checkProcessing(t, []string{"ktor", "-v", "new", "some"}, &Input{Command: NewCommand, CommandArgs: []string{"some"}, Verbose: true, CommandOptions: map[Flag]string{}})
	checkProcessing(
		t,
		[]string{"ktor", "openapi", "-o", "dir", "file.yml"},
		&Input{Command: OpenAPI, CommandArgs: []string{"file.yml"}, CommandOptions: map[Flag]string{OutDir: "dir"}},
	)
	checkProcessing(
		t,
		[]string{"ktor", "-v", "openapi", "-o", "dir", "file.yml"},
		&Input{Command: OpenAPI, CommandArgs: []string{"file.yml"}, CommandOptions: map[Flag]string{OutDir: "dir"}, Verbose: true},
	)
	checkProcessing(
		t,
		[]string{"ktor", "openapi", "--output=dir", "file.yml"},
		&Input{Command: OpenAPI, CommandArgs: []string{"file.yml"}, CommandOptions: map[Flag]string{OutDir: "dir"}},
	)
	checkProcessing(
		t,
		[]string{"ktor", "dev", "-p", "path/to/project"},
		&Input{Command: DevCommand, CommandArgs: []string{}, CommandOptions: map[Flag]string{ProjectDir: "path/to/project"}},
	)
}

func checkProcessing(t *testing.T, args []string, expected *Input) {
	input, err := ProcessArgs(ParseArgs(args))

	if err != nil {
		t.Fatalf("unexpected error %#v", err)
	}

	if !reflect.DeepEqual(input, expected) {
		t.Fatalf("expected %#v, got %#v", expected, input)
	}
}

func checkProcessError(t *testing.T, args []string, expected *Error) {
	_, err := ProcessArgs(ParseArgs(args))

	if err == nil {
		t.Fatalf("expected error, got none")
	}

	if !reflect.DeepEqual(err, expected) {
		t.Fatalf("expected error %#v, got %#v", expected, err)
	}
}
