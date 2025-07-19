package formatterdemo

import "fmt"

func ExampleAddress_Format() {
	addr := Address{Host: "localhost", Port: 8080}
	fmt.Printf("%H\n", addr) // Host: localhost
	fmt.Printf("%P\n", addr) // Port: 8080
	fmt.Printf("%#v\n", addr)

	// output:
	// localhost
	// 8080
	// formatterdemo.Address{Host: "localhost", Port: 8080}
}
