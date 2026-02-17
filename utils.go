package main

import "strings"

// SanitizeInput cleans and normalizes user input.
func SanitizeInput(input string) string {
	return strings.ToLower(strings.TrimSpace(input))
}

// ValidateInput checks if input is non-empty after sanitization.
func ValidateInput(input string) bool {
	return len(SanitizeInput(input)) > 0
}

// ValidateLength checks if input meets minimum length requirement.
func ValidateLength(input string, minLen int) bool {
	return len(strings.TrimSpace(input)) >= minLen
}
