package main

func main() {
	// closure function stored in a variable
	f := increment()
	println(f())
	println(f())
	println(f())
	println(f())
	println(f())
	println(f())

	// closure function stored in a variable
	// the counter is reset to 0 every time
	s := increment()
	println(s())
	println(s())
	println(s())
	println(s())
	println(s())
	println(s())

}

func increment() func() int {
	counter := 0
	return func() int {
		counter++
		return counter
	}
}
