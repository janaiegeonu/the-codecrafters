package main

import (
	"fmt"
	"strings"
)

func Title(text string) string {

	word := strings.Fields(text)
	connectors := " a an the and but or for nor on at to by in of up as is it "

	for i, value := range word {
		if strings.Contains(connectors, " "+value+" ") {
			word[i] = value
		} else {
			word[i] = strings.Title(value)
		}
	}
	if len(word) > 0 && strings.ToLower(word[0]) == "title" {

		for i := 0; i < len(word); i++ {
		}
		word = word[1:]
	}

	word[0] = strings.Title(word[0])
	return strings.Join(word, " ")
}

func main() {
	fmt.Println(Title("title a threat in the north"))
}
