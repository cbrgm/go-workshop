# Session 03 - Building a go application

## Go tool

*  Your Go installation comes with a tool we call the go tool, because that's its name.
* The go tool can
    * compile your programs
    * run your tests
    * display documentation for a package
    * fetch packages from the internet.

## $GOPATH

* The go tool works inside a workspace where all your Go source code is stored.
* All the source code for this workshop is included with this repository.
* You can set $GOPATH to be the base directory where you checked out this repository. eg.
* `export GOPATH=$HOME/introduction-to-go`
* Using a workspace allows you to import code from other packages with a fixed name. eg.

```
import "github.com/pkg/profile"
```

Will import the code for the profile package stored in

```
$GOPATH/src/github.com/pkg/profile
```

## go build

Go is a compiled language, so the usual work flow is

 * Edit code
 * go build
* Run program

Let's try building a Go program

* cd $GOPATH/src/helloworld
* Read the source for hello.go
* Build the source with go build
* Run the program ./helloworld

## Let's write our first app

 Let's write a clock, a program that prints out the current time

* cd $GOPATH/src/whattimeisit
* create a `main.go` and finish the program (tip: `time.Now()`)
* Build the program with, go build (if you make an error, go back and edit main.go)
* Run your program ./whattimeisit, it should print something like this

```
The current time is 2019-02-10 12:33:18.222821474 +0900 JST
```

## Testing your application

Testing is easy in Go

* We will have a look at `03_testing\sum.go`.
* Examine `03_testing\sum_test.go`
* Run `go test` in the directory

## Error Handling

* You probably spotted that lots of methods and functions in the Go standard library return a value of type error.
* error is a predeclared type, just like int, string, etc.
* error is an interface, it's declaration is

```
type error interface {
        Error() string
}

```

* Any type that has an Error() string method, implements the error interface. 

