package main

import (
	"fmt"
	"time"
)

func sayHello() {
	fmt.Println("Hello")
}	

func main() {
	go sayHello() // Starts a new goroutine
	time.Sleep(9 * time.Second)
	fmt.Println("World")
}

