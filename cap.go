package main

import (
	"fmt"
	"strings"
)

func Cap(text string) string {
	text = strings.ToLower(text)
	word := strings.Fields(text)
	if len(word) > 0 && strings.ToLower(word[0]) == "cap" {
		for i := 0; i < len(word); i++ {
			word[i] = strings.Title(word[i])
		}
		word = word[1:]
	}
	return strings.Join(word, " ")

}
func main() {
	fmt.Println(Cap("cap janai loves golang"))
}
