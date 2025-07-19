package auth

import "fmt"

func ExampleMyCustomError_Error() {
	err1 := &MyCustomError{
		Code:    404,
		Message: "Not Found",
	}
	err2 := &MyCustomError{Code: 10001, Message: "Internal Server Error"}
	// fmt.Println("Error 1 : ", err1)
	// fmt.Println("Error 2 : ", err2)
	fmt.Println(err1.Error())
	fmt.Println(err1.String())
	fmt.Println(err2.Error())
	fmt.Println(err2.String())
	// Output:
	// Error 404: Not Found
	// MyCustomError{Code: 404, Message: Not Found}
	// Error 10001: Internal Server Error
	// MyCustomError{Code: 10001, Message: Internal Server Error}
}
