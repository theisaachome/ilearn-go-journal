package main

func main() {
	defer println("defer main func 1")
	defer println("defer main func 2")
	defer println("defer main func 3")
	defer func() {
		println("defer func")
	}()

}

// defer func() {} run in last in first out
