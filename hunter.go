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
	foundAnimals := make([]IAnimal, 0)

	for xr, xl, a := h.x-h.unitRange, h.x+h.unitRange, 0; xr <= xl; xr, xl, a = xr+1, xl-1, a+1 {
		for y := h.y + a; y >= h.y-a; y-- {
			if y == h.y {
				continue
			}

			if xr < 0 || xl < 0 || y < 0 || xr >= AREA_SIZE || xl >= AREA_SIZE || y >= AREA_SIZE {
				continue
			}

			look := func(a IMover) {
				if a != nil {
					a := a.(IAnimal)
					foundAnimals = append(foundAnimals, a)
				}
			}

			animalLeft := area.At(xl, y)
			animalRight := area.At(xr, y)
			look(animalLeft)
			look(animalRight)
		}
	}

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
