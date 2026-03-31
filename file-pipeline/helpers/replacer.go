package helpers

import (
	"strings"
)

func ReplaceWord(text string) string {
	text = strings.ReplaceAll(text, "CLASSIFIED:", "[REDACTED]:")
	return text
}

func Trimmer(text string) string {
	return strings.TrimSpace(text)
}
