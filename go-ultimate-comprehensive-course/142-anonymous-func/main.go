package main

import "fmt"

func main() {
	foo()
	func() {
		fmt.Println("anonymous func")
	}()

	func(s string) {
		fmt.Println("This is annonymous func with parameter: ", s, "!.")
	}("Todo")
}

func foo() {
	fmt.Println("foo")
}

// a named function with an identifier
// func (r receiver) identifier(p parameters)(returns){code}
// an anonymous function
// func (p parameters)(returns){code}
// func(){}() an anonymous function with no parameters
