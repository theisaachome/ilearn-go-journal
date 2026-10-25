package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	favIC     []string
}

func main() {
	p1 := person{
		firstName: "John",
		lastName:  "Doe",
		favIC:     []string{"instagram", "twitter", "facebook"},
	}
	p2 := person{
		firstName: "Jane",
		lastName:  "Doe",
		favIC:     []string{"instagram", "twitter"},
	}
	fmt.Println(p1)
	fmt.Println(p2)
	for _, v := range p1.favIC {
		fmt.Println(v)
	}
	fmt.Println("%s favorite social media sites are:", p2.firstName)
	for _, v := range p2.favIC {
		fmt.Println(v)
	}
}
