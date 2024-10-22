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

	checkProcessing(t, []string{"ktor", "--version"}, &Input{Command: VersionCommand})
	checkProcessing(t, []string{"ktor", "-V"}, &Input{Command: VersionCommand})
	checkProcessing(t, []string{"ktor", "-h"}, &Input{Command: Help})
	checkProcessing(t, []string{"ktor", "--help"}, &Input{Command: HelpCommand})
	checkProcessing(t, []string{"ktor", "--version", "--help"}, &Input{Command: VersionCommand})
	checkProcessing(t, []string{"ktor", "--help", "--version"}, &Input{Command: HelpCommand})
	checkProcessing(t, []string{"ktor", "--version", "new"}, &Input{Command: VersionCommand})
	checkProcessing(t, []string{"ktor", "--help", "new", "some"}, &Input{Command: HelpCommand})
	checkProcessing(t, []string{"ktor", "new", "some"}, &Input{Command: NewCommand, CommandArgs: []string{"some"}})
	checkProcessing(t, []string{"ktor", "-v", "new", "some"}, &Input{Command: NewCommand, CommandArgs: []string{"some"}, Verbose: true})
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
