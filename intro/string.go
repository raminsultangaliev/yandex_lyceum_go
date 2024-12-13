package main

import (
	"unicode"
)

func isLatin(input string) bool {
	for _, r := range input {
		if !unicode.Is(unicode.Latin, r) {
			return false
		}
	}
	return true
}
