package main

import "fmt"

func main() {
	foo()
}

// func (r receiver) identifier(p parameters)(returns){code}

func foo() {
	fmt.Println("Called foo function")
}
