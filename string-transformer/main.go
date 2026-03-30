package stringtransformer

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// upper function
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

// lower function
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

// cap function
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

// snake function
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
		for i := 0; i < len(word); i++ {
		}
		word = word[1:]
	}
	return strings.Join(word, "_")

}

//title function

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

	fmt.Println("\t THE STRING TRANSFORMER", "\n")
	fmt.Println("Markers & their Operations ;", "\n")
	fmt.Println("1. up ----> to UPPERCASE the string input")
	fmt.Println("2. low ----> to lowercase the string input")
	fmt.Println("3. cap ----> to TitleCase the string input")
	fmt.Println("4. snake ----> to snake_case the string input")
	fmt.Println("5. title ----> to fix articles or sentence in the string input")
	fmt.Println("6. reverse ----> to reverse the string input", "\n")
	fmt.Println("Input Format ;", "\n")
	fmt.Println("operator <text>  |   e.g up janai loves golang", "\n")
	fmt.Println("________________________________________________________")

	reader := bufio.NewReader(os.Stdin)

	var text string

	for {

		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		if input == "" {
			fmt.Println("  PLEASE READ THE INSTRUCTIONS  ")
			fmt.Println("Markers & their Operations ;", "\n")
			fmt.Println("1. up ----> to UPPERCASE the string input")
			fmt.Println("2. low ----> to lowercase the string input")
			fmt.Println("3. cap ----> to TitleCase the string input")
			fmt.Println("4. snake ----> to snake_case the string input")
			fmt.Println("5. title ----> to fix articles or sentence in the string input")
			fmt.Println("6. reverse ----> to reverse the string input", "\n")
			fmt.Println("Input Format ;", "\n")
			fmt.Println("operator <text>  |   e.g up janai loves golang", "\n")
			fmt.Println("________________________________________________________")
			fmt.Println()
			continue

		}

		if input == "cap" {
			fmt.Println("ERROR: No text provided. Usage: cap <text>", "\n")
			continue
		}

		if input == "lower" {
			fmt.Println("ERROR: No text provided. Usage: lower <text>", "\n")
			continue
		}

		if input == "upper" {
			fmt.Println("ERROR: No text provided. Usage: upper <text>", "\n")
			continue
		}

		if input == "snake" {
			fmt.Println("ERROR: No text provided. Usage: snake <text>", "\n")
			continue
		}

		if input == "reversed" {
			fmt.Println("ERROR: No text provided. Usage: reversed <text>", "\n")
			continue
		}

		if input == "title" {
			fmt.Println("ERROR: No text provided. Usage: title <text>", "\n")
			continue
		}

	}

}
