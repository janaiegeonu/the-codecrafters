package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Red(text string) string {
	red := "\033[31m"
	reset := "\033[0m"

	return red + text + reset
}

func Green(text string) string {
	green := "\033[32m"
	reset := "\033[0m"

	return green + text + reset
}

func Yellow(text string) string {
	
	yellow := "\033[33m"
	reset := "\033[0m"

	return yellow + text + reset
}

func Grey(text string) string {

	grey := "\033[30m"
	reset := "\033[0m"

	return grey + text + reset
}

func Casing(text string) string {

	return strings.ToUpper(text)
}
func main() {

	reader := bufio.NewReader(os.Stdin)

	fmt.Println(Yellow("\t JANAI CLI CALCULATOR 📟 🧮"), "\n")

	fmt.Println("INSTRUCTIONS 📄 :", "\n")
	fmt.Println(Grey(" --Enter your first number, then an operator following with your second number"), "\n")
	fmt.Println("Operations to use:", "\n")
	fmt.Println(Grey("(+) = to ADD numbers"))
	fmt.Println(Grey("(-) = to SUBSCRACT numbers"))
	fmt.Println(Grey("(/) = to DIVIDE numbers"))
	fmt.Println(Grey("(*) = to MULTIPLY numbers"))
	fmt.Println()
	fmt.Println(Grey("_______________________________________________________________"))

	var first float64
	var operator string
	var second float64

top:

	for {

		fmt.Println()
		fmt.Println()
		fmt.Print(Yellow("🔸 ENTER YOUR FIRST NUMBER 🔢 : "))

		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		input = strings.ReplaceAll(input, " ", "")

		if input == "" {
			fmt.Println(Red("ERROR : empty input"), "\n")
			continue

		}

		digits, err := strconv.ParseFloat(input, 64)
		if err != nil {
			fmt.Println(Red("ERROR : Number should be in NUMERICAL format only"), "\n")
			continue
		}

		first = digits

		break
	}

	fmt.Println(Grey("_______________________________________________________________"))

	for {

		fmt.Println()
		fmt.Print(Yellow("🔸 ENTER AN OPERATOR (+,-,/,*) 🔣 : "))

		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		if input == "" {
			fmt.Println(Red("ERROR : please input an Operator"), "\n")
			continue
		}

		operator = input

		switch operator {
		case "+":

		case "-":

		case "/":

		case "*":

		default:
			fmt.Println(Red("ERROR 404: OPERATOR Not Found!!"), "\n")
			continue

		}

		break
	}

	fmt.Println(Grey("_______________________________________________________________"))

	for {

		fmt.Println()
		fmt.Print(Yellow("🔸 ENTER YOUR SECOND NUMBER 🔢 : "))

		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		input = strings.ReplaceAll(input, " ", "")

		if input == "" {
			fmt.Println(Red("ERROR : empty input"), "\n")
			continue

		}

		digits, err := strconv.ParseFloat(input, 64)
		if err != nil {
			fmt.Println(Red("ERROR : Number should be in NUMERICAL format only"), "\n")
			continue
		}

		second = digits

		break
	}

	fmt.Println(Grey("_______________________________________________________________"))

	if operator == "+" {
		result := first + second
		fmt.Println()
		fmt.Println(Green("RESULT :"), first, "+", second, "=", result, "✅")
	}

	if operator == "-" {
		result := first - second

		fmt.Println()
		fmt.Println(Green("RESULT :"), first, "-", second, "=", result, "✅")
	}

	if operator == "/" {

		result := first / second

		fmt.Println()
		fmt.Println(Green("RESULT :"), first, "/", second, "=", result, "✅")
	}

	if operator == "*" {
		result := first * second

		fmt.Println()
		fmt.Println(Green("RESULT :"), first, "*", second, "=", result, "✅")
	}

	var again string

	for {

		fmt.Println()
		fmt.Println(Grey("_______________________________________________________________"))
		fmt.Println()
		fmt.Print(Yellow("🔴 DO YOU WISH TO CALCULATE AGAIN ?? (YES/NO) : "))

		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		if input == "" {
			fmt.Println(Red("ERROR : empty input"), "\n")
			continue
		}

		again = Casing(input)

		switch again {
		case "YES":
			goto top

		case "NO":
			fmt.Println()
			fmt.Println(Red("EXITING PROGRAM ...............😵"))
			fmt.Println()
			return
		}

	}

}
