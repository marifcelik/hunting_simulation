package main

import (
	"fmt"
	"math/rand"
	"slices"
)

const AREA_SIZE = 500

var (
	area *Area

	huntersHuntedPreds = make(map[PredatorType]int)
	huntersHuntedPreys = make(map[PreyType]int)
	predatorsHunted    = make(map[PreyType]int)
	predatorsBreeded   = make(map[PredatorType]int)
	preysBreeded       = make(map[PreyType]int)
)

type Area struct {
	preys     []*Prey
	predators []*Predator
	hunter    *Hunter
	action    int
	matrix    [][]IMover
	count     int
}

func NewArea(hunter *Hunter, predators []*Predator, preys []*Prey) *Area {
	newArea := new(Area)
	newArea.hunter = hunter
	newArea.predators = predators
	newArea.preys = preys

	matrix := make([][]IMover, AREA_SIZE)
	for i := range matrix {
		matrix[i] = make([]IMover, AREA_SIZE)
	}
	newArea.matrix = matrix

	newArea.add(hunter)
	for _, p := range predators {
		newArea.add(p)
	}
	for _, p := range preys {
		newArea.add(p)
	}

	return newArea
}

func (a *Area) add(i IMover) {
	for a.At(i.X(), i.Y()) != nil {
		fmt.Println("Found a collision")
		i.SetX(rand.Intn(AREA_SIZE))
		i.SetY(rand.Intn(AREA_SIZE))
	}
	a.matrix[AREA_SIZE-i.Y()-1][i.X()] = i
}

func (a *Area) calculateTotal() int {
	count := 0
	for y := range AREA_SIZE {
		for x := range AREA_SIZE {
			if a.At(x, y) != nil {
				count++
			}
		}
	}
	a.count = count
	return count
}

func (a *Area) At(x, y int) IMover {
	return a.matrix[AREA_SIZE-y-1][x]
}

func (a *Area) AddAnimal(animal IAnimal) {
	switch animal := animal.(type) {
	case *Predator:
		a.predators = append(a.predators, animal)
	case *Prey:
		a.preys = append(a.preys, animal)
	default:
		panic("Unknown animal")
	}
	a.add(animal)
}

func (a *Area) Print() {
	a.calculateTotal()
	fmt.Printf("\ntotal number of living things in the area: %v\n", a.count)

	printMap("hunters hunts", huntersHuntedPreys)
	for k, v := range huntersHuntedPreds {
		fmt.Printf("\t%v: %v, ", k, v)
	}
	fmt.Printf("\n")

	printMap("predators hunts", predatorsHunted)
	printMap("predators breed", predatorsBreeded)
	printMap("preys breed", preysBreeded)
}

func printMap[T comparable](str string, iterate map[T]int) {
	fmt.Printf("%s : \n", str)
	if len(iterate) == 0 {
		fmt.Printf("\tnothing\n")
	} else {
		for k, v := range iterate {
			fmt.Printf("\t%v: %v, ", k, v)
		}
		fmt.Printf("\n")
	}
}

func (a *Area) Delete(x, y int) {
	a.matrix[AREA_SIZE-y-1][x] = nil
}

func Move[T IMover](movers ...T) {
	for _, m := range movers {
		if area.action >= 1000 {
			return
		}
		oldX, oldY := m.X(), m.Y()
		m.Move()
		area.Delete(oldX, oldY)
		area.add(m)

		area.action += m.Unit()
	}
}

func Hunt[T IHunter](hunters ...T) {
	for _, hunter := range hunters {
		hunt, hunted := hunter.Hunt()
		if hunted {
			area.Delete(hunt.X(), hunt.Y())

			switch hunt := hunt.(type) {
			case *Predator:
				area.predators = slices.DeleteFunc(area.predators, func(a *Predator) bool {
					return a.x == hunt.X() && a.y == hunt.Y()
				})
			case *Prey:
				area.preys = slices.DeleteFunc(area.preys, func(a *Prey) bool {
					return a.x == hunt.X() && a.y == hunt.Y()
				})
			default:
				panic("Unknown animal")
			}
		}
	}
}

func Breed[T IAnimal](breeders ...T) {
	for _, breeder := range breeders {
		breed, born := breeder.Breed()
		if born {
			area.AddAnimal(breed)
		}
	}
}

func init() {
	hunter := NewHunter(1)

	preys := make([]*Prey, 60)
	for i := 0; i < 30; i++ {
		preys[i] = NewPrey(Sheep, 2, i%2 == 0)
	}
	for i := 30; i < 40; i++ {
		preys[i] = NewPrey(Cow, 2, i%2 == 0)
	}
	for i := 40; i < 60; i++ {
		preys[i] = NewPrey(Chicken, 1, i%2 == 0)
	}

	predators := make([]*Predator, 18)
	for i := 0; i < 10; i++ {
		predators[i] = NewPredator(Wolf, []PreyType{Sheep, Chicken}, 3, i%2 == 0)
	}
	for i := 10; i < 18; i++ {
		predators[i] = NewPredator(Lion, []PreyType{Cow, Sheep}, 4, i%2 == 0)
	}

	area = NewArea(hunter, predators, preys)
}

func main() {
	for area.action < 1000 {
		Move(area.hunter)
		Move(area.predators...)
		Move(area.preys...)

		// first love
		Breed(area.preys...)
		Breed(area.predators...)

		// then fight
		Hunt(area.hunter)
		Hunt(area.predators...)
	}

	area.Print()
}
