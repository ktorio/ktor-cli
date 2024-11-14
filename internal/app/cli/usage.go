package cli

import (
	"fmt"
	"io"
	"strings"
)

func WriteUsage(w io.Writer) {
	fmt.Fprintf(w, "Ktor is a tool mainly for generating Ktor projects.\n\n")
	fmt.Fprintf(w, "Usage: ktor [options] <command> [arguments]\n\n")
	fmt.Fprintln(w, "The options are:")

	maxLen := 0
	for _, spec := range allFlagsSpec {
		if l := len(formatFlags(&spec)); l > maxLen {
			maxLen = l
		}
	}

	for _, spec := range allFlagsSpec {
		fmt.Fprintf(w, "\t%-*s    %s\n", maxLen, formatFlags(&spec), spec.description)
	}

	fmt.Fprintln(w, "The commands are:")

	maxLen = 0
	for command := range allCommandsSpec {
		if l := len(formatCommand(command)); l > maxLen {
			maxLen = l
		}
	}

	for command, spec := range allCommandsSpec {
		fmt.Fprintf(w, "\t%-*s    %s\n", maxLen, formatCommand(command), spec.description)
	}
}
func formatCommand(command Command) string {
	return fmt.Sprintf("%s %s", command, formatArgs(allCommandsSpec[command].args))
}

func formatFlags(spec *flagSpec) string {
	return strings.Join(spec.aliases, ", ")
}
