package main

import (
	"fmt"
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

func main() {
	fmt.Println(isLatin("abcdeа"))
}
