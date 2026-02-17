package main

import "strings"

// Sanitize cleans user input.
func Sanitize(input string) string {
	return strings.TrimSpace(input)
}

// Validate checks if input is non-empty.
func Validate(input string) bool {
	return len(strings.TrimSpace(input)) > 0
}
