package main

import (
	"fmt"
	"strings"
)

func Upper(text string) string {

	word := strings.Fields(text)

	if len(word) > 0 && strings.ToLower(word[0]) == "upper" {
		for i := 0; i < len(word); i++ {
			word[i] = strings.ToUpper(word[i])
		}
		word = word[1:]
	}

	return strings.Join(word, " ")

}

func main() {

	fmt.Println(Upper("UppER janai love golang"))
}
