package main

import "fmt"

// In this example the assignment of y = x fails because x and y are different integer types.
// This will fail:
func main() {
	var x uint = 700
	var y int = 700

	fmt.Println(x == y)
}
