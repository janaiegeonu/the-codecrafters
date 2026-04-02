package helpers

import (

	"strings"
)

func CapToTitle(text string) string {
	clean := strings.TrimSpace(text)

	
	if clean != "" && clean == strings.ToUpper(clean) && clean != strings.ToLower(clean) {
		words := strings.Fields(strings.ToLower(clean))

		for i := range words {
			words[i] = strings.ToUpper(words[i][:1]) + words[i][1:]
		}

		return strings.Join(words, " ")
	}

	return text
}

func LowToUp(text string) string {
	clean := strings.TrimSpace(text)

	if clean != "" && clean == strings.ToLower(clean) && clean != strings.ToUpper(clean) {
		return strings.ToUpper(clean)
	}

	return text
}

