package helpers

import (
	"strings"
)

func ReverseWord(text string) string {

	if strings.Contains(text, "REVERSE") {

		text = strings.ReplaceAll(text, "REVERSE", "")

		words := strings.Fields(text)

		for i, word := range words {
			var reversed []string
			for j := len(word) - 1; j >= 0; j-- {
				reversed = append(reversed, string(word[j]))
			}
			words[i] = strings.Join(reversed, "")
		}

		return strings.Join(words, " ")
	}

	return text
}