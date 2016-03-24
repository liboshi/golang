package main

import "fmt"

type Simpler interface {
	Set(v int)
	Get() int
}

type Simple struct {
	value int
}

func (s *Simple) Set(v int) {
	s.value = v
}

func (s *Simple) Get() int {
	return s.value
}

func main() {
	sim := new(Simple)
	sim.Set(5)
	fmt.Printf("The valuel is %d\n", sim.Get())
}
