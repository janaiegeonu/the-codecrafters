package main

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

	if len(word) > 0 {
		word[0] = strings.Title(word[0])
	}
	return strings.Join(word, " ")
}

// reversed function
func Reversed(text string) string {
	if strings.HasPrefix(text, "reverse ") {
		text = strings.TrimPrefix(text, "reverse ")
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
	var history []string

	fmt.Println("\t THE STRING TRANSFORMER", "\n")
	fmt.Println("Markers & their Operations ;", "\n")
	fmt.Println("1. up ----> to UPPERCASE the string input")
	fmt.Println("2. low ----> to lowercase the string input")
	fmt.Println("3. cap ----> to TitleCase the string input")
	fmt.Println("4. snake ----> to snake_case the string input")
	fmt.Println("5. title ----> to fix articles or sentence in the string input")
	fmt.Println("6. reverse ----> to reverse the string input", "\n")
	fmt.Println("7. exit ----> to immedietly end and EXIT the program", "\n")

	fmt.Println("Input Format ;", "\n")
	fmt.Println("operator <text>  |   e.g upper janai loves golang", "\n")
	fmt.Println("________________________________________________________")

	reader := bufio.NewReader(os.Stdin)

	var text string

start:
	for {
		fmt.Println()
		fmt.Print("ENTER INPUT : ")

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

		if input == "reverse" {
			fmt.Println("ERROR: No text provided. Usage: reversed <text>", "\n")
			continue
		}

		if input == "title" {
			fmt.Println("ERROR: No text provided. Usage: title <text>", "\n")
			continue
		}

		if strings.ToLower(input) == "exit" {
			fmt.Println()
			fmt.Println("exiting program ......................")
			return
		}
		text = input

		word := strings.Fields(text)
		if len(word) == 0 {
			continue
		}
		first := strings.ToLower(word[0])

		value := []string{"upper", "lower", "snake", "title", "reverse", "cap", "history"}
		found := false

		for _, v := range value {
			if first == v {
				found = true
				break
			}
		}

		if !found {
			fmt.Println("Unknown command/marker:", word[0])
			fmt.Println("Valid commands: upper, lower, cap, title, snake, reverse, exit", "\n")
			continue
		}

		break

	}

	if strings.HasPrefix(strings.ToLower(text), "upper ") {
		result := Upper(text)
		fmt.Println("RESULT :", result)
		history = append(history, text+": "+result)
		goto start
	}

	if strings.HasPrefix(strings.ToLower(text), "lower ") {
		result := Lower(text)
		fmt.Println("RESULT :", result)
		history = append(history, text+": "+result)
		goto start
	}

	if strings.HasPrefix(strings.ToLower(text), "cap ") {
		result := Cap(text)
		fmt.Println("RESULT :", result)
		history = append(history, text+": "+result)
		goto start
	}

	if strings.HasPrefix(strings.ToLower(text), "reverse ") {
		result := Reversed(text)
		fmt.Println("RESULT :", result)
		history = append(history, text+": "+result)
		goto start
	}

	if strings.HasPrefix(strings.ToLower(text), "snake ") {
		result := Snake(text)
		fmt.Println("RESULT :", result)
		history = append(history, text+": "+result)
		goto start
	}

	if strings.HasPrefix(strings.ToLower(text), "title ") {
		result := Title(text)
		fmt.Println("RESULT :", result)
		history = append(history, text+": "+result)
		goto start
	}

	if text == "history" {
		if len(history) < 5 {
			fmt.Println(history)

		} else {
			println(history[len(history)-1])
			println(history[len(history)-2])
			println(history[len(history)-3])
			println(history[len(history)-4])
			println(history[len(history)-5])

		}
		goto start
	}
}
