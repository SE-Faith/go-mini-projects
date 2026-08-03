package sanitizer

import (
	"strings"
)

// Function CleanString removes spaces and converts the string to Title Case

func CleanString(word string) string {
	trimmed := strings.TrimSpace(word)
	cleaned := strings.Title(trimmed)
	return cleaned
}
