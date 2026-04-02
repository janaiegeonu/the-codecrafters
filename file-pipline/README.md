This project implements a file processing pipeline that reads a raw text string, applies a series of structured transformation rules, and writes a cleaned, formatted version to a new file called output.txt

The program processes the input line-by-line, handling inconsistent formatting such as extra spaces, mixed casing, special markers (e.g. CLASSIFIED:), and invalid lines. Each line is passed through transformation functions to produce a clean refined text line by line

The final output includes properly formatted lines, optional numbering, and a standardized structure, while a summary of processing statistics (lines read, written, removed, and rules applied) is displayed in the terminal.
