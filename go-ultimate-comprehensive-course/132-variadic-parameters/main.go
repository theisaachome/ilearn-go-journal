package main

import "fmt"

func main() {
	x := sum(1, 2, 3, 4, 5)
	fmt.Println("The sum is: ", x)
	y := sumWithString("Hello", 1, 2, 3, 4, 5)
	fmt.Println("The sum is: ", y)
}

// func (r receiver) identifer(p parameters)(returns){code}

func sum(nums ...int) int {
	fmt.Println(nums)
	fmt.Printf("%T\n", nums)

	n := 0
	for _, v := range nums {
		n += v
	}
	return n
}

// variadic parameters should be the last parameter

func sumWithString(s string, nums ...int) int {
	fmt.Println(nums)
	fmt.Printf("%T\n", nums)

	n := 0
	for _, v := range nums {
		n += v
	}
	return n
}
