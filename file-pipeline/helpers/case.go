package helpers

import (
	"regexp"
	"strings"
)

func CapToTitle(text string) string {

	var word string

	cap, _ := regexp.MatchString(`[A-Z]`, text)

	if cap {
		word = strings.ToLower(text)
	}
	return strings.Title(word)

}

func LowToUp(text string) string {

	var word string

	low, _ := regexp.MatchString(`[a-z]`, text)

	if low {
		word = strings.ToUpper(text)
	}
	return word
}
