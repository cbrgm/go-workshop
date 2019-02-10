package main

func main() {

	// Task: Uncomment the lines below, we can break out of our infite loop!
	var i = 1
	for {
		// if i > 10 {
		//    break
		// }
		if i%2 == 0 {
			println(i)
		}
		i++
	}
}
