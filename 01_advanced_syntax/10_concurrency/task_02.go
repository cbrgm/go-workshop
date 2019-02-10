package main

import "fmt"

// Example: Buffered channels
func main() {
	ch := make(chan int, 2) // Modifiy the buffer size and see what happens
	ch <- 1
	ch <- 5234234
	ch <- 5234234
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
