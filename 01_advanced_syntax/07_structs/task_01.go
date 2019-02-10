package main

import "fmt"

type Employee struct {
	Id   int
	Name string
}

// Example: Create a new employee using build-in new() function
func main() {
	// You can also get a pointer to a struct using the built-in new() function
	// It allocates enough memory to fit a value of the given struct type, and returns a pointer to it
	employeePointer := new(Employee)

	employeePointer.Id = 1000
	employeePointer.Name = "Sachin"

	fmt.Println(employeePointer)
}