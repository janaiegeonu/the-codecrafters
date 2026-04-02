package main

import (
	"file-pipline/helpers"
	"fmt"
	"os"
	"strings"
)

func main() {

	if len(os.Args) != 3 {
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

	lines := strings.Split(strings.ReplaceAll(string(data), "\r\n", "\n"), "\n")


var output []string
linesRead := 0
linesWritten := 0
linesRemoved := 0

for _, line := range lines {
	linesRead++

	line = helpers.Trimmer(line)

if line == "" || strings.Trim(line, "- ") == "" {
	linesRemoved++
	continue
}

line = helpers.ReplaceWord(line)
line = helpers.ReverseWord(line)
line = helpers.CapToTitle(line)
line = helpers.LowToUp(line)    

	output = append(output, line)
	linesWritten++
}


result := strings.Join(output, "\n")
        
err = os.WriteFile(output_file, []byte(result), 0644)
	if err != nil {
		fmt.Println("Error writing output file:", err)
	}

	fmt.Println("✦ Lines read    :", linesRead)
	fmt.Println("✦ Lines written :", linesWritten)
	fmt.Println("✦ Lines removed :", linesRemoved)
	fmt.Println("✦ Rules applied : Trim, Replace, Reverse, Case rules")
	
	}
