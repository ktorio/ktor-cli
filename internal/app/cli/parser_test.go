package cli

import (
	"reflect"
	"testing"
)

func TestParseArgs(t *testing.T) {
	checkParsing(t, []string{}, &Args{})
	checkParsing(t, []string{"ktor"}, &Args{})
	checkParsing(t, []string{"ktor", "-v"}, &Args{Flags: []string{"-v"}})
	checkParsing(t, []string{"ktor", "--version", "--help"}, &Args{Flags: []string{"--version", "--help"}})
	checkParsing(t, []string{"ktor", "command"}, &Args{Command: "command"})
	checkParsing(t, []string{"ktor", "-v", "command"}, &Args{Flags: []string{"-v"}, Command: "command"})
	checkParsing(
		t,
		[]string{"ktor", "command", "arg1", "arg2"},
		&Args{CommandArgs: []string{"arg1", "arg2"}, Command: "command"},
	)
	checkParsing(
		t,
		[]string{"ktor", "-v", "command", "arg1", "arg2"},
		&Args{Flags: []string{"-v"}, CommandArgs: []string{"arg1", "arg2"}, Command: "command"},
	)
	checkParsing(
		t,
		[]string{"ktor", "command", "-v1", "-v2"},
		&Args{Command: "command", CommandArgs: []string{"-v1", "-v2"}},
	)
}

func checkParsing(t *testing.T, rawArgs []string, expected *Args) {
	args := ParseArgs(rawArgs)

	if !reflect.DeepEqual(ParseArgs(rawArgs), expected) {
		t.Fatalf("expected %#v, got %#v", expected, args)
	}
}
