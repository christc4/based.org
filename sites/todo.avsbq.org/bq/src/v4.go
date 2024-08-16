package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
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

func displayHelpFromFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Help file not present", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading help file:", err)
	}
}

func printFileContents(file *os.File) {
	_, err := file.Seek(0, 0) // Move the file pointer to the beginning
	if err != nil {
		fmt.Println("Error seeking file:", err)
		return
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}
}

func runSystemCommand(cmd string) {
	command := exec.Command("bash", "-c", cmd)
	output, err := command.Output()
	if err != nil {
		fmt.Println("Error running command:", err)
		return
	}
	fmt.Println(string(output))
}
