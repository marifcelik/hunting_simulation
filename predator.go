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
	foundPredators := Scan(p.x, p.y, p.unitRange, func(m IMover) (*Predator, bool) {
		v, ok := m.(*Predator)
		return v, ok && v.kind == p.kind && v.gender != p.gender
	})

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
	foundPreys := Scan(p.x, p.y, p.unitRange, func(m IMover) (*Prey, bool) {
		v, ok := m.(*Prey)
		return v, ok && slices.Contains(p.canHunt, v.kind)
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

	prey = foundPreys[0]
	hunted = true
	fmt.Printf("%s hunts %s at %v,%v -> %v,%v\n", p.kind, prey.Kind(), p.x, p.y, prey.X(), foundPreys[0].Y())
	predatorsHunted[foundPreys[0].kind]++

	return
}
