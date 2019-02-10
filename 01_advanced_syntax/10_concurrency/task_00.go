package main

import (
	"fmt"
	"time"
)

// Example: Prints out hello world 5 times
func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(1000 * time.Millisecond)
		fmt.Println(s)
	}
}

func main() {
	go say("world")
	say("hello")
}
