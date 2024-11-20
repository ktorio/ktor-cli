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
	for _, spec := range AllFlagsSpec {
		if l := len(formatFlags(&spec)); l > maxLen {
			maxLen = l
		}
	}

	for _, spec := range AllFlagsSpec {
		fmt.Fprintf(w, "\t%-*s    %s\n", maxLen, formatFlags(&spec), spec.Description)
	}

	fmt.Fprintln(w, i18n.Get(i18n.CommandsCaption))

	maxLen = 0
	for command := range AllCommandsSpec {
		if command == CompletionCommand {
			continue
		}

		if l := len(formatCommand(command)); l > maxLen {
			maxLen = l
		}
	}

	for command, spec := range AllCommandsSpec {
		if command == CompletionCommand {
			continue
		}

		fmt.Fprintf(w, "\t%-*s    %s\n", maxLen, formatCommand(command), spec.Description)
	}
}
func formatCommand(command Command) string {
	var opts strings.Builder

	sep := ""
	if spec, ok := commandFlagSpec[command]; ok {
		for _, s := range spec {
			opts.WriteString(sep)
			opts.WriteString(formatFlags(&s))

			if s.hasArg {
				opts.WriteString(" <arg>")
			}

			sep = ", "
		}
	}

	if opts.String() != "" {
		return fmt.Sprintf("%s [%s] %s", command, opts.String(), formatArgs(AllCommandsSpec[command].args))
	}

	return fmt.Sprintf("%s %s", command, formatArgs(AllCommandsSpec[command].args))
}

func formatFlags(spec *flagSpec) string {
	return strings.Join(spec.Aliases, ", ")
}
