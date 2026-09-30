package main

import "fmt"

func main() {
	adam := 42
	fmt.Println(adam)

	a := 40
	fmt.Println(a)

	b, c, d, _, f := 1, 2, 3, 4, 5
	fmt.Println(b, c, d, f)

	// this would not work
	/*
	   b,c,d,e := 1,2,3,4,5;
	   fmt.Println(b,c,d);
	*/

	// this does work
	/*
	   var g int
	   fmt.Println(g)
	   g = 43;
	   fmt.Println(g)

	   // declare a variable to holde a VALUE of a certain TYPE
	   // Initializing a variable
	   var h int = 44;
	   fmt.Println(h)
	*/
}
