package main

type Chair interface {
}

type VictorianChair interface {
}

type ModernChair interface{}

type FurnitureFactory interface {
	CreateChair() Chair
	CreateCoffeeTable()
	CreateSofa()
}
