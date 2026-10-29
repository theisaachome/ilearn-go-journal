package main

import "fmt"

func main() {
	fmt.Println(factorial(5))

	first, last := Names()
	fmt.Println(first, last)
	// now I don't need to use the return value
	_, last = Names()
	fmt.Println(last)
}

func factorial(n int) int {
	fmt.Println("factorial(", n, ")")
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

// factorial(5) = 5 * factorial(4) = 4 * factorial(3) = 3 * factorial(2) = 2 * factorial(1) = 1

func Names() (string, string) {
	return "John", "Doe"
}
