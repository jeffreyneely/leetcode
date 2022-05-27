package main

import "strings"

func licenseKeyFormatting(s string, k int) string {
	s = strings.ToUpper(strings.ReplaceAll(s, "-", ""))
	i := len(s) - k

	for i > 0 {
		s = s[:i] + "-" + s[:i]
		i -= k
	}
	return string(s)
}
