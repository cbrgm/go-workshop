package main

import "fmt"

type Employee struct {
	Id   int
	Name string
}

func SetName(e Employee, name string) {
	e.Name = name
}


// Task: Do not execute the programm, what do you expect? How can we fix this?
func main() {
	emp := Employee{1, "Bargmann"}
	SetName(emp, "Burgmann")

	fmt.Println(emp.Name)
}