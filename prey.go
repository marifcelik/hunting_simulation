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
	foundPreys := make([]*Prey, 0)

	for xr, xl, a := p.x-3, p.x+3, 0; xr <= xl; xr, xl, a = xr+1, xl-1, a+1 {
		for y := p.y + a; y >= p.y-a; y-- {
			if xr < 0 || xl < 0 || y < 0 || xr >= AREA_SIZE || xl >= AREA_SIZE || y >= AREA_SIZE {
				continue
			}

			look := func(a IMover) {
				if a != nil {
					if animal, ok := a.(*Prey); ok && animal.kind == p.kind && animal.gender != p.gender {
						foundPreys = append(foundPreys, animal)
					}
				}
			}

			animalLeft := area.At(xl, y)
			animalRight := area.At(xr, y)
			look(animalLeft)
			look(animalRight)
		}
	}

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
