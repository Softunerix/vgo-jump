package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 4 {
		fmt.Println("Usage: vgo-jump <file> <line> <col>")
		os.Exit(1)
	}

	filepath := os.Args[1]
	lineNum, _ := strconv.Atoi(os.Args[2])
	colNum, _ := strconv.Atoi(os.Args[3])

	packageName, err := getWordAtPosition(filepath, lineNum, colNum)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	packagePath, err := getPackagePath(filepath, packageName)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	fmt.Println(packagePath)
}
