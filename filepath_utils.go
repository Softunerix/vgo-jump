package main

import (
	"bufio"
	"os"
	"strings"
	"unicode"
	"unicode/utf8"
)

// finding package path (this is for laravel, it will be updated later)
func GetPackagePath(filepath string, packageName string, langConfig *LanguageConfig) (string, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	packageNameLen := len(packageName)

	scanner := bufio.NewScanner(file)
	for cursorPosition := 1; scanner.Scan(); cursorPosition++ {
		currentLine := scanner.Text()
		currentLine = trimLine(currentLine)

		if len(currentLine) == 0 {
			continue
		}

		if !strings.HasPrefix(currentLine, langConfig.PackageStartsWith) {
			continue
		}
		// Remove "packageStartsWith" (like use or from) and the semicolon
		currentLine = currentLine[len(langConfig.PackageStartsWith):]
		if currentLine[len(currentLine)-1] == ';' {
			currentLine = currentLine[:len(currentLine)-1]
		}
		currentLine = trimLine(currentLine)

		// just have to match the last part to check the package name
		// if len(currentLine) > packageNameLen && currentLine[len(currentLine)-packageNameLen:] == packageName {
		// 	unixPath := fixPath(currentLine[:len(currentLine)-packageNameLen-1])
		// 	// return the path
		// 	return unixPath, nil
		// }

		// checks every words in the line
		for cursorPosition := len(currentLine); cursorPosition > 0; cursorPosition-- {
			if cursorPosition-packageNameLen < 0 {
				break
			}

			if currentLine[cursorPosition-packageNameLen:cursorPosition] == packageName {
				unixPath := fixPath(currentLine[:cursorPosition], langConfig)
				// return the path
				return unixPath, nil
			}
		}
	}

	return "", nil
}

func trimLine(s string) string {
	runes := []rune(s)
	for i := 0; i < len(runes)-1; i++ {
		if runes[i] == ' ' && runes[i+1] == ' ' {
			// Remove one space by deleting the space at i+1
			return string(append(runes[:i+1], runes[i+2:]...))
		}
	}
	return s
}

func fixPath(path string, langConfig *LanguageConfig) string {
	path = strings.TrimSpace(path)

	// Replace all separators with "/"
	path = strings.ReplaceAll(path, langConfig.Separator, "/")

	if path == "" || !langConfig.LowercaseFirstLetter {
		return path
	}

	r, size := utf8.DecodeRuneInString(path)
	return string(unicode.ToLower(r)) + path[size:]
}
