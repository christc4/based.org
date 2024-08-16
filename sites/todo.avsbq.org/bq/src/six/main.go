package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	var filename string
	if len(os.Args) < 2 {
		fmt.Println("Provide filename")

		scanner := bufio.NewScanner(os.Stdin) // read standard input
		for {
			fmt.Print("Filename: ")
			scanner.Scan()
			input := scanner.Text()

			if input == "q" {
				fmt.Println("Exiting...")
				os.Exit(1)
			} else if input != "" {
				filename = input
				break
			} else {
				fmt.Println("Invalid input. Please provide a valid filename or type 'q' to quit.")
			}
		}

	} else {
		filename = os.Args[1]
	}

	fmt.Printf("Editing file: %s\n", filename)

	// Display file stats
	wordCount, lineCount, err := getFileStats(filename)
	if err != nil {
		fmt.Println("New file, word & line count, blank", err)
	} else {
		fmt.Printf("Word count: %d, Line count: %d\n", wordCount, lineCount)
	}

	modTime, err := getFileModTime(filename)
	if err != nil {
		fmt.Println("New file, created now", err)
	} else {
		fmt.Printf("Last modified: %s\n", modTime.Format(time.RFC1123))
	}

	file, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE, 0644)

	if err != nil {
		fmt.Println("Error opening file:", err)
		os.Exit(1)
	}

	defer file.Close()

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		switch line {

		case ".":
			fmt.Println("Saving and exiting...")
			return

		case "?":
			displayHelpFromFile("help.txt")
			continue

		case "p":
			fmt.Println("\n")
			printFileContents(file)
			continue
		case "s":
			fmt.Println("\nSaving")
			if err := file.Sync(); err != nil {
				fmt.Println("Error saving file:", err)
			}
			continue

		default:
			if strings.HasPrefix(line, "!") {
				runSystemCommand(line[1:])
				continue
			}

			_, err := fmt.Fprintln(file, line) // Append the line to the file
			if err != nil {
				fmt.Println("Error writing to file:", err)
				os.Exit(1)
			}
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading standard input:", err)
		os.Exit(1)
	}
	fmt.Println("File saved successfully!")
}


