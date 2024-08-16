package main

import (
    "bufio"
    "fmt"
    "io/ioutil"
    "os"
)

func main() {
    fmt.Println("Welcome to the Go Text Editor!")

    if len(os.Args) < 2 {
        fmt.Println("Please provide a filename")
        os.Exit(1)
    }

    filename := os.Args[1]
    fmt.Printf("Editing file: %s\n", filename)

    contents, err := ioutil.ReadFile(filename)
    if err != nil {
        fmt.Println("Error reading file:", err)
        os.Exit(1)
    }

    fmt.Printf("\nContents of %s:\n%s", filename, contents)

    scanner := bufio.NewScanner(os.Stdin)
    var lines []string
    for scanner.Scan() {
        line := scanner.Text()
        if line == "." {
            break
        }
        lines = append(lines, line)
    }

    newContents := []byte{}
    for _, line := range lines {
        newContents = append(newContents, []byte(line)...)
        newContents = append(newContents, '\n')
    }

    err = ioutil.WriteFile(filename, newContents, 0644)
    if err != nil {
        fmt.Println("Error writing file:", err)
        os.Exit(1)
    }

    fmt.Println("File saved successfully!")
}
