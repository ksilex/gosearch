package main

import (
	"log"
	"math"
)

// По условиям задачи, координаты не могут быть меньше 0.

type Geom struct {
	X1, Y1, X2, Y2 float64
}

func New(X1, Y1, X2, Y2 float64) Geom {
	if X1 < 0 || X2 < 0 || Y1 < 0 || Y2 < 0 {
		log.Println("Координаты не могут быть меньше нуля")
		return Geom{}
	}
	return Geom{X1: X1, Y1: Y1, X2: X2, Y2: Y2}
}

func (g Geom) CalculateDistance() (distance float64) {
	xDiff, yDiff := getDiff(g.X2, g.X1, 2), getDiff(g.Y2, g.Y1, 2)

	// возврат расстояния между точками
	return math.Sqrt(xDiff + yDiff)
}

func getDiff(p1, p2, pow float64) float64 {
	return math.Pow(p2-p1, pow)
}
