package cli

import (
	"fmt"
	"github.com/ktorio/ktor-cli/internal/app/i18n"
	"io"
	"strings"
)

func WriteUsage(w io.Writer) {
	fmt.Fprintf(w, i18n.Get(i18n.ToolSummary))
	fmt.Fprintf(w, i18n.Get(i18n.UsageLine))
	fmt.Fprintln(w, i18n.Get(i18n.OptionsCaption))

	maxLen := 0
	for _, spec := range allFlagsSpec {
		if l := len(formatFlags(&spec)); l > maxLen {
			maxLen = l
		}
	}

	for _, spec := range allFlagsSpec {
		fmt.Fprintf(w, "\t%-*s    %s\n", maxLen, formatFlags(&spec), spec.description)
	}

	fmt.Fprintln(w, i18n.Get(i18n.CommandsCaption))

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
