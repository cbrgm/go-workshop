package main

import "fmt"

type Employee struct {
	Id   int
	Name string
}

// pointer indirection
func (e *Employee) SetName(name string) {
	e.Name = name
}


// Task: Do not execute the programm, what do you expect? How can we fix this?
func main() {
	emp := Employee{1, "Bargmann"}

	emp.SetName("Burgmann") // We can now call the function directly on a struct! :-)

	fmt.Println(emp.Name)
}