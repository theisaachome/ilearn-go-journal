package main

import "fmt"

type Employee struct {
	Name string
	Age  int
}

func GetInformation() {
	// Annonymous struct
	person := struct {
		name  string
		email string
		phone string
	}{
		name:  "John",
		email: "john@doe.com",
		phone: "1234567890",
	}
	fmt.Println(person)
}

func main() {
	GetInformation()
}
