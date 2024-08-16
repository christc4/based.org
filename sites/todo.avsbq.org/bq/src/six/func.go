package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"
)

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

func getFileStats(filename string) (int, int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return 0, 0, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	wordCount := 0
	lineCount := 0

	for scanner.Scan() {
		lineCount++
		wordCount += len(strings.Fields(scanner.Text()))
	}

	if err := scanner.Err(); err != nil {
		return 0, 0, err
	}

	return wordCount, lineCount, nil
}

func getFileModTime(filename string) (time.Time, error) {
	fileInfo, err := os.Stat(filename)
	if err != nil {
		return time.Time{}, err
	}
	return fileInfo.ModTime(), nil
}
