package utils

import (
	"net/http"
	"strconv"
)

func ContentLength(resp *http.Response) int {
	if cl := resp.Header.Get("Content-Length"); cl != "" {
		if l, err := strconv.Atoi(cl); err == nil {
			return l
		}
	}

	return 0
}
