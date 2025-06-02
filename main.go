package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) < 4 {
		fmt.Println("Usage: vgo-jump <file> <line> <col>")
		os.Exit(1)
	}

	filePath := os.Args[1]
	lineNum, _ := strconv.Atoi(os.Args[2])
	colNum, _ := strconv.Atoi(os.Args[3])

	// Load config once at start
	if err := LoadConfig(); err != nil {
		fmt.Println("Failed to load config:", err)
		os.Exit(1)
	}

	// Load  language configuration for the file
	langConfig, fileExtension := GetConfigFor(filePath)
	if langConfig == nil {
		fmt.Println("No language config found for file:", filePath)
		os.Exit(1)
	}

	packageName, err := GetWordAtPosition(filePath, lineNum, colNum)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	packagePath, err := GetPackagePath(filePath, packageName, langConfig)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	// fmt.Println("Project root found at:", filePath, "with package path:", packagePath)

	projectRoot, err := FindProjectRoot(filePath, langConfig)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	// Join project root with cleanPath to get absolute path to target file
	packagePath = langConfig.AddToRootMarker + strings.TrimRight(packagePath, string(os.PathSeparator))
	finalPath := filepath.Join(projectRoot, packagePath) + "." + fileExtension

	fmt.Println(finalPath)
}

// FindProjectRoot tries to find the root of your Laravel project by searching upwards for a marker (like composer.json)
func FindProjectRoot(startPath string, langConfig *LanguageConfig) (string, error) {
	dir := filepath.Dir(startPath)

	for {
		for _, marker := range langConfig.RootMarker {
			if _, err := os.Stat(filepath.Join(dir, marker)); err == nil {
				return dir, nil // found root
			}
		}

		parent := filepath.Dir(dir)
		if parent == dir {
			break // reached filesystem root
		}
		dir = parent
	}

	return "", fmt.Errorf("project root not found")
}
