package main

import (
	"strconv"
)


type Person struct{
	Name string
	Age  int
}
func (p Person) Greet() string {
	return "Hello, my name is " + p.Name + " and I am " + strconv.Itoa(p.Age) + " years old."
}

// receiver function to increment age
// This function modifies the Age field of the Person struct
// It is a method with a pointer receiver, allowing it to modify the original struct
func (p *Person) HaveBirthday() {
	p.Age++
}
func main(){
	p :=Person{Name: "Alice", Age: 30}
	println(p.Greet())
	p.HaveBirthday()
	println(p.Greet())
}