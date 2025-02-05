package gradle

import (
	"github.com/ktorio/ktor-cli/internal/app/lang"
	"os"
	"strings"
)

func ParseProps(fp string) map[string]string {
	m := make(map[string]string)

	b, err := os.ReadFile(fp)

	if err != nil {
		return m
	}

	for _, line := range strings.Split(string(b), "\n") {
		parts := strings.Split(line, "=")

		if len(parts) == 2 {
			m[strings.TrimSpace(parts[0])] = lang.Unquote(strings.TrimSpace(parts[1]))
		}
	}

	return m
}
