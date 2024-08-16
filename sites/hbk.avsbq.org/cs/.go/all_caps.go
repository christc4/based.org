package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// Check if a filename is provided as an argument
	if len(os.Args) < 2 {
		fmt.Println("Please provide a filename.")
		return
	}

	// Get the filename from the command line arguments
	filename := os.Args[1]

	// Open the file
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)

	// Process each line in the file
	for scanner.Scan() {
		line := scanner.Text()
		uppercaseLine := strings.ToUpper(line)
		fmt.Println(uppercaseLine)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}
}

