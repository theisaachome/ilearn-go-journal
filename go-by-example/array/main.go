package main

import "fmt"

func main() {
	fmt.Println("Hello Go Array!!!")
	//FiveDigitsArray()
	//InitArrayAndAssign()
	//CountNumberOfElement()
	TwoDimensionalArray()
}

// FiveDigitsArray create array that hold 5 digit
func FiveDigitsArray() {
	var a [5]int
	fmt.Println("emp:", a)

	// set value with index as follows:
	a[2] = 100
	fmt.Println("set:", a)
	fmt.Println(len(a))
}

// InitArrayAndAssign create array and assign values
func InitArrayAndAssign() {
	b := [5]int{100, 200, 300, 400, 500}
	fmt.Println("emp:", b)
}

// CountNumberOfElement let the compiler count the array element
func CountNumberOfElement() {
	b := [...]string{"John", "Paul", "George", "Ringo"}
	fmt.Println("emp:", b)
	fmt.Println("dcl:", b)
}

// TwoDimensionalArray create 2 Dimensional array
func TwoDimensionalArray() {
	var twoD [2][3]int
	for i := range 2 {
		for j := range 3 {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("emp:", twoD)

	my2D := [2][3]int{
		{1, 2, 3},
		{3, 4, 5},
	}
	fmt.Println("emp:", my2D)
}
