package main

import (
	"fmt"
	"math"
)

type MyFloat float64

type Vertex struct {
	X, Y float64
}

type Abser interface {
	Abs() float64
}

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	fmt.Println("======")
	fmt.Println("Plugin YouCompleteMe was installed successfully...")
	f := MyFloat(-1.23)
	fmt.Println(f.Abs())
	v := &Vertex{3, 4}
	v.Scale(5)
	fmt.Println(v.Abs())
	var a Abser
	f = MyFloat(-math.Sqrt2)
	a = f
	fmt.Println(a.Abs())
	fmt.Println("Bello vim-go")
}
