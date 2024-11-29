package kotlin

import "strings"

func GetVarId(ref string) string {
	if len(ref) < 2 {
		return ref
	}

	if strings.HasPrefix(ref, "$") {
		return ref[1:]
	}

	if len(ref) < 4 {
		return ref
	}

	if strings.HasPrefix(ref, "${") && strings.HasSuffix(ref, "}") {
		return ref[2 : len(ref)-1]
	}

	return ref
}
