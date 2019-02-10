package main

import "fmt"

// Task: Fix the type conversion below. Does it work?
func main() {
	var x int16 = 32000
	var y int64 = x // Fix me

	fmt.Println(y)
}

