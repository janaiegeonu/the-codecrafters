package main

import (
	"file-pipline/helpers"
	"fmt"
	"os"
)

func main() {

	if len(os.Args) != 3 {
		fmt.Println("ERROR!!.. Usage: go run . <input.txt> <output.txt>")
		return
	}

	if len(os.Args) == 0 {
		fmt.Println("ERROR!!.. Usage: go run . <input.txt> <output.txt>")
		return
	}

	input_file := os.Args[1]
	output_file := os.Args[2]

	data, err := os.ReadFile(input_file)
	if err != nil {
		fmt.Println("Error reading input file:", err)
		return
	}

	result := helpers.Complier(string(data))

	err = os.WriteFile(output_file, []byte(result+"\n"), 0644)
	if err != nil {
		fmt.Println("Error writing output file:", err)
	}

}
