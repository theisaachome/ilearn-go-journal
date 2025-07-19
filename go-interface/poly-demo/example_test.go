package polydemo

func ExampleLogger() {
	book := Book{Title: "The Go Programming Language", Author: "Alan A. A. Donovan and Brian W. Kernighan"}
	WriteLog(book)
	// Output:
	// Book Title: The Go Programming Language, Author: Alan A. A. Donovan and Brian W. Kernighan
}
