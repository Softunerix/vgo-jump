package main

import (
	"bufio"
	"os"
	"unicode"
	"unicode/utf8"
)

// finding package path (this is for laravel, it will be updated later)
func getPackagePath(filepath string, packageName string) (string, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return "", err
	}
	defer file.Close()
	// have to find the path by checking the "use" statement in the file

	packageNameLen := len(packageName)

	scanner := bufio.NewScanner(file)
	for cursorPosition := 1; scanner.Scan(); cursorPosition++ {
		currentLine := scanner.Text()
		currentLine = trimLine(currentLine)

		if len(currentLine) == 0 {
			continue
		}

		// Remove "use " and the semicolon
		currentLine = currentLine[4:]
		if currentLine[len(currentLine)-1] == ';' {
			currentLine = currentLine[:len(currentLine)-1]
		}
		currentLine = trimLine(currentLine)

		// just have to match the last part to check the package name
		if len(currentLine) > packageNameLen && currentLine[len(currentLine)-packageNameLen:] == packageName {
			unixPath := fixPath(currentLine[:len(currentLine)-packageNameLen-1])
			// return the path
			return unixPath, nil
		}
	}

	return "", nil
}

func trimLine(s string) string {
	return string([]rune(s))
}

func fixPath(path string) string {
	// Replace backslashes with forward slashes
	path = string([]rune(path))
	for i := 0; i < len(path); i++ {
		if path[i] == '\\' {
			path = path[:i] + "/" + path[i+1:]
		}
	}

	// make the first letter lowercase
	if path == "" {
		return path
	}
	r, size := utf8.DecodeRuneInString(path)
	return string(unicode.ToLower(r)) + path[size:]
}
