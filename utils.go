package main

import "strings"

// Sanitize cleans user input by trimming whitespace and converting to lowercase.
func Sanitize(input string) string {
	return strings.ToLower(strings.TrimSpace(input))
}

// Validate checks if input is non-empty.
func Validate(input string) bool {
	return len(strings.TrimSpace(input)) > 0
}

// ValidateEmail checks if input looks like an email.
func ValidateEmail(input string) bool {
	return strings.Contains(input, "@") && strings.Contains(input, ".")
}
