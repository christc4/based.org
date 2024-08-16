package main
import "fmt"

func countDown(number int) {

  if number > -50 {
    fmt.Println(number)

    // recursive call
    countDown(number - 1)
  } else {
    fmt.Println("Countdown Stops")
  }

 }

func main() {
  countDown(3)
}
