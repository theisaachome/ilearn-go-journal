package main

import (
	"fmt"
	"math"
)


func add(a int, b int) int {
	return a + b
}
func addThree(a,b,c int)int{
	return a + b + c
}

func vals ( a, b int) (int, int) {
	return a + b, a - b
}


func hypot(x,f float64) float64{
	return  math.Sqrt(x*x + f*f)
}
func main(){
	res1:= add(1, 2)
	fmt.Println("Sum of 1 and 2 is:", res1)

	res2 := addThree(1, 2, 3)
	fmt.Println("Sum of 1, 2 and 3 is:", res2)

	a,b:= vals(5, 3)
	fmt.Println(a) // addition result
	fmt.Println(b) // subtraction result

	_, c := vals(10, 4)
	fmt.Println("Subtraction result is:", c) // only print subtraction result

	result := hypot(3, 4)
	fmt.Println("Hypotenuse of a right triangle with sides 3 and 4 is:",result);
}
/*
func findLinks(url string) ([]string, error) {
	resp,err := http.Get(url)
	if err != nil {		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("failed to fetch %s: %s", url, resp.Status)
	}
	doc, err:= html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		return nil, fmt.Errorf("failed to parse HTML: %v", err)
	}

	return []string{"http://example.com", "http://example.org"}, nil
}
	*/