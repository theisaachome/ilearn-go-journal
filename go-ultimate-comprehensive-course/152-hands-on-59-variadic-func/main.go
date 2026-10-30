package main

import "fmt"

func main() {
	fmt.Println(foo(1, 2, 3, 4, 5))
	fmt.Println(bar(1, 2, 3, 4, 5))
}

func foo(nums ...int) int {
	sum := 0
	for _, i := range nums {
		sum += i
	}
	return sum
}

func bar(nums ...int) int {
	t := 0
	for _, i := range nums {
		t += i
	}
	return t
}
