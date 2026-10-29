package main

import "fmt"

func main() {
	fmt.Println("%T\n", add)
	fmt.Println("%T\n", subtract)
	fmt.Println("%T\n", doMath)

	x := doMath(10, 20, add)
	fmt.Println(x)

	y := doMath(10, 20, subtract)
	fmt.Println(y)

	data := []int{1, 2, 3, 4, 5}
	processData(data, printValue)
}

func processData(data []int, callback func(int)) {
	for _, value := range data {
		callback(value)
	}
}

// callback function is a function that is passed as an argument to another function
func printValue(x int) {
	fmt.Println(x)
}

func add(a, b int) int {
	return a + b
}

func subtract(a, b int) int {
	return a - b
}
func doMath(a int, b int, callback func(a, b int) int) int {
	return callback(a, b)
}
