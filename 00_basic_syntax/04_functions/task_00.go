package main

// Task: Extract "moin moin" into a new function moin()
func main() {
	var i int = 0
	for i = 1; i < 3; i++ {
		println("moin")
	}
}

/*
Spoiler:

func moin() {
	println("moin")
}

func main() {
	var i int = 0
	for i = 1; i < 3; i++ {
		moin()
	}
}
 */



