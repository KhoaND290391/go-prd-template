package main

import (
	"strings"
	"unicode"
)

//PrettyInput remove indigit character
func PrettyInput(s string) string {
	return strings.Map(
		func(r rune) rune {
			if unicode.IsDigit(r) {
				return r
			}
			return -1
		},
		s,
	)
}
