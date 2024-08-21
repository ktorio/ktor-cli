package cli

import (
	"fmt"
	"io"
)

func WriteUsage(w io.Writer) {
	fmt.Fprintf(w, "Ktor is a tool mainly for generating Ktor projects.\n\n")
	fmt.Fprintf(w, "Usage: ktor [options] <command> [arguments]\n\n")
	fmt.Fprintln(w, "The options are:")
	fmt.Fprintln(w, "\t-v\tenable verbose mode")
	fmt.Fprintln(w, "The commands are:")
	fmt.Fprintln(w, "\tnew <project-name>\tgenerate new Ktor project")
}
