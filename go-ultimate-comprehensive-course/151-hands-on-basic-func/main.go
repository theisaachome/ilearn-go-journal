package main

import "fmt"

func foo() int {
	return 100
}

func bar() (string, int) {
	return "Hello", 10
}

func main() {
	x := foo()
	fmt.Println(x)

	s, i := bar()
	fmt.Println("String: ", s, " Int: ", i, "!")
}
