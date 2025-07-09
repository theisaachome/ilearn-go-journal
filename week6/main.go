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

// io package provides basic interfaces to I/O primitives.
func IODemo(){
	fileName := "example.txt"

	 err := os.WriteFile(fileName, []byte("Hello, Go's io package!"), 0644);
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer os.Remove(fileName) // Clean up after ourselves

	// Method 1: Read the entire file into memory (for small files)
	content, err :=os.ReadFile(fileName)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	fmt.Println("File content:", string(content))

	// Method 2: Open and read with a buffer
	fille, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer fille.Close()
	buffer:=make([]byte, 64) // Create a buffer of 64 bytes
	for{
		n, err:=fille.Read(buffer)
		if err == io.EOF {break} // End of file
		if err != nil {
			fmt.Println("Error reading file:", err)
			return
		}
		fmt.Printf("Read %d bytes: %s\n",n,buffer)
	}
}
func main(){
	// CopyDemo()
	IODemo()
}