package main

import (
	"fmt"
	"strings"
)

func Lower(text string) string {

	word := strings.Fields(text)

	if len(word) > 0 && strings.ToUpper(word[0]) == "LOWER" {
		for i := 0; i < len(word); i++ {
			word[i] = strings.ToLower(word[i])
		}
		word = word[1:]
	}
	return strings.Join(word, " ")

}

func main() {
	fmt.Println(Lower("lower JANAI LOVE GOLANG"))
}
