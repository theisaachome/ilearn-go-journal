package auth

import "fmt"

type MyCustomError struct {
	Code    int
	Message string
}

func (c *MyCustomError) Error() string {
	return fmt.Sprintf("Error %d: %s", c.Code, c.Message)
}

func (c *MyCustomError) String() string {
	return fmt.Sprintf("MyCustomError{Code: %d, Message: %s}", c.Code, c.Message)
}
