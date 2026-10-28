package main

import "fmt"

func main() {
	x := foo()
	fmt.Println(x)

	y := bar()
	fmt.Println(y())
}

func foo() int {
	return 42
}

// return a function from a function

func bar() func() int {
	return func() int {
		return 43
	}
}
