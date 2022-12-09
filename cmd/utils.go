package cmd

import "strings"

func joinStrings(s string, strs ...string) string {
	if len(strs) == 0 {
		return s
	}

	n := len(s)
	for _, s := range strs {
		n += len(s)
	}

	var buf strings.Builder
	buf.Grow(n)
	buf.WriteString(s)

	for _, s := range strs {
		buf.WriteString(s)
	}

	return buf.String()
}
