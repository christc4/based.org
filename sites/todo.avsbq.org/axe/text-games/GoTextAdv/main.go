package main

import (
	"fmt"
	"os"
	"math/rand"
	"time"
	"bufio"

	"github.com/fatih/color"
)

var Out *os.File
var In *os.File

var player Character

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
	Out = os.Stdout
	In = os.Stdin
}

func main() {
	player = *new(Character)
	player.Name = "Paul"
	player.Speed = 1 + rand.Intn(100)
	player.Health = 100
	player.Alive = true
	player.Weap = 1
	player.CurrentLocation = "Bridge"

	player.Play()
}

func Outputf(c string, format string, args ...interface{}) {
	s := fmt.Sprintf(format, args...)
	Output(c, s)
}

func Output(c string, args ...interface{}) {
	s := fmt.Sprint(args...)
	col := color.WhiteString
	switch(c) {
	case "green":
		col = color.GreenString
	case "red":
		col = color.RedString
	case "blue":
		col = color.BlueString
	case "yellow":
		col = color.YellowString
	}
	fmt.Fprintln(Out, col(s))
}

func UserInput(i *int) {
	fmt.Fscan(In, i)
}

func UserInputln() string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("\n >>> ")
	text, _ := reader.ReadString('\n')
	return text
}

