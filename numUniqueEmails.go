package main

import "strings"

func NumUniqueEmails(emails []string) int {
	s := make(map[string]bool)

	for _, e := range emails {
		i := strings.Index(e, "@")
		l := string(e[:i])

		if j := strings.Index(l, "+"); j > -1 {
			l = string(e[:j])
		}

		l = strings.Replace(l, ".", "", -1)
		s[l+e[i:]] = true
	}

	return len(s)
}
