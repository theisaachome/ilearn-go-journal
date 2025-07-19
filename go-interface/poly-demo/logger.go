package polydemo

import (
	"fmt"
	"log"
	"os"
)

type Book struct {
	Title  string
	Author string
}

func (b Book) String() string {
	return fmt.Sprintf("Book Title: %s, Author: %s", b.Title, b.Author)
}

func WriteLog(s fmt.Stringer) {
	log.SetOutput(os.Stdout) // Redirect log output to standard output
	log.SetFlags(0)          // Remove all log prefixes (date, time, file)

	log.Print(s.String())
}
