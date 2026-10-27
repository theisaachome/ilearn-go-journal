package main

import (
	"fmt"
	"log"
	"strconv"
)

type book struct {
	title string
}

func (b book) String() string {
	return fmt.Sprintf("Book Title: %s", b.title)
}

type count int

func (c count) String() string {
	return fmt.Sprint("This is the number ", strconv.Itoa(int(c)))
}

func logInfo(s fmt.Stringer) {
	log.Println("LOG FROM 138", s.String())
}
func main() {

	book1 := book{title: "Go Programming"}
	fmt.Println(book1.String())

	var n count = 42
	fmt.Println(n.String())

	log.Println("Hello, world!")
	log.Println("Hello, world!", "This is a test")
	logInfo(book1)
	logInfo(n)
}
