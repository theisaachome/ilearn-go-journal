package main

type Student struct {
	Name    string
	Age     int
	Address Address
}
type Address struct {
	Street string
	City   string
	State  string
}

func main() {
	student := Student{
		Name:    "John",
		Age:     30,
		Address: Address{Street: "123 Main St", City: "New York", State: "NY"},
	}
	println(student.Name, student.Age, student.Address.Street, student.Address.City, student.Address.State)
}
