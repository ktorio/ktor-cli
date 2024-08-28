package command

import (
	"fmt"
	"runtime/debug"
	"strings"
)

func Version(mainVersion string) {
	fmt.Printf("Ktor CLI version %s\n", getVersion(mainVersion))
}

func getVersion(mainVersion string) string {
	if mainVersion != "" {
		return strings.Trim(mainVersion, "\n\r")
	}

	if info, ok := debug.ReadBuildInfo(); ok {
		for _, setting := range info.Settings {
			if setting.Key == "vcs.revision" {
				return fmt.Sprintf("dev-%s", setting.Value[:7])
			}
		}
	}

	return "dev"
}
