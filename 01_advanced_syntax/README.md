# Session 02 - Advanced Syntax (~60 min)

Welcome back! In the second session we will cover:

* how zero values work in Go
* Equality and Type Conversions
* Pointers
* Arrays
* Slices
* Ranges
* Maps
* Structs
* Defer
* Concurrency

## Zero Values

* In Go, there is no unitialised memory. The Go runtime will always ensure that the memory allocated for each variable is initalised before use. 
* For example: `var name string` Then the memory assigned to the variable `name` will be zeroed, as we have not provided an initaliser. 
* Result: Value of `name` will be "" because it's the value f a string with zero length

Zero values summarized:

* The zero value for integer types: int, int8, uint, uint64, etc, is 0.
* The zero value for floating point types: float32, float64, complex128, etc, is 0.0.
* The zero value for arrays is the zero value for each element, ie. [3]int is 0, 0, 0.
* The zero value for slices is nil.
* The zero value for boolean is false.
* The zero value for structs is the zero value for each field.

## Equality & Type Conversions

* As Go is a strongly typed language, for two variables to be equal, both their type and their value must be equal. 
* See: `00_equality/task_00.go`
* Sometimes you have variables of different integer types, you can convert from one type to another using a conversion expression. 
* The expression T(v) converts the value v to the type T.

Example:

```
package main

import "fmt"

func main() {
    var x uint = 700
    var y int = int(x) // Type Conversion from uint to int

    fmt.Println(y)
}
```

* See: `01_types/task_00-01.go`
* **Attention**: When you convert a variable with a smaller number of bits to a variable with a larger number of bits, this is fine, because they all fit. 
* When you convert a variable with a larger number of bits to a variable with a smaller number of bits there is a risk of truncation, because there are less bits available to represent your number. 

## Pointers

* Go has pointers. A pointer holds the memory address of a value. 
* The type `*T` is a pointer to a `T` value. Its zero value is `nil`. 
* Example: `var p *int`
* The `&` operator generates a pointer to its operand. 

```
i := 42
p = &i
```
* The * operator denotes the pointer's underlying value. Example:

```
i := 42
p = &i

fmt.Println(*p) // read i through the pointer p
*p = 21         // set i through the pointer p
```

* See: `02_pointers/task_00.go`

## Arrays

* An array type definition specifies a length and an element type.
* For example, the type `[4]int` represents an array of four integers. An array's size is fixed; its length is part of its type.
* Example: `[4]int` and `[5]int` are distinct, incompatible types
* Arrays can be indexed in the usual way, so the expression s[n] accesses the nth element, starting from zero. 

Example:
```
var a [4]int
a[0] = 1
i := a[0]

// i == 1
```

To initialize a string array you can write:

```
b := [2]string{"Christian", "Bargmann"}
```

Or, you can have the compiler count the array elements for you: 

```
b := [...]string{"Christian", "Bargmann"}

```

See: `03_types/task_00.go`

## Slices

* Slices are like references to arrays
* Go's slice type provides a convenient and efficient means of working with sequences of typed data
* A slice literal is declared just like an array literal, except you leave out the element count: 
* Example: `letters := []string{"a", "b", "c", "d"}`
* A slice is formed by specifying two indices, a low and high bound, separated by a colon:  `a[low : high]`

Example:
```
package main

import "fmt"

func main() {
	primes := [6]int{2, 3, 5, 7, 11, 13}

	var s []int = primes[1:4]
	fmt.Println(s)
}

// Returns 3,5,7

```

* See: `04_slices/task_00.go`
* A slice does not store any data, it just describes a section of an underlying array. 
* Changing the elements of a slice modifies the corresponding elements of its underlying array. 
* See: `04_slices/task_01.go`

### Slice Literals

* A slice literal is like an array literal without the length. 

Example: This is an array literal:
```
[3]bool{true, true, false}
```

Example: This is a slice literal. It creates the same array as above, then builds a slice that references it: 

```
[]bool{true, true, false}
```

You can also use Go's build in `make` function to create a new slice:

```
package main

import "fmt"

func main() {
    i := make([]int, 20)
    fmt.Println(len(i))
}
```

### Length & Capacity

* A slice has both a **length** and a **capacity**.
* The length of a slice is the number of elements it contains.
* The capacity of a slice is the number of elements in the underlying array, counting from the first element in the slice.
* The length and capacity of a slice s can be obtained using the expressions `len(s` and `cap(s)`
* See: `04_slices/task_03.go`
* The zero value of a slice is `nil`. A `nil` slice has a length and capacity of 0 and has no underlying array.  Example: `var s []int`

### Appending to slices

* It is common to append new elements to a slice, and so Go provides a built-in append function. The documentation of the built-in package describes append. 
* The resulting value of append is a slice containing all the elements of the original slice plus the provided values. 
* If the backing array of s is too small to fit all the given values a bigger array will be allocated. The returned slice will point to the newly allocated array. 
* See: `04_slices/task_04.go`

## Range

* The range form of the for loop iterates over a slice or map. 
* When ranging over a slice, two values are returned for each iteration. The first is the index, and the second is a copy of the element at that index. 

Example:
```
package main

import "fmt"

var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

func main() {
	for i, v := range pow {
		fmt.Printf("2*%d = %d\n", i, v)
	}
}

```
* See: `04_range/task_05.go`
* You can skip the index or value by assigning to `_`.
* If you only want the index, drop the , value entirely. 

## Maps

* Go has a built in Hash Map type, called a `map`. 
* Maps map values of key type K to values of type V 
* Example: `var m map[string]int`
* Just like making a slice, making a map is accomplished with the `make` built-in. 

Example:
```
package main

import "fmt"

func main() {
    m := make(map[string]int)
    fmt.Println(m)
}
```

Inserting a value into a map looks similar to assigning a value to a slice element. 

```
package main

import "fmt"

func main() {
    days := make(map[int]string)
    days[1] = "Monday"
    days[2] = "Tuesday"
    days[3] = "Wednesday"
    days[4] = "Thursday"
    days[5] = "Friday"
    days[6] = "Saturday"
    days[7] = "Sunday"
    fmt.Println(days)
}

```

 You can also use the sort literal initialization:
 
 ```
 package main
 
 import "fmt"
 
 func main() {
     days := map[int]string{
         1: "Monday",
         2: "Tuesday",
         3: "Wednesday",
         4: "Thursday",
         5: "Friday",
         6: "Saturday",
         7: "Sunday",
     }
     fmt.Println(days)
 }
 ```
 
 ### Retrieve values from a map
 
 Just like a slice, you can retrieve the value stored in a map with the syntax m[key]. 
 Example:
 
```
package main

import "fmt"

func main() {
    // map of planets to their number of moons
    moons := map[string]int{
        "Mercury": 0,
        "Venus":   0,
        "Earth":   1,
        "Mars":    2,
        "Jupiter": 67,
    }

    fmt.Println("Earth:", moons["Earth"])
    fmt.Println("Neptune:", moons["Neptune"])
}

```

### Check if value in map exists
Map look ups support a second syntax.
Example:

```
package main

import "fmt"

func main() {
    moons := map[string]int{"Mercury": 0, "Venus": 0, "Earth": 1, "Mars": 2, "Jupiter": 67}

    n, found := moons["Earth"]
    fmt.Println("Earth:", n, "Found:", found)

    n, found = moons["Neptune"]
    fmt.Println("Neptune:", n, "Found:", found)
}
```

### Delete a value from map

To delete a value from a map, you use the built in delete function. 
Example:

```
package main

import "fmt"

func main() {
    moons := map[string]int{"Mercury": 0, "Venus": 0, "Earth": 1, "Mars": 2, "Jupiter": 67}

    const planet = "Mars"

    n, found := moons[planet]
    fmt.Println(planet, n, "Found:", found)

    delete(moons, planet) // Removing element from map

    n, found = moons[planet]
    fmt.Println(planet, n, "Found:", found)
}
```

### Iterating over a map

We can use the `range` operator presented earlier :-)

```
for key, value := range myMap {
    fmt.Println("Key:", key, "Value", value)
}
```

## Structs

* A struct is a collection of fields. They allow us to create our "own" types in Go.
* For example we can create a Person Type like this:

```
type Person struct {
	FirstName string
	LastName  string
	Age       int
}

```

* The `type` keyword introduces a new type. 
* It is followed by the name of the type (Person) and the keyword struct to indicate that we’re defining a struct. 
* The struct contains a list of fields inside the curly braces. Each field has a name and a type.

### Creating a struct

You can initialize a variable of a struct type using a struct literal like so:

```
// Initialize a struct by supplying the value of all the struct fields.
var p = Person{"Christian", "Bargmann", 25}
```

* Note: Note that you need to pass the field values in the same order in which they are declared in the struct (This is not Python! :-))
* Go also supports the name: value syntax for initializing a struct (the order of fields is irrelevant when using this syntax).

```
// Initialize a struct by supplying the value of all the struct fields.
var p = Person{FirstName: "Christian", LastName: "Bargmann", Age: 25}
```

or like this:

```
var p = Person{
    FirstName: "Christian",
    LastName: "Bargmann", 
    Age: 25,
}
```


* You can also use the built-in `new()` function to create an instance of a struct. 
* The `new()` function allocates enough memory to fit all the struct fields, sets each of them to their zero value and returns a pointer to the newly allocated struct.
* See: `06_structs/task_01.go`

### Accessing struct fields

* You can access individual fields of a struct using the dot (.) operator -

```
c := Car{
	Name:       "Ferrari",
	Model:      "GTC4",
	Color:      "Red",
	WeightInKg: 1920,
}

// Accessing struct fields using the dot operator
fmt.Println("Car Name: ", c.Name)
fmt.Println("Car Color: ", c.Color)
```

You can also access struct fields through pointers

```
c := Car{
	Name:       "Ferrari",
	Model:      "GTC4",
	Color:      "Red",
	WeightInKg: 1920,
}

pc := &c // Create a pointer to struct c

// Accessing struct fields through a pointer using the dot operator
fmt.Println("Car Name: ", pc.Name)
fmt.Println("Car Color: ", pc.Color)
```

### "OOP"

* Using pointer receivers, we can do stuff that feels like "object orientated programming"
* Lets have a look at: `06_structs/task_02-03.go`

## Interfaces

* An interface type is defined as a set of method signatures. 
* A value of interface type can hold any value that implements those methods. 
* **Interfaces are implemented implicitly** - no need for `foo implements bar`
* A type implements an interface by implementing its methods. There is no explicit declaration of intent, no "implements" keyword. 

Example: Dense Interface implemented by Geode + Rock:

```
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

```

* See: `08_interfaces/task_00.go`

### The empty interface 

Every type implements the empty interface `interface{}`. This allows us to pass any type to a function. Whoop!

```
func PrintIt(a interface{}) {
  // a can have any methods. We dont care. :-)
  fmt.Println(a)
}

```

## Defer

* A defer statement defers the execution of a function until the surrounding function returns. 
* The deferred call's arguments are evaluated immediately, but the function call is not executed until the surrounding function returns. 
* It's like Java's try-with-resources :-)

Example:

```
package main

import "fmt"

func main() {
	defer fmt.Println("world")
	fmt.Println("hello")
}

// prints hello world
```

* Deferred function calls are pushed onto a stack. When a function returns, its deferred calls are executed in last-in-first-out order. 

```
package main

import "fmt"

func main() {
	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")
}

```

Returns:

```
counting
done
9
8
7
6
5
4
3
2
1
0
```

* `defer` is useful when opening resources like database connections or file syscalls :-)
* See: `09_defer/task_00.go`

## Optional: Concurrency

* Concurrent programming is a large topic but it’s also one of the most interesting aspects of the Go language.
* Do not communicate by sharing memory; instead, share memory by communicating.
* Go encourages a different approach in which shared values are passed around on channels and, in fact, never actively shared by separate threads of execution.

### Go Routines

* A goroutine is a lightweight thread managed by the Go runtime. 
* Example: `go f(x, y, z)` starts a new goroutine running `f(x,y,z)`
* The evaluation of f, x, y, and z happens in the current goroutine and the execution of f happens in the new goroutine. 
* Goroutines run in the same address space, so access to shared memory must be synchronized.

### Channels

* Channels are a typed conduit through which you can send and receive values with the channel operator, `<-`. 

```
ch <- v    // Send v to channel ch.
v := <-ch  // Receive from ch, and
           // assign value to v.
```

* Like maps and slices, channels must be created before use
* Example: `ch := make(chan int)`
* By default, sends and receives block until the other side is ready. This allows goroutines to synchronize without explicit locks or condition variables. 
* See: `10_concurrency/task_00.go`

### Buffered Channels

* Channels can be buffered. Provide the buffer length as the second argument to make to initialize a buffered channel
* Example: `ch := make(chan int, 100)`
* Sends to a buffered channel block only when the buffer is full. Receives block when the buffer is empty. 
* See: `10_concurrency/task_01.go`

### Range and close

* A sender can close a channel to indicate that no more values will be sent. Receivers can test whether a channel has been closed by assigning a second parameter to the receive expression: after 
* Example: `v, ok := <-ch`
* ok is false if there are no more values to receive and the channel is closed. 
* See: `10_concurrency/task_02-03.go`

### Select Statement

* The select statement lets a goroutine wait on multiple communication operations.
* A select blocks until one of its cases can run, then it executes that case. It chooses one at random if multiple are ready.
* See: `10_concurrency/task_04.go`

# End of section 02

In this section we convered:

* how zero values work in Go
* Equality and Type Conversions
* Pointers
* Arrays
* Slices
* Ranges
* Maps
* Structs
* Defer
* Concurrency

