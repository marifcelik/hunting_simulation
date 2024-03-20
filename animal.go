package main

type AnimalType int

type Animal struct {
	gender bool // true: female, false: male
}

func NewAnimal(gender bool) *Animal {
	return &Animal{gender}
}
