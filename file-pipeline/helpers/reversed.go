package helpers

import (
	"strings"
)

func ReverseWord(text string) string {

	var words string

	if strings.Contains(text, "REVERSE") {

		text = strings.ReplaceAll(text, "REVERSE", "")

		words := strings.Fields(text)

		for i, word := range words {
			Reverse := []string{}
			for j := len(word) - 1; j >= 0; j-- {
				Reverse = append(Reverse, string(word[j]))

			}
			words[i] = strings.Join(Reverse, "")
		}
		return strings.Join(words, " ")

	}
	return words
}

