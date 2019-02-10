package main

import "fmt"

// Task: First run the program. Uncomment the line below afterwards. What behavior do you see? Any ideas?
func main() {
	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)

	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b)

	// b[0] = "XXX"
	fmt.Println(a, b)
	fmt.Println(names)
}
