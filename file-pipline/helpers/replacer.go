package helpers

import (
	"strings"
)

func ReplaceWord(line string) string {
	line = strings.ReplaceAll(line, "CLASSIFIED:", "[REDACTED]:")
	return line
}

func Trimmer(line string) string {
	return strings.TrimSpace(line)
}
