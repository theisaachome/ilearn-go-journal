package main

import "fmt"

func main() {
	x := sum([]int{1, 2, 3, 4, 5})
	fmt.Println(x)
}

// name return type
func sum(data []int) (total int) {
	total = 0
	for _, v := range data {
		total += v
	}
	return
}
