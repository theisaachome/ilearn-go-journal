package main

import "fmt"

type Person struct {
	firstName string
}

// Person struct receiver method
func (p Person) Greet() {
	fmt.Println("Hello, I'm " + p.firstName)
}

func main() {
	p1 := Person{firstName: "John"}
	p2 := Person{firstName: "Jane"}

	p1.Greet()
	p2.Greet()
}

// func (r receiver) identifier(p parameters)(returns){code}
