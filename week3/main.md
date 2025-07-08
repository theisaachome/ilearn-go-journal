# Attaching methods to structs


## What are methods ?

Methods are functions associated with a specific type (in this case struct) 

They are declared with a receiver arguments , which specifies the type the method operate on.

```go
type MyStruct struct {
    // fields
}

// Method with a value receiver
func (s MyStruct) MethodName() {
    // Access fields of s
}

// Method with a pointer receiver (more common for modifying the struct)
func (s *MyStruct) AnotherMethod() {
    // Access and modify fields of s
}
```

## Value vs Pointer Receivers

- Value receiver 
