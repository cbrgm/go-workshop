# Session 01 - Basic Syntax (~45 min)

We will first have a look at the basic Go syntax. You will learn:

* how to declare constants and variables
* how to write for loops and use if.
* how types work.
* how to write your own functions.
* how packages and import statements work.

## Introduction

The Go programming language

* Modern (2009)
* Compact, general-purpose
* Imperative, statically type-checked, dynamically type-safe
* Garbage-collected
* Compiles to native code, statically linked
* Fast compilation, efficient execution

Designed by programmers for programmers! 

###  Safety first, a few examples!


Example 1: Typed, and type safe

```
var i int = -1
var u uint = 200
i = u   // nope, incompatible types
```

Example 2: Array accesses are bounds checked

```
s := make([]string, 10)
x := s[20] // will panic at runtime
```

Example 3: No implicit conversions; booleans and integers are not aliases

```
i := 2
if i { ... }    // nope, no coercion to bool
```

### Awesome support for concurrency

* Multicore CPUs are a reality.
* Multiprocessing is not a solution.
* Networking support baked into the standard library, integrated into the runtime.

### Garbage collected

Go is a garbage collected language.

* Eliminates the bookkeeping errors related to ownership of shared values.
* Eliminates an entire class of use after free and memory leak bugs.
* Enables simpler, cleaner, APIs.

The garbage collector handles heaps into the 100's of GB range!

### Opinionated

Go is an opinionated language.

* Unused local variables are an error.
* Unused imports are also an error.
* The compiler does not issue warnings, only errors.
* A single way to format code as defined by `go fmt`.

## Let's get started: Constants & Identifier

Here are some examples of constants:


    1
    "hello"
    false
    1.3


To make a constant, we declare one with the const keyword.

    const name = "Christian"
    println(name)


In Go an identifier is any word that starts with a letter.
All source code in Go is UTF-8, you can use even Emojis!

In Go an identifier is any word that starts with a letter.

    const students = 22
    const workshop = "go-workshop 😁"
    println(students, workshop)

Identifiers must start with a unicode letter, or underscore `_`. This does not work:

    const 1student = "Christian"
    const 😁workshop = "go-workshop"

## Comments


* Inline comments, which start with a double forward slash, //.
* Block comments, which start with a forward slash and a star, /*, and end with a star and forward slash, */.


```
// const a = 1

/*
const b = 2
const c = 3
*/

println(a, b, c)
```

## Declarations

There are six kinds of declarations in Go

* `const`: declares a new constant.
* `var`: declares a new variable.
* `type`: declares a new type.
* `func`: declares a new function, or method.
* `package`: declares the package this .go source file belongs to.
* `import`: declares that this package imports declarations from another.

## Variables

* A variable holds a value that can be changed over time.
* You declare a new variable with the `var` declaration.

Example:
```
var π = 3.14159
var radius = 6371.0 // radius of the Earth in km
var circumference = 2 * π * radius

println("Circumference of the earth is", circumference, "km")
```
Variables can be assigned in different ways:

```
var x int         // Variable x of type int with a zero value
var x int = v     // Variable x of type int with value v
var x = v       // Variable x with value v, implicit typing

x := v          // Short variable declaration (type inferred)
x, y := v1, v2  // Double declaration (similar with var)
```

* Unused variables are often the source of bugs.
* If you declare a variable in the scope of your function but do not use it, the Go compiler will complain.
* See: `01_variables/task_00.go`

## Incrementing / Decrementing

* Go supports a limited form of variable post-increment and post-decrement
* Examples: x++, x--, there is no pre-de/incrementing
* i++ and i-- are statements, not an expressions, they do not produce a value.

```
var i = 1
i++
println(i)
```

See: `01_variables/task_01.go`

## Looping

Go has a single for loop construct that combines

    while condition { … }
    do { … } while condition
    do { … } until condition

into one syntax. There is only `for` in Go:

    for (init statement); condition; (post statement) { … }

The parts of a for statement are:

* **init statement**: used to initalise the loop variable; i = 0.
* **condition**: user to test if the loop is done; i < 10, true means keep looping.
* **post statement**: user to increment the loop variable; i++, i = i - 1.


Our first loop:

```
var i = 0
for i = 0; i < 10; i++ {
    println(i)
}
```

* Cool: not need to put ( braces around the for condition )
* See: `02_lopps/task_00-02.go`

## Conditions

* Go has two conditional statements, `if` and `switch`.
* if is used to choose between two choices based on a condition:

```
if v > 0 {
        println("v is greater than zero")
} else {
        println("v is less than or equal to zero")
}
```

* The `else` part may be omitted, for example when checking preconditions:

```
if v == 0 {
        // nothing to do
        return
}
// handle v
```

* See: `03_conditions/task_00.go`

### Continue & break

* Unlike languages like Java, if statements in Go are often used as guard clauses.
* We say that when everything is true the code reads from the top to the bottom of the page.
* We can rewrite the previous program using a new statement, `continue`, which skips the body of the loop.
* See: `03_conditions/task_01.go`
* Also we can use `break` to break out of loops, see: `03_conditions/task_03.go`

### Switch Statement

* A switch statement runs the first case equal to the condition expression
* The cases are evaluated from top to bottom, stopping when a case succeeds.
* If no case matches and there is a default case, its statements are executed.

```
var c rune = 't'

switch c {
    case ' ', 't', 'n', 'f', 'r':
        println("Found!")
    default:
        println("Not found!)
}
```



* A fallthrough statement transfers control to the next case.
* Example below results in `2,3`

```
switch 2 { / missing expression means "true"
    case 1:
        fmt.Println("1")
        fallthrough
    case 2:
        fmt.Println("2")
        fallthrough
    case 3:
        fmt.Println("3")
}
```

## Types

Go is a strongly typed language, like Java, C, C++, and Python. Go has nine kinds of types, they are:

* `strings`: string.
* `signed integers`: int8, int16, int32, int64.
* `unsigned integers`: uint8, uint6, uint32, uint64.
* `aliases`: byte, rune, int, uint.
* `booleans`: bool.
* `IEEE floating point`: float32, float64.
* `Complex types`: complex64, complex128.
* `Compound types`: array, slice, map, struct.
* `Pointer types`: *int, *bytes.Buffer.

## Functions

* You can declare your own functions with the func declaration.
* A function's name must be a valid identifier, just like const and var.
* A simple function definitions looks like this:

```
func foo() {} // Simple function definition
```

Example:
```
func moin() {
    println("Moin")
}
```

* Functions may have parameters and return types. 
* To pass an argument to a function, the type of the argument and the type of the function's formal parameter must be the same.
* Let's have a look at `04_functions/task_00-01.go`
* There are many different ways to define a function: 

```
func f1() {}                // Simple function definition
func f2(s string, i int) {} // Function that accepts two args
func f3(s1, s2 string) {}   // Two args of the same type
func f4(s ...string) {}     // Variadic function

func f5() int {             // Return type declaration
	return 42
}

func f6() (int, string) {   // Multiple return values
	return 42, "foo"
}
```

## Packages

* A package is the unit in which software is shared and reused in Go. All Go code is arranged into packages. 
* Each source file in a package must begin with the same package declaration. 
* A package's name must be a valid identifier, just like const, var, and func.
* We already introduced the `main` package. To run a Go program it needs a `main()` function within a `main package`
* See: `05_packages/task_00.go`

 Go ships with a rich standard library of packages. This includes

* file input / output
* string handling
* compression
* encoding and decoding of JSON and XML
* network handling
* HTTP client and server


## Imports 

* The final declaration we'll cover in this section is the import declaration.
* The import declaration allows you to use code from other packages into your package.  
* When you import a package, the public types, functions, variables, types, and constants, are available with a prefix of the package's name. 
* Go already ships a lot of default packages

```
time.Now    // denotes the Now function in package time
```

You can import different packages using `import`.

```
import "fmt"
import "time"
```

You can also write:

```
import (
    "fmt"
    "time"
)
```

Example:

```
package main

import "fmt"
import "time"

func main() {
    var now = time.now()
    fmt.println(now)
}
```

## End of Section 1 - Time for a recap! :-)


What we've learned so far:

* how to declare constants and variables
* how to write for loops and use if.
* how types work.
* how to write your own functions.
* how packages and import statements work.

See you in ~15 min.
