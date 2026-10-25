package main

import "fmt"

func main() {
	xi := []int{1, 2, 3, 4, 5}
	// dot dot dot operator to unpack the slice
	x := sum(xi...)
	fmt.Println("The sum is: ", x)
}

// func (r receiver) identifier(p parameters)(returns){code}

func sum(nums ...int) int {
	fmt.Println(nums)
	fmt.Printf("%T\n", nums)

	n := 0
	for _, v := range nums {
		n += v
	}
	return n
}
