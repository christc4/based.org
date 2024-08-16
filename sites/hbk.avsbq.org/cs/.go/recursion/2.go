package main
import "fmt"

func countDown(number int) {

	fmt.Println(number)

	countDown(number - 1)

	}

func main() {
	countDown(3)
}
