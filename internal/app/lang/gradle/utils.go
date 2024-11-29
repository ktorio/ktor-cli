package gradle

import "strings"

func PlatformSuffix(artifact string) string {
	if strings.HasSuffix(artifact, "-jvm") {
		return "-jvm"
	}

	return ""
}
