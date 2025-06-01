package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
)

func main() {
	if len(os.Args) < 4 {
		fmt.Println("Usage: vgo-jump <file> <line> <col>")
		os.Exit(1)
	}

	filePath := os.Args[1]
	lineNum, _ := strconv.Atoi(os.Args[2])
	colNum, _ := strconv.Atoi(os.Args[3])

	packageName, err := getWordAtPosition(filePath, lineNum, colNum)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	packagePath, err := getPackagePath(filePath, packageName)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	fmt.Println("Project root found at:", filePath, "with package path:", packagePath)

	projectRoot, err := FindProjectRoot(filePath)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	// Join project root with cleanPath to get absolute path to target file
	finalPath := filepath.Join(projectRoot, packagePath)

	fmt.Println("Opening file:", finalPath)
	cmd := exec.Command("alacritty", "-e", "vim", finalPath)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		fmt.Println("Failed to open in Vim:", err)
		os.Exit(1)
	}
}

// FindProjectRoot tries to find the root of your Laravel project by searching upwards for a marker (like composer.json)
func FindProjectRoot(startPath string) (string, error) {
	dir := filepath.Dir(startPath)

	for {
		if _, err := os.Stat(filepath.Join(dir, "artisan")); err == nil {
			return dir, nil // found root
		}

		parent := filepath.Dir(dir)
		if parent == dir {
			break // reached filesystem root
		}
		dir = parent
	}

	return "", fmt.Errorf("project root not found")
}
