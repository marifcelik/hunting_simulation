package main

type AnimalType int

type Animal struct {
	gender bool // true: female, false: male
	dead   bool
}

func NewAnimal(gender bool) *Animal {
	return &Animal{gender, false}
}

func (a *Animal) Die() {
	a.dead = true
}
