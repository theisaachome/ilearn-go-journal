package main

import "fmt"

// This program defines a Person struct, creates an instance of it using a pointer,
// sets its fields, and then sets the pointer to nil.
type Person struct{
	Name string
	Age  int
}

// make() operator
func makeDemo(){
	s := make([]int,10,15)

	for i:=0;i <10;i++ {
		s[i]=i+1
	}
	fmt.Println(s)
	println("Slice created with make:", s)
}

func main(){
	person := new (Person)
	person.Name = "John"
	person.Age = 30
	println("Name:", person.Name)
	println("Age:", person.Age)
	person = nil // Setting the pointer to nil
	println("Person pointer is now nil")

	makeDemo() // Call the makeDemo function to demonstrate make
}