package main

import (
	"encoding/json"
	"fmt"
)


type Employee struct{
	Name string `json:"name"`
	Age int	`json:"age"`
	Position string `json:"position"`
	IsRemote bool `json:"is_remote"`
	Address `json:"address"`
}

type Address struct{
	Street string `json:"street"`
	City string	`json:"city"`
}

func (a Address)printAddress(){
	fmt.Println("Address : ", a.Street, " , ", a.City)
}
func (e *Employee)updateName(newName string){
   e.Name = newName	
}
func main(){
	address := Address{
		Street: "123 Main St",
		City: "New York",
	}
	emp1 := Employee{
		Name: "Alice",
		Age: 30,
		Position: "Software Engineer",
		IsRemote: true,
		Address: address,
	}

	emp1.updateName("Bob") // This won't change emp1's name due to value receiver
	fmt.Println("=======Employee Details:=========")
	fmt.Println("Name:", emp1.Name)
	fmt.Println("Age:", emp1.Age)
	fmt.Println("Position:", emp1.Position)
	fmt.Println("Remote:", emp1.IsRemote)

	emp1.printAddress() // Calling method on embedded struct

	fmt.Println("=========================")
	// anonymous struct
	// This is an example of an anonymous struct
	job := struct{
		Title string
		Salary int
	}{"Software Engineer", 80000}
	fmt.Println("Job Title:", job.Title)
	fmt.Println("Job Salary:", job.Salary)

	fmt.Println("=========================")
	jsonData, _ := json.Marshal(emp1)
	fmt.Println("JSON Data:", string(jsonData))
}