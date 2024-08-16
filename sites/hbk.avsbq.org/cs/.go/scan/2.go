package main

import (
	"fmt"
)

func main() {

	var name string
	var num_val int
	var float_val float32
	var bool_val bool

	// Calling Scan()
	fmt.Scan(&name)
	fmt.Scan(&num_val)
	fmt.Scan(&float_val)
	fmt.Scan(&bool_val)	

	fmt.Printf("%s %d %g %t", name, num_val, float_val, bool_val)

}
