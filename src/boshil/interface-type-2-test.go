package main

import (
	"fmt"
	"math"
)

type Square struct {
	side float32
}

type Circle struct {
	redius float32
}

type Shaper interface {
	Area() float32
}

func (sq *Square) Area() float32 {
	return sq.side * sq.side
}

func (ci *Circle) Area() float32 {
	return ci.redius * ci.redius * math.Pi
}

func main() {
	var areaIntf Shaper
	sq1 := new(Square)
	sq1.side = 5

	areaIntf = sq1

	switch t := areaIntf.(type) {
	case *Square:
		fmt.Printf("Type Square %T with value %v\n", t, t)
	case *Circle:
		fmt.Printf("Type Square %T with value %v\n", t, t)
	case nil:
		fmt.Println("nil value: nothing to check?")
	default:
		fmt.Printf("Unexpected type %T\n", t)
	}
}
