package main

import "fmt"

type Shaper interface {
	Area() float32
	Perimeter() float32
}

type Square struct {
	side float32
}

type SquareA struct {
	side float32
}

func (sq *Square) Area() float32 {
	return sq.side * sq.side
}

func (sq *Square) Perimeter() float32 {
	return sq.side * 3
}

func (sq *SquareA) Perimeter() float32 {
	return sq.side * 4
}

func main() {
	sq1 := new(Square)
	sq1.side = 5
	sq2 := new(SquareA)
	sq2.side = 4
	areaIntf := sq1
	fmt.Printf("The square has area: %f\n", areaIntf.Area())
	fmt.Printf("The perimeter is: %f\n", areaIntf.Perimeter())
	areaIntfA := sq2
	fmt.Printf("The perimeter is: %f\n", areaIntfA.Perimeter())
}
