package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}
type contactInfo struct {
	email   string
	zipcode int
}

func main() {
	alex := person{
		firstName: "Alex",
		lastName:  "Henderson",
		contact:   contactInfo{email: "alex@gmail.com", zipcode: 56100},
	}
	alex.printInfo()
	alex.updateInfo("name update")
}

func (p person) updateInfo(name string) {
	p.firstName = name
}
func (p person) printInfo() {
	fmt.Printf("First Name : %s\n", p.firstName)
	fmt.Printf("Last Name  : %s\n", p.lastName)
	fmt.Println("*****Contact****")
	fmt.Printf("Email   : %s \n", p.contact.email)
	fmt.Printf("Zipcode : %v \n", p.contact.zipcode)
}
