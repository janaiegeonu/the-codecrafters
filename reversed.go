package main

import (
	"fmt"
	"strings"
)

func Reversed(text string) string {
	if strings.HasPrefix(text, "reversed ") {
		text = strings.TrimPrefix(text, "reversed ")
	}

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

func main() {

	fmt.Println(Reversed("reversed hello world"))
}
