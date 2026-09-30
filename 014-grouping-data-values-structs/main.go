package main

import "fmt"

type Person struct {
	firstName string
	lastName  string
	age       int
}

type secretAgent struct {
	person    Person
	ltk       bool
	firstName string
	lastName  string
}

func main() {
	john := Person{
		firstName: "John",
		lastName:  "Doe",
		age:       30,
	}
	agent := secretAgent{
		person: Person{
			firstName: "Secret Agent",
			lastName:  "Fisher Bond",
			age:       50,
		},
		ltk:       true,
		firstName: "James",
		lastName:  "Bond",
	}
	fmt.Println(agent)
	fmt.Println(agent.person)
	fmt.Println(agent.firstName, agent.lastName, agent.person.firstName, agent.person.lastName, agent.ltk)
	mary := Person{
		firstName: "Mary",
		lastName:  "Smith",
		age:       25,
	}
	fmt.Println(john)
	fmt.Println(mary)

	fmt.Println("%T \t %#v", john, john)
	fmt.Println("%T \t %#v", mary, mary)

	fmt.Println(john.firstName, john.lastName, john.age)

}
