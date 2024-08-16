package main

import "fmt"

type Person struct {
    Name string
    Age  int
}

func (p Person) Greet() {
    fmt.Println("Hello, my name is", p.Name)
}

func main() {
    person := Person{Name: "Alice", Age: 30}
    person.Greet()
}

