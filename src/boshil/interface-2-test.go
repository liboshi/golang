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

type Rectangle struct {
	length, width float32
}

func (sq *Square) Area() float32 {
	return sq.side * sq.side
}

func (sq *Square) Perimeter() float32 {
	return sq.side * 3
}

func (sq *Square) ShowSide() float32 {
	return sq.side
}

func (sq *SquareA) Perimeter() float32 {
	return sq.side * 4
}

func (r *Rectangle) Area() float32 {
	return r.length * r.width
}

func (r *Rectangle) Perimeter() float32 {
	return (r.length + r.width) * 2
}

func maina() {
	sq1 := new(Square)
	sq1.side = 5
	sq2 := new(SquareA)
	sq2.side = 4
	areaIntf := sq1
	fmt.Printf("The side of square is: %f\n", areaIntf.ShowSide())
	fmt.Printf("The square has area: %f\n", areaIntf.Area())
	fmt.Printf("The perimeter is: %f\n", areaIntf.Perimeter())
	areaIntfA := sq2
	fmt.Printf("The perimeter is: %f\n", areaIntfA.Perimeter())
}

func main() {
	r := &Rectangle{5, 4}
	sq := &Square{5}
	shapes := []Shaper{r, sq}
	for n, _ := range shapes {
		fmt.Println("Shape details: ", shapes[n])
		fmt.Println("Area of this shape is:", shapes[n].Area())
	}
}
