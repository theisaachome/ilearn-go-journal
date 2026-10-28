package main

import (
	"bytes"
	"fmt"
)

func main() {

	b := bytes.NewBufferString("Hello, ")
	fmt.Println(b.String())
	b.WriteString("World!")
	fmt.Println(b.String())
	// Reset the buffer to its initial state
	b.Reset()
	b.WriteString("Goodbye!")
	fmt.Println(b.String())

	b.Write([]byte("HelloWorldfrombytes.Buffer"))
	fmt.Println(b.String())
}
