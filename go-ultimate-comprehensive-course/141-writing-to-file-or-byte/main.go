package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
)

type Book struct {
	title string
}

func (b Book) writeOut(w io.Writer) {
	w.Write([]byte(b.title))
}

func main() {
	book := Book{"Go Programming"}

	// Create a new file
	f, err := os.Create("output.txt")
	if err != nil {
		log.Fatalf("error %s", err)
	}
	defer f.Close()

	var b bytes.Buffer

	book.writeOut(f)
	book.writeOut(&b)
	// print the buffer
	fmt.Println(b.String())

}
