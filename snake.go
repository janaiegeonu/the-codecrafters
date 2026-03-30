package main

import (
	"fmt"
	"strings"
)

func Snake(text string) string {

	text = strings.ToLower(text)

	text = strings.ReplaceAll(text, "!", "")
	text = strings.ReplaceAll(text, ".", "")
	text = strings.ReplaceAll(text, "?", "")
	text = strings.ReplaceAll(text, ",", "")
	text = strings.ReplaceAll(text, "(", "")
	text = strings.ReplaceAll(text, ")", "")
	text = strings.ReplaceAll(text, "&", "")
	text = strings.ReplaceAll(text, ";", "")
	text = strings.ReplaceAll(text, ":", "")

	word := strings.Fields(text)
	if len(word) > 0 && strings.ToLower(word[0]) == "snake" {
		word = word[1:]
	}
	return strings.Join(word, "_")

}

func main() {
	fmt.Println(Snake("SNAKE Alert! Level 5 detected."))
}
