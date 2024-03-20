package main

import "math/rand"

type LivingThing struct {
	x         int
	y         int
	unitRange int
}

func NewLivingThing(unit int) *LivingThing {
	return &LivingThing{rand.Intn(AREA_SIZE), rand.Intn(AREA_SIZE), unit}
}

func (l *LivingThing) Move() (int, int) {
	// fmt.Printf("area.action: %v\n", area.action)
again:
	// look after for intermetiate directions
	randDirection := rand.Intn(4)
	x, y := l.x, l.y

	switch randDirection {
	case 0: // up
		y += l.unitRange
	case 1: // right
		x += l.unitRange
	case 2: // down
		y -= l.unitRange
	case 3: // left
		x -= l.unitRange
		// case 4: // up-right
		// 	x += l.unit
		// 	y += l.unit
		// case 5: // down-right
		// 	x += l.unit
		// 	y -= l.unit
		// case 6: // down-left
		// 	x -= l.unit
		// 	y -= l.unit
		// case 7: // up-left
		// 	x -= l.unit
		// 	y += l.unit
	}

	l.checkBoundaries(&x, &y)

	if area.At(x, y) != nil {
		goto again
	}

	l.x, l.y = x, y
	return x, y
}

func (l *LivingThing) checkBoundaries(x, y *int) {
	if *x < 0 {
		*x = 0
	} else if *x > AREA_SIZE-1 {
		*x = AREA_SIZE - 1
	}

	if *y < 0 {
		*y = 0
	} else if *y > AREA_SIZE-1 {
		*y = AREA_SIZE - 1
	}
}

func (l *LivingThing) X() int {
	return l.x
}

func (l *LivingThing) Y() int {
	return l.y
}

func (l *LivingThing) SetX(x int) {
	l.x = x
}

func (l *LivingThing) SetY(y int) {
	l.y = y
}

func (l *LivingThing) Unit() int {
	return l.unitRange
}
