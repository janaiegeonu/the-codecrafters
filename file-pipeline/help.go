if isLowerCaseLine(text) {
	text = strings.ToUpper(text)
}












lines := strings.Split(string(data), "\n")


var output []string
linesRead := 0
linesWritten := 0
linesRemoved := 0

for _, line := range lines {
	linesRead++

	line = helpers.Trimmer(line)

	// Example rule: remove empty lines
	if line == "" {
		linesRemoved++
		continue
	}

	line = helpers.LowToUp(line)
	line = helpers.ReplaceWord(line)
	line = helpers.ReverseWord(line)
	line = helpers.CapToTitle(line)

	output = append(output, line)
	linesWritten++
}


result := strings.Join(output, "\n")
        
err = os.WriteFile(output_file, []byte(result), 0644)