package main

import (
	"fmt"
	"math"
	"slices"
)

type Hunter struct {
	LivingThing
}

func NewHunter(unit int) *Hunter {
	return &Hunter{*NewLivingThing(unit)}
}

// Search is a method that searches for prey and retruns the first prey found
//
// if there is no prey found it returns nil
func (h *Hunter) Hunt() (hunt IAnimal, hunted bool) {
	foundAnimals := Scan(h.x, h.y, h.unitRange, func(m IMover) (IAnimal, bool) {
		v, ok := m.(IAnimal)
		return v, ok
	})

	slices.SortFunc(foundAnimals, func(a, b IAnimal) int {
		return int(
			math.Abs(float64(h.x-a.X())) + math.Abs(float64(h.y-a.Y())) -
				math.Abs(float64(h.x-b.X())) + math.Abs(float64(h.y-b.Y())),
		)
	})

	if len(foundAnimals) == 0 {
		return
	}

	hunt = foundAnimals[0]
	hunted = true
	fmt.Printf("The hunter hunts %s at %v,%v -> %v,%v\n", hunt.Kind(), h.x, h.y, hunt.X(), hunt.Y())

	if p, ok := hunt.(*Prey); ok {
		huntersHuntedPreys[p.kind]++
	} else {
		huntersHuntedPreds[hunt.(*Predator).kind]++
	}

	return
}
