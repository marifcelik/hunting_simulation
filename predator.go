package main

import (
	"fmt"
	"math"
	"math/rand"
	"slices"
)

type PredatorType AnimalType

const (
	Wolf PredatorType = iota + 4
	Lion
)

func (p PredatorType) String() string {
	return [...]string{"Wolf", "Lion"}[p-4]
}

type Predator struct {
	kind    PredatorType
	canHunt []PreyType
	Hunter
	Animal
}

func NewPredator(kind PredatorType, canHunt []PreyType, unit int, gender bool) *Predator {
	return &Predator{kind, canHunt, *NewHunter(unit), *NewAnimal(gender)}
}

func (p *Predator) Kind() string {
	return p.kind.String()
}

// TODO Write a general method to reduce code repetition in Hunt and Breed methods

func (p *Predator) Breed() (animal IAnimal, born bool) {
	foundPredators := make([]*Predator, 0)

	for xr, xl, a := p.x-3, p.x+3, 0; xr <= xl; xr, xl, a = xr+1, xl-1, a+1 {
		for y := p.y + a; y >= p.y-a; y-- {
			if xr < 0 || xl < 0 || y < 0 || xr >= AREA_SIZE || xl >= AREA_SIZE || y >= AREA_SIZE {
				continue
			}

			look := func(a IMover) {
				if a != nil {
					if animal, ok := a.(*Predator); ok && animal.kind == p.kind && animal.gender != p.gender {
						foundPredators = append(foundPredators, animal)
					}
				}
			}

			animalLeft := area.At(xl, y)
			animalRight := area.At(xr, y)
			look(animalLeft)
			look(animalRight)
		}
	}

	if len(foundPredators) == 0 {
		return
	}

	slices.SortFunc(foundPredators, func(a, b *Predator) int {
		return int(
			math.Abs(float64(p.x-(*a).X())) + math.Abs(float64(p.y-(*a).Y())) -
				math.Abs(float64(p.x-(*b).X())) + math.Abs(float64(p.y-(*b).Y())),
		)
	})

	nearest := foundPredators[0]
	fmt.Printf("%ss breed at %v,%v -> %v,%v. ", p.kind, p.x, p.y, nearest.x, nearest.y)

	animal = NewPredator(p.kind, p.canHunt, p.unitRange, rand.Intn(2) == 0)
	born = true
	fmt.Printf("newborn is at %v,%v\n", animal.X(), animal.Y())
	predatorsBreeded[p.kind]++

	return
}

func (p *Predator) Hunt() (prey IAnimal, hunted bool) {
	foundPreys := make([]*Prey, 0)

	for xr, xl, a := p.x-p.unitRange, p.x+p.unitRange, 0; xr <= xl; xr, xl, a = xr+1, xl-1, a+1 {
		for y := p.y + a; y >= p.y-a; y-- {
			if xr < 0 || xl < 0 || y < 0 || xr >= AREA_SIZE || xl >= AREA_SIZE || y >= AREA_SIZE {
				continue
			}

			look := func(a IMover) {
				if a != nil {
					preyLeft, ok := a.(*Prey)
					if ok && slices.Contains(p.canHunt, preyLeft.kind) {
						fmt.Printf("FOUND prey: %v at %v,%v\n", preyLeft.kind, preyLeft.X(), preyLeft.Y())
						foundPreys = append(foundPreys, preyLeft)
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

	prey = foundPreys[0]
	hunted = true
	fmt.Printf("%s hunts %s at %v,%v -> %v,%v\n", p.kind, prey.Kind(), p.x, p.y, prey.X(), foundPreys[0].Y())
	predatorsHunted[foundPreys[0].kind]++

	return
}
