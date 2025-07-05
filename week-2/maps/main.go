package main

// This program demonstrates the use of maps in Go.
// It defines a map to store string keys and integer values.
// The map is declared but not initialized, so it will be nil until assigned.
var myMap map[string]int

func mapProduct(){
	productMap :=make(map[string]int)
	productMap["apple"] = 10
	productMap["banana"] = 20
	productMap["cherry"] = 30
	for key,value :=range productMap {
		println("Item: ", key, " Price : $", value)
	}
}
func main() {
	mapProduct()
}