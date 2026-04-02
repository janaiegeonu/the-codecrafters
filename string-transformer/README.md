This project is a command-line tool built in Go that performs different types of string transformations based on user input. The idea was to create a flexible text processor where a user can type a command (marker) followed by text, and the program applies the requested transformation instantly.

operators being used;

upper → converts the entire text to uppercase
lower → converts the text to lowercase
cap → capitalizes the first letter of each word
title → formats text like a proper sentence (ignores small connector words like and, the, of, etc.)
snake → converts text into snake_case and removes punctuation
reverse → reverses each word in the sentence
history → shows the most recent transformations
exit → exits the program


While working on this, I improved my understanding of:

String manipulation in Go
Building command-based CLI tools
Structuring programs with multiple helper functions
Handling user input and validation properly