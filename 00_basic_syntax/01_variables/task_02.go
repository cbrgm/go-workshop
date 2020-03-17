package main

func main() {
  // Beispiel das es keine Referenzmodifikation gibt
  	aa := "t"
  	ab := aa
  	aa = "s"
  	println(aa)
  	println(ab)


  	// Konstanten koennen nur Konstanten zugewiesen werden
  	const x = "hallo"
  	const c = x
  	println(c)

  	// Konstante mit einer Variable initialisieren schlaegt fehl
  	var v = "hallo"
  	const d = v
  	println(d)
}
