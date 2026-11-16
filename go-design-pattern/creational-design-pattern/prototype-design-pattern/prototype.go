package main

type Cloner[T any] interface {
	Clone() T
}

/*
func (e *Employee) Clone() *Employee {
	cloned := *e
	cloned.Skills = append([]string{}, e.Skills...)
	return &cloned
}
*/
