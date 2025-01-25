package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

// Function to convert a string to a slice of runes
func convertRunes(input string) []rune {
	return []rune(input)
}

// Function to read lines from the template file (standard, shadow, thinkertoy)
func readFile(path string) []string {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err) // Log an error and stop execution if the file cannot be opened
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text()) // Add each line to the slice
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err) // Log any errors encountered during scanning
	}
	return lines
}

// Function to calculate the indices of runes in the template file
func getRunesIndex(runes []rune) []int {
	var indices []int
	for _, r := range runes {
		// Calculate the index for each rune based on ASCII value
		index := (int(r) - 32) * 9
		indices = append(indices, index)
	}
	return indices
}

// Recursive function to print the interleaved lines for each rune
func recursiveGetLines(indices []int, lines []string, currentLine int) {
	if currentLine >= 9 {
		return // Base case: stop after 9 lines
	}
	// Print the current line for each rune
	for _, index := range indices {
		if index+currentLine < len(lines) {
			fmt.Printf("%s ", lines[index+currentLine])
		}
	}
	fmt.Println()
	// Recursive call for the next line
	recursiveGetLines(indices, lines, currentLine+1)
}

func main() {
	// Check the number of arguments passed
	if len(os.Args) != 3 {
		fmt.Println("Usage: go run . [STRING] [BANNER]") // Provide correct usage instructions
		os.Exit(1)
	}

	// Retrieve the input string
	inputString := os.Args[1]

	// Retrieve the template name
	template := os.Args[2]

	// Validate if the template is among the allowed ones
	if template != "standard" && template != "shadow" && template != "thinkertoy" {
		fmt.Println("Usage: go run . [STRING] [BANNER]") // Repeat usage instructions for invalid templates
		os.Exit(1)
	}

	// Load the corresponding template file
	templatePath := template + ".txt"
	lines := readFile(templatePath)

	// Replace occurrences of `\n` in the string with actual newlines
	inputString = strings.ReplaceAll(inputString, `\n`, "\n")

	// Split the string into segments by newlines
	segments := strings.Split(inputString, "\n")

	// Process and display each segment
	for i, segment := range segments {
		// Handle empty segments (corresponding to \n in the input)
		if segment == "" {
			if i < len(segments)-1 {
				fmt.Println() // Print a single newline for `\n`
			}
			continue
		}

		// Convert the segment into runes
		runes := convertRunes(segment)

		// Calculate the indices for the runes in the template file
		indices := getRunesIndex(runes)

		// Print the lines for each rune recursively
		recursiveGetLines(indices, lines, 0)
	}
}
