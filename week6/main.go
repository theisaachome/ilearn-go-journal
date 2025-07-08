package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func CopyDemo(){
	reader :=strings.NewReader("Hello, Go's io package!")

	fmt.Println("Copying data from string reader to standard output:")
	n,err :=io.Copy(os.Stdout,reader)
	if err != nil {
		fmt.Println("Error during copy:", err)
		return
	}

	fmt.Printf("\n%d bytes copied successfully.\n", n)

	var builder strings.Builder
	readerII :=strings.NewReader("Hello,Another Go's io package with a string builder!");
	n2,err2 :=io.Copy(&builder, readerII)
	if err2 != nil {
		fmt.Println("Error during copy:", err)
		return
	}
	fmt.Printf("\n%d bytes copied successfully to string builder.\n", n2)
	fmt.Println("\nCopying data from string reader to string builder:")
}
func main(){
	CopyDemo()
}