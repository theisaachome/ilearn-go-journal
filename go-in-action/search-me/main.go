package main

import (
	"log"
	"os"
)


func init(){
	log.SetOutput(os.Stdout)
}
func main(){
	log.Println("Hello, World!")
	log.Println("This is a simple Go program that logs messages to the console.")
	log.Println("Make sure to run this in an environment where you can see the output.")
	log.Println("Goodbye!")
}