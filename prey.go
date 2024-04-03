package main

import (
	"fmt"
	"math"
	"math/rand"
	"slices"
)

type PreyType AnimalType

const (
	Sheep PreyType = iota + 1
	Cow
	Chicken
)

func (p PreyType) String() string {
	return [...]string{"Sheep", "Cow", "Chicken"}[p-1]
}

type Prey struct {
	kind PreyType
	LivingThing
	Animal
}

func NewPrey(kind PreyType, unit int, gender bool) *Prey {
	return &Prey{kind, *NewLivingThing(unit), *NewAnimal(gender)}
}

func (p *Prey) Kind() string {
	return p.kind.String()
}

func (p *Prey) Breed() (animal IAnimal, born bool) {
	foundPreys := Scan(p.x, p.y, p.unitRange, func(m IMover) (*Prey, bool) {
		v, ok := m.(*Prey)
		return v, ok && v.kind == p.kind && v.gender != p.gender
	})

	if len(foundPreys) == 0 {
		return
	}

	slices.SortFunc(foundPreys, func(a, b *Prey) int {
		return int(
			math.Abs(float64(p.x-(*a).X())) + math.Abs(float64(p.y-(*a).Y())) -
				math.Abs(float64(p.x-(*b).X())) + math.Abs(float64(p.y-(*b).Y())),
		)
	})

	nearest := foundPreys[0]
	fmt.Printf("%ss breed at %v,%v -> %v,%v ", p.kind, p.x, p.y, nearest.x, nearest.y)

	animal = NewPrey(p.kind, p.unitRange, rand.Intn(2) == 0)
	born = true
	fmt.Printf("newborn is at %v,%v\n", animal.X(), animal.Y())
	preysBreeded[p.kind]++

	return
}
