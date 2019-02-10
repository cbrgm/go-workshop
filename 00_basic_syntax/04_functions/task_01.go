package main

// Task: Write an add function that takes two parameters of type int and returns the sum of both.
// Notice that the type comes after the variable name!
func main() {
	x := 42
	y := 13

	println(x + y) // this should be add(x,y)!
}

/*
Spoiler:

func add(x int, y int) int {
	return x + y
}

func main() {
	println(add(42, 13))
}
 */