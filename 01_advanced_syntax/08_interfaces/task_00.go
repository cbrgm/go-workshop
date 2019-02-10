package main

import "fmt"

type Dense interface { // The Dense interface
	Density() int
}

// IsItDenser takes two interfaces as formal parameters
func IsItDenser(a, b Dense) bool {
	return a.Density() > b.Density()
}

type Rock struct {
	Mass   int
	Volume int
}

// Rock has Density() int method
func (r Rock) Density() int {
	return r.Mass / r.Volume
}

type Geode struct {
}

// Geode has Density() int method
func (g Geode) Density() int {
	return 100
}

func main() {
	r := Rock{10, 1}
	g := Geode{}

	fmt.Println(IsItDenser(g, r)) // We can pass both to IsItDenser, both implement Denser interface implicitly!
}
