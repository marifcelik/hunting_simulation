package main

type IMover interface {
	Move() (int, int)
	Unit() int
	X() int
	Y() int
	SetX(int)
	SetY(int)
}

type IAnimal interface {
	IMover
	Kind() string
	Breed() (IAnimal, bool)
}

type IHunter interface {
	IMover
	Hunt() (IAnimal, bool)
}

type IPredator interface {
	IHunter
	IAnimal
}
