package main

import (
	"fmt"
	"math"
)

// Структура
type Point struct {
	x, y int
}

// Коснтруктор
func NewPoint(x int, y int) *Point {
	return &Point{
		x: x,
		y: y,
	}
}

// Функция получения X
func (p *Point) GetX() int {
	return p.x
}

// Функция получения Y
func (p *Point) GetY() int {
	return p.y
}

// Функция рассчета расстояния
func (p *Point) Distance(other *Point) float64 {
	return math.Sqrt(math.Pow(float64(p.GetX()-other.GetX()), 2) + math.Pow(float64(p.GetY()-other.GetY()), 2))

}

func main() {
	p1 := NewPoint(1, 0)
	p2 := NewPoint(2, 0)

	fmt.Println(p1.Distance(p2))
}
