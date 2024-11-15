package i18n

import "fmt"

func Get(msg Message, args ...any) string {
	format, ok := en[msg]

	if !ok {
		return "[No translation]"
	}

	return fmt.Sprintf(format, args...)
}
