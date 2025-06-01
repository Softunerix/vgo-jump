package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

// Retrieve the word from the file at the specified line and column
func getWordAtPosition(filepath string, lineNum, colNum int) (string, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	currentLine := ""
	for i := 1; scanner.Scan(); i++ {
		if i == lineNum {
			currentLine = scanner.Text()
			break
		}
	}

	if currentLine == "" {
		return "", fmt.Errorf("line %d not found in file %s", lineNum, filepath)
	}
	if colNum < 1 || colNum > len(currentLine) {
		return "", fmt.Errorf("column number out of range")
	}

	colNum-- // Convert to 0-based index

	start, end := getWordPosition(currentLine, colNum)
	word := extractWord(currentLine, start, end)
	if word == "" {
		return "", fmt.Errorf("no word found at line %d, column %d", lineNum, colNum+1)
	}

	return word, nil
}

// Extract the word from the line based on start and end positions
func extractWord(line string, start, end int) string {
	if start < 0 || end > len(line) || start >= end {
		return ""
	}
	return line[start:end]
}

// start/end positions of the word in the line
func getWordPosition(line string, colNum int) (int, int) {
	start := colNum
	// getting the start position by checking if the character is a letter or digit
	// (word will have syntax or space before the start of the word)
	for start > 0 && (unicode.IsLetter(rune(line[start])) || unicode.IsDigit(rune(line[start]))) {
		start--
	}
	start++ // Move to the first character of the word

	end := colNum
	// getting the end position by checking if the character is a letter or digit
	// (word will have syntax or space before the start of the word)
	for end > 0 && (unicode.IsLetter(rune(line[end])) || unicode.IsDigit(rune(line[end]))) {
		end++
	}

	return start, end
}
