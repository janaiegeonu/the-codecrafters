This project is a command-line calculator built with Golang. The idea behind it was to create an interactive terminal tool that can perform basic arithmetic operations while also improving the overall user experience with styled output in the terminal section.

How my code works;

The program walks the user through a step-by-step input process:

Enter the first number
Choose an operator (+, -, *, /)
Enter the second number

Once all inputs are valid, the program performs the calculation and displays the result.

It also gives the user the option to keep calculating without restarting the program.

 Terminal Styling

 One of the main features of this project is the use of ANSI escape codes to style the terminal output.

helper functions were created to color different parts of the interface:

Red()  for error messages
Green()  for results
Yellow()  for prompts and headings
Grey() for instructions and separators

The program includes validation to prevent common issues:

Handles empty input
Ensures numbers are in the correct format
Validates that only supported operators are used
Keeps prompting the user until valid input is provided