package main

import "fmt"

type person struct {
	name string
	age  int
}

func (p person) speak() {
	fmt.Println("Hello, I'm " + p.name + " and I am " + fmt.Sprint(p.age) + " years old.")
}

func main() {
	john := person{"John", 25}
	john.speak()
}
