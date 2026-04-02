package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Casing(text string) string {
	return strings.ToLower(text)
}

func Hextodec(data string) int64 {
	value, err := strconv.ParseInt(data, 16, 64)
	if err != nil {
		fmt.Println("ERROR during converstion")
	}
	return value

}

func Bintodec(data string) int64 {
	value, err := strconv.ParseInt(data, 2, 64)
	if err != nil {
		fmt.Println("ERROR during converstion")
	}
	return value

}

func Dectobin(num int) string {

	data := int64(num)
	bin := strconv.FormatInt(data, 2)
	return bin
}

func DectoHex(num int) string {

	data := int64(num)
	Hex := strconv.FormatInt(data, 16)
	return Hex
}

func main() {

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("\t JANAI BASE CONVERTER", "\n")
	fmt.Println("INSTRUCTIONS (PLEASE READ) :", "\n")
	fmt.Println("1. This is a base converter program to a BASE into another BASE & display the result/output", "\n")
	fmt.Println("2. it Support three input bases: hex, bin, dec.")
	fmt.Println("3. For (dec) input, it will output both binary and hex results.")
	fmt.Println("4. For (hex) and (bin) input, output only decimal results.", "\n")
	fmt.Println(" ACCPECTED INPUT FORMAT :", "\n")
	fmt.Println("convert [base value] [base]     e.g  convert 1E hex", "\n")
	fmt.Println("____________________________________________________________")

	var data string
	var base string

	for {

		fmt.Println()
		fmt.Print("ENTER YOUR INPUT : ")

		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		input = strings.ReplaceAll(input, " ", "")

		if input == "" {
			fmt.Println("ERROR: empty input", "\n")
			continue
		}

		data = input

		break
	}

	for {

		fmt.Print("ENTER YOUR BASE : ")

		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		input = strings.ReplaceAll(input, " ", "")

		base = Casing(input)

		switch base {
		case "bin":
			fmt.Println()
			fmt.Println("convert", data, "to bin")

		case "hex":
			fmt.Println()
			fmt.Println("convert", data, "to hex")

		case "dec":
			fmt.Println()
			fmt.Println("convert", data, "to dec")

		default:
			fmt.Println()
			fmt.Println("ERROR 404: Base Not Found")
			continue
		}

		break
	}

	if base == "hex" {
		fmt.Println()
		result := Hextodec(data)
		fmt.Println("RESULT : Decimal", result)

	}

	if base == "bin" {
		fmt.Println()
		result := Bintodec(data)
		fmt.Println("RESULT : Decimal", result)

	}

	var dec int

	if base == "dec" {

		for {
			fmt.Println()
			val, err := strconv.Atoi(data)
			if err != nil {
				fmt.Println("only NUMERIC format is allowed", "\n")
				continue
			}
			dec = val

			break
		}
		result_1 := Dectobin(dec)
		result_2 := DectoHex(dec)
		fmt.Println("RESULT : Binary", result_1)
		fmt.Println()
		fmt.Println("RESULT : Hex", strings.ToUpper(result_2))

	}

}
