package main

import "fmt"

// Task: Run the program as it is. Afterwards change x to 32000. What's the result?
func main() {
	var x int64 = 64000
	var y int16 = int16(x)

	fmt.Println(y)
}
