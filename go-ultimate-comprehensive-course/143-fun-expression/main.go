package main

import "fmt"

func main() {
	foo()

	x := func() {
		fmt.Println("Annonymous func")
	}
	y := func(s string) {
		fmt.Println("Annonymous func with parameter: ", s, "!")
	}

	x()
	y("Hello")
}

func foo() {
	fmt.Println("foo ran")
}
