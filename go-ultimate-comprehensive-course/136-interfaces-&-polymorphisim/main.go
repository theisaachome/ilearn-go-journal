package main

import "fmt"

type person struct {
	firstName string
}

type secretAgent struct {
	person
	lk bool
}

func (p person) speak() {
	fmt.Println("Hello, I'm " + p.firstName)
}

func (s secretAgent) speak() {
	fmt.Println("Hello, I'm " + s.firstName + " and I am a secret agent")
}

// define interface for polymorphisim
type human interface {
	speak()
}

func saySomething(h human) {
	h.speak()
}

func main() {
	sa1 := secretAgent{person: person{firstName: ""}, lk: true}
	p1 := person{firstName: "John"}
	//sa1.speak()
	//p1.speak()

	// polymorphisim implementation
	saySomething(sa1)
	saySomething(p1)
}
