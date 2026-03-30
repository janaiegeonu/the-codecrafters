package main

import (
	"fmt"
	"strings"
)

func Reversed(text string) string {
	text = strings.ToLower(text)
	var result []rune

	if strings.HasPrefix(text, "reversed ") {
		text = strings.TrimPrefix(text, "reversed ")

		rev := []rune(text)
		for i := len(rev) - 1; i >= 0; i-- {
			result = append(result, rev[i])
		}
	}

	return string(result)

}

func main() {

	fmt.Println(Reversed("reversed hello world"))
}
