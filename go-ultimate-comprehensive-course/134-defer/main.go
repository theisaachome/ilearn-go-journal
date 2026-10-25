package main

import "fmt"

// A "defer statement invokes a function whose execution is deferred to the
// the moment the surrounding function returns
// the surrendered control to the caller"
// "The deferred function is called after the surrounding function returns,
// even if the surrounding function panics

func main() {
	defer foo()
	bar()
}

// func (r receiver) identifier(p parameters)(returns){code}

func foo() {
	fmt.Println("foo")
}
func bar() {
	fmt.Println("bar")
}
